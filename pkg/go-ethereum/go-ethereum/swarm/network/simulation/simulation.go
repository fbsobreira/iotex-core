//go:binary-only-package
package simulation

import (
	_ "context"
	_ "encoding/hex"
	_ "encoding/json"
	_ "errors"
	_ "io/ioutil"
	_ "math/rand"
	_ "net/http"
	_ "os"
	_ "strings"
	_ "sync"
	_ "time"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/node"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/p2p/discover"
	_ "github.com/ethereum/go-ethereum/p2p/simulations"
	_ "github.com/ethereum/go-ethereum/p2p/simulations/adapters"
	_ "github.com/ethereum/go-ethereum/swarm/network"
)
