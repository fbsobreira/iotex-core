//go:binary-only-package
package core

import (
	_ "bufio"
	_ "bytes"
	_ "context"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io/ioutil"
	_ "math/big"
	_ "os"
	_ "reflect"
	_ "regexp"
	_ "strings"
	_ "sync"

	_ "github.com/davecgh/go-spew/spew"
	_ "golang.org/x/crypto/ssh/terminal"

	_ "github.com/ethereum/go-ethereum/accounts"
	_ "github.com/ethereum/go-ethereum/accounts/abi"
	_ "github.com/ethereum/go-ethereum/accounts/keystore"
	_ "github.com/ethereum/go-ethereum/accounts/usbwallet"
	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/common/hexutil"
	_ "github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/core/types"
	_ "github.com/ethereum/go-ethereum/internal/ethapi"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/rlp"
)
