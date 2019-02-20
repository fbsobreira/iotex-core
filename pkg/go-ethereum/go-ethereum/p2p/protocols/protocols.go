//go:binary-only-package
package protocols

import (
	_ "bufio"
	_ "bytes"
	_ "context"
	_ "fmt"
	_ "io"
	_ "reflect"
	_ "sync"
	_ "time"

	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/metrics"
	_ "github.com/ethereum/go-ethereum/p2p"
	_ "github.com/ethereum/go-ethereum/rlp"
	_ "github.com/ethereum/go-ethereum/swarm/spancontext"
	_ "github.com/ethereum/go-ethereum/swarm/tracing"
	_ "github.com/opentracing/opentracing-go"
)
