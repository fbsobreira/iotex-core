//go:binary-only-package
package mru

import (
	_ "bytes"
	_ "context"
	_ "crypto/ecdsa"
	_ "encoding/binary"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "hash"
	_ "path/filepath"
	_ "sync"
	_ "time"
	_ "unsafe"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/p2p/discover"
	_ "github.com/ethereum/go-ethereum/swarm/chunk"
	_ "github.com/ethereum/go-ethereum/swarm/log"
	_ "github.com/ethereum/go-ethereum/swarm/multihash"
	_ "github.com/ethereum/go-ethereum/swarm/storage"
)
