//go:binary-only-package
package network

import (
	_ "bytes"
	_ "context"
	_ "errors"
	_ "fmt"
	_ "math/rand"
	_ "net"
	_ "strings"
	_ "sync"
	_ "time"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/p2p/discover"
	_ "github.com/ethereum/go-ethereum/p2p/protocols"
	_ "github.com/ethereum/go-ethereum/rpc"
	_ "github.com/ethereum/go-ethereum/swarm/log"
	_ "github.com/ethereum/go-ethereum/swarm/pot"
	_ "github.com/ethereum/go-ethereum/swarm/state"
	_ "github.com/ethereum/go-ethereum/swarm/storage"
)
