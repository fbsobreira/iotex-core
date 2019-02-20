//go:binary-only-package
package ethstats

import (
	_ "context"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "math/big"
	_ "net"
	_ "regexp"
	_ "runtime"
	_ "strconv"
	_ "strings"
	_ "time"

	_ "golang.org/x/net/websocket"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/mclock"
	_ "github.com/ethereum/go-ethereum/consensus"
	_ "github.com/ethereum/go-ethereum/core"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/eth"
	_ "github.com/ethereum/go-ethereum/event"
	_ "github.com/ethereum/go-ethereum/les"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/rpc"
)
