package pkg

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// ParseAccAddressArg is a helper function to parse an account address CLI argument.
func ParseAccAddressArg(argName, argValue string) (sdk.AccAddress, error) {
	addr, err := sdk.AccAddressFromBech32(argValue)
	if err != nil {
		return sdk.AccAddress{}, fmt.Errorf("parsing %s argument: invalid address: %w", argName, err)
	}

	return addr, nil
}

// ParseUint64Arg is a helper function to parse uint64 CLI argument.
func ParseUint64Arg(argName, argValue string) (uint64, error) {
	v, err := strconv.ParseUint(argValue, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parsing %s argument: invalid uint64 value: %w", argName, err)
	}

	return v, nil
}
