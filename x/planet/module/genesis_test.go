package planet_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/test/planet/testutil/keeper"
	"github.com/test/planet/testutil/nullify"
	planet "github.com/test/planet/x/planet/module"
	"github.com/test/planet/x/planet/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx, _ := keepertest.PlanetKeeper(t)
	planet.InitGenesis(ctx, k, genesisState)
	got := planet.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
