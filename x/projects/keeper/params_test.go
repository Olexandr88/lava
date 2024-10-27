package keeper_test

import (
	"testing"

	testkeeper "github.com/lavanet/lava/v4/testutil/keeper"
	"github.com/lavanet/lava/v4/x/projects/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ProjectsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
