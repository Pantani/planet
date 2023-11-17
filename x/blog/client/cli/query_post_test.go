package cli_test

import (
	"fmt"
	"testing"

	tmcli "github.com/cometbft/cometbft/libs/cli"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"planet/testutil/network"
	"planet/testutil/nullify"
	"planet/x/blog/client/cli"
	"planet/x/blog/types"
)

func (s *IntegrationTestSuite) networkWithPostObjects(n int) (*network.Network, []types.Post) {
	s.T().Helper()
	state := types.GenesisState{PortId: types.PortID}
	for i := 0; i < n; i++ {
		post := types.Post{
			Id: uint64(i),
		}
		nullify.Fill(&post)
		state.PostList = append(state.PostList, post)
	}
	return s.network(&state), state.PostList
}

func (s *IntegrationTestSuite) TestShowPost() {
	var (
		net, objs = s.networkWithPostObjects(2)
		ctx       = net.Validators[0].ClientCtx
		common    = []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
	)
	tests := []struct {
		desc string
		id   string
		args []string
		err  error
		obj  types.Post
	}{
		{
			desc: "found",
			id:   fmt.Sprintf("%d", objs[0].Id),
			args: common,
			obj:  objs[0],
		},
		{
			desc: "not found",
			id:   "not_found",
			args: common,
			err:  status.Error(codes.NotFound, "not found"),
		},
	}
	for _, tc := range tests {
		s.T().Run(tc.desc, func(t *testing.T) {
			args := append([]string{tc.id}, tc.args...)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdShowPost(), args)
			if tc.err != nil {
				stat, ok := status.FromError(tc.err)
				require.True(t, ok)
				require.ErrorIs(t, stat.Err(), tc.err)
				return
			}
			require.NoError(t, err)
			var resp types.QueryGetPostResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.NotNil(t, resp.Post)
			require.Equal(t,
				nullify.Fill(&tc.obj),
				nullify.Fill(&resp.Post),
			)
		})
	}
}

func (s *IntegrationTestSuite) TestListPost() {
	var (
		net, objs = s.networkWithPostObjects(5)
		ctx       = net.Validators[0].ClientCtx
	)
	request := func(next []byte, offset, limit uint64, total bool) []string {
		args := []string{
			fmt.Sprintf("--%s=json", tmcli.OutputFlag),
		}
		if next == nil {
			args = append(args, fmt.Sprintf("--%s=%d", flags.FlagOffset, offset))
		} else {
			args = append(args, fmt.Sprintf("--%s=%s", flags.FlagPageKey, next))
		}
		args = append(args, fmt.Sprintf("--%s=%d", flags.FlagLimit, limit))
		if total {
			args = append(args, fmt.Sprintf("--%s", flags.FlagCountTotal))
		}
		return args
	}
	s.T().Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(objs); i += step {
			args := request(nil, uint64(i), uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPost(), args)
			require.NoError(t, err)
			var resp types.QueryAllPostResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Post), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Post),
			)
		}
	})
	s.T().Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(objs); i += step {
			args := request(next, 0, uint64(step), false)
			out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPost(), args)
			require.NoError(t, err)
			var resp types.QueryAllPostResponse
			require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
			require.LessOrEqual(t, len(resp.Post), step)
			require.Subset(t,
				nullify.Fill(objs),
				nullify.Fill(resp.Post),
			)
			next = resp.Pagination.NextKey
		}
	})
	s.T().Run("Total", func(t *testing.T) {
		args := request(nil, 0, uint64(len(objs)), true)
		out, err := clitestutil.ExecTestCLICmd(ctx, cli.CmdListPost(), args)
		require.NoError(t, err)
		var resp types.QueryAllPostResponse
		require.NoError(t, net.Config.Codec.UnmarshalJSON(out.Bytes(), &resp))
		require.NoError(t, err)
		require.Equal(t, len(objs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(objs),
			nullify.Fill(resp.Post),
		)
	})
}
