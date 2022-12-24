package hal

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gaia/v7/x/hal/keeper"
	"github.com/cosmos/gaia/v7/x/hal/types"
)

// InitGenesis performs module's genesis initialization.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)

	for _, entry := range genState.RedeemEntries {
		accAddr, err := sdk.AccAddressFromBech32(entry.Address)
		if err != nil {
			panic(fmt.Errorf("parsing RedeemEntry.Address (%s): %w", entry.Address, err))
		}

		k.SetRedeemEntry(ctx, entry)
		for _, op := range entry.Operations {
			k.InsertToRedeemQueue(ctx, op.CompletionTime, accAddr)
		}
	}
}

// ExportGenesis returns the current module genesis state.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	params := k.GetParams(ctx)

	var entries []types.RedeemEntry
	k.IterateRedeemEntries(ctx, func(entry types.RedeemEntry) (stop bool) {
		entries = append(entries, entry)
		return false
	})

	return &types.GenesisState{
		Params:        params,
		RedeemEntries: entries,
	}
}
