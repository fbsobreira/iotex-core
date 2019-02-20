//go:binary-only-package
package console

import (
	_ "encoding/json"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "os"
	_ "os/signal"
	_ "path/filepath"
	_ "regexp"
	_ "sort"
	_ "strings"
	_ "syscall"
	_ "time"

	_ "github.com/mattn/go-colorable"
	_ "github.com/peterh/liner"
	_ "github.com/robertkrimen/otto"

	_ "github.com/ethereum/go-ethereum/accounts/usbwallet"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/rpc"
	_ "github.com/ethereum/go-ethereum/internal/jsre"
	_ "github.com/ethereum/go-ethereum/internal/web3ext"
	_ "github.com/ethereum/go-ethereum/rpc"
)

