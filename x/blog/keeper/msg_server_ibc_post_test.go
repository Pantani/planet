package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"

	keepertest "github.com/test/planet/testutil/keeper"
	"github.com/test/planet/x/blog/keeper"
	"github.com/test/planet/x/blog/types"
)

func TestMsgServerSendIbcPost(t *testing.T) {
	k, ctx, addressCodec := keepertest.BlogKeeper(t)
	creator, err := addressCodec.BytesToString([]byte("signerAddr__________________"))
	require.NoError(t, err)
	srv := keeper.NewMsgServerImpl(k)

	tests := []struct {
		name string
		msg  types.MsgSendIbcPost
		err  error
	}{
		{
			name: "invalid address",
			msg: types.MsgSendIbcPost{
				Creator:          "invalid address",
				Port:             "port",
				ChannelID:        "channel-0",
				TimeoutTimestamp: 100,
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "invalid port",
			msg: types.MsgSendIbcPost{
				Creator:          creator,
				Port:             "",
				ChannelID:        "channel-0",
				TimeoutTimestamp: 100,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid channel",
			msg: types.MsgSendIbcPost{
				Creator:          creator,
				Port:             "port",
				ChannelID:        "",
				TimeoutTimestamp: 100,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "invalid timeout",
			msg: types.MsgSendIbcPost{
				Creator:          creator,
				Port:             "port",
				ChannelID:        "channel-0",
				TimeoutTimestamp: 0,
			},
			err: sdkerrors.ErrInvalidRequest,
		}, {
			name: "valid message",
			msg: types.MsgSendIbcPost{
				Creator:          creator,
				Port:             "port",
				ChannelID:        "channel-0",
				TimeoutTimestamp: 100,
			},
			err: channeltypes.ErrChannelCapabilityNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = srv.SendIbcPost(ctx, &tt.msg)
			if tt.err != nil {
				require.ErrorContains(t, err, tt.err.Error())
				return
			}
			require.NoError(t, err)
		})
	}
}
