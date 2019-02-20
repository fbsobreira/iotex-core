//go:binary-only-package
package simulations

import (
	_ "bufio"
	_ "bytes"
	_ "context"
	_ "encoding/json"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "math/rand"
	_ "net/http"
	_ "strconv"
	_ "strings"
	_ "sync"
	_ "time"

	_ "github.com/julienschmidt/httprouter"
	_ "golang.org/x/net/websocket"

	_ "github.com/ethereum/go-ethereum/event"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/p2p/discover"
	_ "github.com/ethereum/go-ethereum/p2p/simulations/adapters"
	_ "github.com/ethereum/go-ethereum/rpc"
)
