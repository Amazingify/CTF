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
	"github.com/stretchr/testify/require"
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

	// * --------- setting the mock starts --------- * //
	ctx := testCtx.WithBlockHeader(tmproto.Header{Time: tmtime.Now()})
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

func SetBankMocks(bank *testutil.MockBankKeeper) {
	bank.EXPECT().MintCoins(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	bank.EXPECT().SendCoinsFromAccountToModule(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	bank.EXPECT().SendCoinsFromModuleToAccount(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	bank.EXPECT().GetSupply(gomock.Any(), gomock.Any())
	bank.EXPECT().GetAllBalances(gomock.Any(), gomock.Any())
}

func TestConfigurations(t *testing.T) {
	_, _, context, keeper, _ := setupKeeperAndContext(t)
	ctx := sdk.UnwrapSDKContext(context)
	hal.InitGenesis(ctx, *keeper, *types.DefaultGenesisState())

	var collateralMeta []types.TokenMeta
	collateralMeta = append(collateralMeta, types.TokenMeta{
		Denom:       "coin",
		Decimals:    6,
		Description: "",
	})

	keeper.SetParams(ctx, types.NewParams(10000, 10, collateralMeta, types.TokenMeta{Denom: "hal", Decimals: 6, Description: ""}))

	halMeta := keeper.HALMeta(ctx)
	require.EqualValues(t, types.TokenMeta{
		Denom: "hal", Decimals: 6, Description: "",
	}, halMeta)

	params := keeper.GetParams(ctx)

	require.EqualValues(t, types.Params{
		RedeemDur:        10000,
		MaxRedeemEntries: 10,
		CollateralMetas:  collateralMeta,
		HalMeta:          types.TokenMeta{Denom: "hal", Decimals: 6, Description: ""},
	}, params)

}
