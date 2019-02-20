//go:binary-only-package
package consensus

import (
	_ "math/big"
	_ "errors"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/state"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/params"
	_ "github.com/ethereum/go-ethereum/rpc"
)
