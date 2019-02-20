//go:binary-only-package
package mem

import (
	_ "archive/tar"
	_ "bytes"
	_ "encoding/json"
	_ "io"
	_ "io/ioutil"
	_ "sync"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/swarm/storage/mock"
)
