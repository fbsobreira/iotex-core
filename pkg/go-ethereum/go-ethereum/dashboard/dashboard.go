//go:binary-only-package
package dashboard

import (
	_ "bytes"
	_ "crypto/sha256"
	_ "encoding/json"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "net"
	_ "net/http"
	_ "os"
	_ "path/filepath"
	_ "regexp"
	_ "runtime"
	_ "sort"
	_ "strings"
	_ "sync"
	_ "sync/atomic"
	_ "syscall"
	_ "time"

	_ "github.com/elastic/gosigar"
	_ "github.com/mohae/deepcopy"
	_ "github.com/rjeczalik/notify"
	_ "golang.org/x/net/websocket"

	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/metrics"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/params"
	_ "github.com/ethereum/go-ethereum/rpc"
)
