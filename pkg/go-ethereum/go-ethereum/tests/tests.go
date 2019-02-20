//go:binary-only-package
package tests

import (
	_ "bytes"
	_ "encoding/hex"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "math/big"
	_ "strings"
	_ "testing"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/common/math"
	_ "github.com/ethereum/go-ethereum/consensus/ethash"
	_ "github.com/ethereum/go-ethereum/core"
	_ "github.com/ethereum/go-ethereum/core/state"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/core/vm"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/crypto/sha3"
	_ "github.com/ethereum/go-ethereum/ethdb"
	_ "github.com/ethereum/go-ethereum/params"
	_ "github.com/ethereum/go-ethereum/rlp"
)
