//go:binary-only-package
package rules

import (
	_ "encoding/json"
	_ "fmt"
	_ "os"
	_ "strings"

	_ "github.com/robertkrimen/otto"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/internal/ethapi"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/signer/core"
	_ "github.com/ethereum/go-ethereum/signer/rules/deps"
	_ "github.com/ethereum/go-ethereum/signer/storage"
)
