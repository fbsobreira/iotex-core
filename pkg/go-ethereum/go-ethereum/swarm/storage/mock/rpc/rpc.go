//go:binary-only-package
package rpc

import (
	_ "fmt"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/rpc"
	_ "github.com/ethereum/go-ethereum/swarm/log"
	_ "github.com/ethereum/go-ethereum/swarm/storage/mock"
)
