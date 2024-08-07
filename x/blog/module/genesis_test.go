package blog_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/test/planet/testutil/keeper"
	"github.com/test/planet/testutil/nullify"
	blog "github.com/test/planet/x/blog/module"
	"github.com/test/planet/x/blog/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		PostList: []types.Post{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		PostCount: 2,
		SentPostList: []types.SentPost{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SentPostCount: 2,
		TimeoutPostList: []types.TimeoutPost{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TimeoutPostCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.BlogKeeper(t)
	blog.InitGenesis(ctx, k, genesisState)
	got := blog.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.PostList, got.PostList)
	require.Equal(t, genesisState.PostCount, got.PostCount)
	require.ElementsMatch(t, genesisState.SentPostList, got.SentPostList)
	require.Equal(t, genesisState.SentPostCount, got.SentPostCount)
	require.ElementsMatch(t, genesisState.TimeoutPostList, got.TimeoutPostList)
	require.Equal(t, genesisState.TimeoutPostCount, got.TimeoutPostCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
