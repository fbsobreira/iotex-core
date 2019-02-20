//go:binary-only-package
package tracers

import (
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "math/big"
	_ "strings"
	_ "sync/atomic"
	_ "time"
	_ "unicode"
	_ "unsafe"

	_ "gopkg.in/olebedev/go-duktape.v3"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/core/vm"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/eth/tracers/internal/tracers"
)
