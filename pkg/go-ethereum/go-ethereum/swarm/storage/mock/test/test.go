//go:binary-only-package
package test

import (
	_ "bytes"
	_ "fmt"
	_ "io"
	_ "strconv"
	_ "testing"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/swarm/storage"
	_ "github.com/ethereum/go-ethereum/swarm/storage/mock"
)
