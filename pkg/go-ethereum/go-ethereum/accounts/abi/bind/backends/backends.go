//go:binary-only-package
package backends

import (
	_ "context"
	_ "errors"
	_ "fmt"
	_ "math/big"
	_ "sync"
	_ "time"

	_ "github.com/ethereum/go-ethereum"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/math"
	_ "github.com/ethereum/go-ethereum/consensus/ethash"
	_ "github.com/ethereum/go-ethereum/core"
	_ "github.com/ethereum/go-ethereum/core/bloombits"
	_ "github.com/ethereum/go-ethereum/core/rawdb"
	_ "github.com/ethereum/go-ethereum/core/state"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/core/vm"
	_ "github.com/ethereum/go-ethereum/eth/filters"
	_ "github.com/ethereum/go-ethereum/ethdb"
	_ "github.com/ethereum/go-ethereum/event"
	_ "github.com/ethereum/go-ethereum/params"
	_ "github.com/ethereum/go-ethereum/rpc"
)
