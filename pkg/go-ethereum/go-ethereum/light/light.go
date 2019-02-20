//go:binary-only-package
package light

import (
	_ "bytes"
	_ "context"
	_ "errors"
	_ "fmt"
	_ "math/big"
	_ "sync"
	_ "sync/atomic"
	_ "time"

	_ "github.com/hashicorp/golang-lru"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/consensus"
	_ "github.com/ethereum/go-ethereum/core"
	_ "github.com/ethereum/go-ethereum/core/rawdb"
	_ "github.com/ethereum/go-ethereum/core/state"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/ethdb"
	_ "github.com/ethereum/go-ethereum/event"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/params"
	_ "github.com/ethereum/go-ethereum/rlp"
	_ "github.com/ethereum/go-ethereum/trie"
)
