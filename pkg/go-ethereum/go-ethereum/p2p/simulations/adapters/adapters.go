//go:binary-only-package
package adapters

import (
	_ "bufio"
	_ "context"
	_ "crypto/ecdsa"
	_ "encoding/hex"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "math"
	_ "net"
	_ "os"
	_ "os/exec"
	_ "os/signal"
	_ "path/filepath"
	_ "regexp"
	_ "runtime"
	_ "strconv"
	_ "strings"
	_ "sync"
	_ "syscall"
	_ "time"

	_ "github.com/docker/docker/pkg/reexec"
	_ "golang.org/x/net/websocket"

	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/event"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/node"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/p2p/discover"
	_ "github.com/ethereum/go-ethereum/p2p/simulations/pipes"
	_ "github.com/ethereum/go-ethereum/rpc"
)
