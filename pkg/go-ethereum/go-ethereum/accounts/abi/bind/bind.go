//go:binary-only-package
package bind

import (
	_ "bytes"
	_ "context"
	_ "crypto/ecdsa"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "math/big"
	_ "reflect"
	_ "regexp"
	_ "strings"
	_ "text/template"
	_ "time"
	_ "unicode"

	_ "golang.org/x/tools/imports"

	_ "github.com/ethereum/go-ethereum"
	_ "github.com/ethereum/go-ethereum/accounts/abi"
	_ "github.com/ethereum/go-ethereum/accounts/keystore"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/event"
	_ "github.com/ethereum/go-ethereum/log"
)

