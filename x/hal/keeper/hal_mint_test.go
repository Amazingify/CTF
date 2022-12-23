package keeper_test

import (
	"testing"

	"github.com/cosmos/gaia/v7/x/hal/mock_types"
)

func TestAdd(t *testing.T) {

	c := mock_types.MockAccountKeeper{}
	c.GetModuleAddress("hal")

	got := 3
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
