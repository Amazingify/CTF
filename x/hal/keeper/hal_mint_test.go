package keeper_test

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cosmosutils "github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/gaia/v7/x/hal"
	"github.com/cosmos/gaia/v7/x/hal/keeper"
	testutil "github.com/cosmos/gaia/v7/x/hal/mock_types"
	"github.com/cosmos/gaia/v7/x/hal/types"
	"github.com/golang/mock/gomock"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"
)

func setupKeeperAndContext(t *testing.T) (
	*testutil.MockBankKeeper,
	*testutil.MockAccountKeeper,
	context.Context,
	*keeper.Keeper,
	types.MsgServer,
) {

	encCfg := testutil.MakeTestEncodingConfig()

	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	tkey := sdk.NewTransientStoreKey("transient_test")
	testCtx := cosmosutils.DefaultContext(storeKey, tkey)

	// * --------- setting the db/store starts --------- * //
	registry := codectypes.NewInterfaceRegistry()
	cdc := codec.NewProtoCodec(registry)
	// * --------- setting the db/store ends --------- * //
	ctx := testCtx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
	// * --------- setting the mock starts --------- * //
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	accountMock := testutil.NewMockAccountKeeper(ctrl)
	context := sdk.WrapSDKContext(ctx)
	ps := paramstypes.NewSubspace(cdc, encCfg.Amino, storeKey, tkey, "HalParams")
	// * --------- setting the mock ends --------- * //
	k := keeper.NewKeeper(cdc, storeKey, accountMock, bankMock, ps)
	k.SetParams(ctx, types.DefaultParams())
	server := keeper.NewMsgServerImpl(k)

	return bankMock, accountMock, context, &k, server
}

func TestDummy(t *testing.T) {
	_, _, context, keeper, _ := setupKeeperAndContext(t)
	ctx := sdk.UnwrapSDKContext(context)
	hal.InitGenesis(ctx, *keeper, *types.DefaultGenesisState())
}
