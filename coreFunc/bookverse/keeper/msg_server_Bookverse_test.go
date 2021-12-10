package keeper

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/goldipatel12/marketplace/x/Bookverses/types"
)

func TestBookverseMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateBookverse(ctx, &types.MsgCreateBookverse{Creator: creator})
		require.NoError(t, err)
		assert.Equal(t, i, int(resp.Id))
	}
}

func TestBookverseMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateBookverse
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateBookverse{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBookverse{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateBookverse{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateBookverse(ctx, &types.MsgCreateBookverse{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateBookverse(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestBookverseMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteBookverse
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteBookverse{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteBookverse{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteBookverse{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		tc := tc
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateBookverse(ctx, &types.MsgCreateBookverse{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteBookverse(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
