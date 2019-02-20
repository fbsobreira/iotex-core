//go:binary-only-package
package contract

import (
	_ "math/big"
	_ "strings"

	_ "github.com/ethereum/go-ethereum"
	_ "github.com/ethereum/go-ethereum/accounts/abi"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/event"
)
