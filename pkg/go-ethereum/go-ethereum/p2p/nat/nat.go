//go:binary-only-package
package nat

import (
	_ "errors"
	_ "fmt"
	_ "net"
	_ "strings"
	_ "sync"
	_ "time"

	_ "github.com/huin/goupnp"
	_ "github.com/huin/goupnp/dcps/internetgateway1"
	_ "github.com/huin/goupnp/dcps/internetgateway2"
	_ "github.com/jackpal/go-nat-pmp"

	_ "github.com/ethereum/go-ethereum/log"
)
