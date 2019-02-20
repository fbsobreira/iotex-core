//go:binary-only-package
package usbwallet

import (
	_ "context"
	_ "encoding/binary"
	_ "encoding/hex"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "math/big"
	_ "runtime"
	_ "sync"
	_ "time"

	_ "github.com/karalabe/hid"
	_ "github.com/golang/protobuf/proto"

	_ "github.com/ethereum/go-ethereum"
	_ "github.com/ethereum/go-ethereum/accounts"
	_ "github.com/ethereum/go-ethereum/accounts/usbwallet/internal/trezor"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/event"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/rlp"
)

