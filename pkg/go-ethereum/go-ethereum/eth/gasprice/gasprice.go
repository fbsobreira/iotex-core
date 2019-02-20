//go:binary-only-package
package gasprice

import (
	_ "context"
	_ "math/big"
	_ "sort"
	_ "sync"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/internal/ethapi"
	_ "github.com/ethereum/go-ethereum/params"
	_ "github.com/ethereum/go-ethereum/rpc"
)
