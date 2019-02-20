//go:binary-only-package
package swap

import (
	_ "context"
	_ "crypto/ecdsa"
	_ "errors"
	_ "fmt"
	_ "math/big"
	_ "os"
	_ "path/filepath"
	_ "sync"
	_ "time"

	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/contracts/chequebook"
	_ "github.com/ethereum/go-ethereum/contracts/chequebook/contract"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/swarm/log"
	_ "github.com/ethereum/go-ethereum/swarm/services/swap/swap"
)
