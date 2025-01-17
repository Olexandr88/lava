package subscription_test

import (
	"testing"

	keepertest "github.com/lavanet/lava/v4/testutil/keeper"
	"github.com/lavanet/lava/v4/testutil/nullify"
	"github.com/lavanet/lava/v4/x/subscription"
	"github.com/lavanet/lava/v4/x/subscription/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SubscriptionKeeper(t)
	subscription.InitGenesis(ctx, *k, genesisState)
	got := subscription.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
