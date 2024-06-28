package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/test/planet/x/blog/keeper"
	"github.com/test/planet/x/blog/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the post
	for _, elem := range genState.PostList {
		if err := k.Post.Set(ctx, elem.Id, elem); err != nil {
			panic(err)
		}
	}

	// Set post count
	if err := k.PostSeq.Set(ctx, genState.PostCount); err != nil {
		panic(err)
	}
	// Set all the sentPost
	for _, elem := range genState.SentPostList {
		if err := k.SentPost.Set(ctx, elem.Id, elem); err != nil {
			panic(err)
		}
	}

	// Set sentPost count
	if err := k.SentPostSeq.Set(ctx, genState.SentPostCount); err != nil {
		panic(err)
	}
	// Set all the timeoutPost
	for _, elem := range genState.TimeoutPostList {
		if err := k.TimeoutPost.Set(ctx, elem.Id, elem); err != nil {
			panic(err)
		}
	}

	// Set timeoutPost count
	if err := k.TimeoutPostSeq.Set(ctx, genState.TimeoutPostCount); err != nil {
		panic(err)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if k.ShouldBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	if err := k.Params.Set(ctx, genState.Params); err != nil {
		panic(err)
	}
}

// ExportGenesis returns the module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	var err error

	genesis := types.DefaultGenesis()
	genesis.Params, err = k.Params.Get(ctx)
	if err != nil {
		panic(err)
	}

	genesis.PortId = k.GetPort(ctx)

	err = k.Post.Walk(ctx, nil, func(key uint64, elem types.Post) (bool, error) {
		genesis.PostList = append(genesis.PostList, elem)
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	genesis.PostCount, err = k.PostSeq.Peek(ctx)
	if err != nil {
		panic(err)
	}

	err = k.SentPost.Walk(ctx, nil, func(key uint64, elem types.SentPost) (bool, error) {
		genesis.SentPostList = append(genesis.SentPostList, elem)
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	genesis.SentPostCount, err = k.SentPostSeq.Peek(ctx)
	if err != nil {
		panic(err)
	}

	err = k.TimeoutPost.Walk(ctx, nil, func(key uint64, elem types.TimeoutPost) (bool, error) {
		genesis.TimeoutPostList = append(genesis.TimeoutPostList, elem)
		return false, nil
	})
	if err != nil {
		panic(err)
	}

	genesis.TimeoutPostCount, err = k.TimeoutPostSeq.Peek(ctx)
	if err != nil {
		panic(err)
	}

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
