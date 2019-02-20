//go:binary-only-package
package shhclient

import (
	_ "context"

	_ "github.com/ethereum/go-ethereum"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/rpc"
	_ "github.com/ethereum/go-ethereum/whisper/whisperv6"
)
