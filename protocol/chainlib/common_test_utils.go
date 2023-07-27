package chainlib

import (
	"context"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/server/grpc/gogoreflection"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/protocol/chaintracker"
	"github.com/lavanet/lava/protocol/common"
	"github.com/lavanet/lava/protocol/lavasession"
	testcommon "github.com/lavanet/lava/testutil/common"
	keepertest "github.com/lavanet/lava/testutil/keeper"
	plantypes "github.com/lavanet/lava/x/plans/types"
	spectypes "github.com/lavanet/lava/x/spec/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/proto/tendermint/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type mockResponseWriter struct {
	blockToReturn *int
}

func (mockResponseWriter) Header() http.Header {
	return http.Header{}
}

func (mockResponseWriter) Write(in []byte) (int, error) {
	return 0, nil
}

func (mrw mockResponseWriter) WriteHeader(statusCode int) {
	*mrw.blockToReturn = statusCode
}

type myServiceImplementation struct {
	*tmservice.UnimplementedServiceServer
	serverCallback http.HandlerFunc
}

func (bbb myServiceImplementation) GetLatestBlock(ctx context.Context, reqIn *tmservice.GetLatestBlockRequest) (*tmservice.GetLatestBlockResponse, error) {
	metadata, exists := metadata.FromIncomingContext(ctx)
	req := &http.Request{}
	if exists {
		headers := map[string][]string{}
		for key, val := range metadata {
			headers[key] = val
		}
		req = &http.Request{
			Header: headers,
		}
	}
	num := 5
	respWriter := mockResponseWriter{blockToReturn: &num}
	bbb.serverCallback(respWriter, req)
	return &tmservice.GetLatestBlockResponse{Block: &types.Block{Header: types.Header{Height: int64(num)}}}, nil
}

// generates a chain parser, a chain fetcher messages based on it
// apiInterface can either be an ApiInterface string as in spectypes.ApiInterfaceXXX or a number for an index in the apiCollections
func CreateChainLibMocks(ctx context.Context, specIndex string, apiInterface string, serverCallback http.HandlerFunc, getToTopMostPath string) (cpar ChainParser, crout ChainRouter, cfetc chaintracker.ChainFetcher, closeServer func(), errRet error) {
	closeServer = nil
	spec, err := keepertest.GetASpec(specIndex, getToTopMostPath, nil, nil)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	index, err := strconv.Atoi(apiInterface)
	if err == nil && index < len(spec.ApiCollections) {
		apiInterface = spec.ApiCollections[index].CollectionData.ApiInterface
	}
	chainParser, err := NewChainParser(apiInterface)
	if err != nil {
		return nil, nil, nil, nil, err
	}
	var chainRouter ChainRouter
	chainParser.SetSpec(spec)
	endpoint := &lavasession.RPCProviderEndpoint{
		NetworkAddress: lavasession.NetworkAddressData{},
		ChainID:        specIndex,
		ApiInterface:   apiInterface,
		Geolocation:    1,
		NodeUrls:       []common.NodeUrl{},
	}
	if apiInterface == spectypes.APIInterfaceGrpc {
		// Start a new gRPC server using the buffered connection
		grpcServer := grpc.NewServer()
		lis, err := net.Listen("tcp", "localhost:0")
		if err != nil {
			return nil, nil, nil, closeServer, err
		}
		endpoint.NodeUrls = append(endpoint.NodeUrls, common.NodeUrl{Url: lis.Addr().String()})
		go func() {
			service := myServiceImplementation{serverCallback: serverCallback}
			tmservice.RegisterServiceServer(grpcServer, service)
			gogoreflection.Register(grpcServer)
			// Serve requests on the buffered connection
			if err := grpcServer.Serve(lis); err != nil {
				return
			}
		}()
		time.Sleep(10 * time.Millisecond)
		chainRouter, err = GetChainRouter(ctx, 1, endpoint, chainParser)
		if err != nil {
			return nil, nil, nil, closeServer, err
		}
	} else {
		mockServer := httptest.NewServer(serverCallback)
		closeServer = mockServer.Close
		endpoint.NodeUrls = append(endpoint.NodeUrls, common.NodeUrl{Url: mockServer.URL})
		chainRouter, err = GetChainRouter(ctx, 1, endpoint, chainParser)
		if err != nil {
			return nil, nil, nil, closeServer, err
		}
	}

	chainFetcher := NewChainFetcher(ctx, chainRouter, chainParser, endpoint)

	return chainParser, chainRouter, chainFetcher, closeServer, err
}

type TestStruct struct {
	Ctx       context.Context
	Keepers   *keepertest.Keepers
	Servers   *keepertest.Servers
	Providers []testcommon.Account
	Spec      spectypes.Spec
	Plan      plantypes.Plan
	Consumer  testcommon.Account
}

func SetupForTests(t *testing.T, numOfProviders int, specID string, getToTopMostPath string) TestStruct {
	ts := TestStruct{}
	ts.Servers, ts.Keepers, ts.Ctx = keepertest.InitAllKeepers(t)
	// init keepers state
	var balance int64 = 100000000000

	// setup consumer
	ts.Consumer = testcommon.CreateNewAccount(ts.Ctx, *ts.Keepers, balance)

	// setup providers
	for i := 0; i < numOfProviders; i++ {
		ts.Providers = append(ts.Providers, testcommon.CreateNewAccount(ts.Ctx, *ts.Keepers, balance))
	}
	sdkContext := sdk.UnwrapSDKContext(ts.Ctx)
	spec, err := keepertest.GetASpec(specID, getToTopMostPath, &sdkContext, &ts.Keepers.Spec)
	if err != nil {
		require.NoError(t, err)
	}
	ts.Spec = spec
	ts.Keepers.Spec.SetSpec(sdk.UnwrapSDKContext(ts.Ctx), ts.Spec)

	ts.Plan = testcommon.CreateMockPlan()
	ts.Keepers.Plans.AddPlan(sdk.UnwrapSDKContext(ts.Ctx), ts.Plan)

	var stake int64 = 50000000000

	// subscribe consumer
	testcommon.BuySubscription(t, ts.Ctx, *ts.Keepers, *ts.Servers, ts.Consumer, ts.Plan.Index)

	// stake providers
	for _, provider := range ts.Providers {
		testcommon.StakeAccount(t, ts.Ctx, *ts.Keepers, *ts.Servers, provider, ts.Spec, stake)
	}

	// advance for the staking to be valid
	ts.Ctx = keepertest.AdvanceEpoch(ts.Ctx, ts.Keepers)
	return ts
}