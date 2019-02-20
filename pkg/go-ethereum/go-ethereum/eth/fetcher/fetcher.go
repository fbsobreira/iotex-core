//go:binary-only-package
package fetcher

import (
	_ "errors"
	_ "math/rand"
	_ "time"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/prque"
	_ "github.com/ethereum/go-ethereum/consensus"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/metrics"
)
