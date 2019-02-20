//go:binary-only-package
package testutil

import (
	_ "io"
	_ "io/ioutil"
	_ "os"
	_ "net/http"
	_ "net/http/httptest"
	_ "strings"
	_ "testing"

	_ "github.com/ethereum/go-ethereum/swarm/api"
	_ "github.com/ethereum/go-ethereum/swarm/storage"
	_ "github.com/ethereum/go-ethereum/swarm/storage/mru"
)
