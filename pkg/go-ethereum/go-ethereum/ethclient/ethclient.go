//go:binary-only-package
package ethclient

import (
	_ "context"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "math/big"

	_ "github.com/ethereum/go-ethereum"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/rlp"
	_ "github.com/ethereum/go-ethereum/rpc"
)
