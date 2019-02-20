//go:binary-only-package
package runtime

import (
	_ "math"
	_ "math/big"
	_ "time"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core"
	_ "github.com/ethereum/go-ethereum/core/state"
	_ "github.com/ethereum/go-ethereum/core/vm"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/ethdb"
	_ "github.com/ethereum/go-ethereum/params"
)

