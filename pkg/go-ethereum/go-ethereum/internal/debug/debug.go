//go:binary-only-package
package debug

import (
	_ "bytes"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "net/http"
	_ "net/http/pprof"
	_ "os"
	_ "os/user"
	_ "path/filepath"
	_ "runtime"
	_ "runtime/debug"
	_ "runtime/pprof"
	_ "runtime/trace"
	_ "strings"
	_ "sync"
	_ "time"

	_ "github.com/fjl/memsize/memsizeui"
	_ "github.com/mattn/go-colorable"
	_ "gopkg.in/urfave/cli.v1"

	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/log/term"
	_ "github.com/ethereum/go-ethereum/metrics"
	_ "github.com/ethereum/go-ethereum/metrics/exp"
)
