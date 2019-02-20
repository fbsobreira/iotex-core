//go:binary-only-package
package client

import (
	_ "context"
	_ "errors"
	_ "fmt"
	_ "sync"
	_ "time"

	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/p2p/discover"
	_ "github.com/ethereum/go-ethereum/p2p/protocols"
	_ "github.com/ethereum/go-ethereum/rlp"
	_ "github.com/ethereum/go-ethereum/rpc"
	_ "github.com/ethereum/go-ethereum/swarm/log"
	_ "github.com/ethereum/go-ethereum/swarm/pss"
)
