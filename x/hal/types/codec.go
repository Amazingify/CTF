package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptoCodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/hal interfaces and concrete types on the provided LegacyAmino codec.
// These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgMintHAL{}, "hal/MsgMintHAL", nil)
	cdc.RegisterConcrete(&MsgRedeemCollateral{}, "hal/MsgRedeemCollateral", nil)
}

// RegisterInterfaces registers the x/hal interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintHAL{},
		&MsgRedeemCollateral{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/hal module codec. Note, the codec should ONLY be used in certain instances
	// of tests and for JSON encoding as Amino is still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/hal and defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptoCodec.RegisterCrypto(amino)
	amino.Seal()
}
