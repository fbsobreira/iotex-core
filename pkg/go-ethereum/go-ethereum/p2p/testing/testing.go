//go:binary-only-package
package testing

import (
	_ "bytes"
	_ "fmt"
	_ "errors"
	_ "io"
	_ "io/ioutil"
	_ "strings"
	_ "sync"
	_ "testing"
	_ "time"

	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/node"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/p2p/discover"
	_ "github.com/ethereum/go-ethereum/p2p/simulations"
	_ "github.com/ethereum/go-ethereum/p2p/simulations/adapters"
	_ "github.com/ethereum/go-ethereum/rlp"
	_ "github.com/ethereum/go-ethereum/rpc"
)
