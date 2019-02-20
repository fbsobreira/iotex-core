//go:binary-only-package
package fuse

import (
	_ "context"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "os"
	_ "os/exec"
	_ "path/filepath"
	_ "runtime"
	_ "strings"
	_ "sync"
	_ "time"

	_ "bazil.org/fuse"
	_ "bazil.org/fuse/fs"
	_ "golang.org/x/net/context"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/swarm/api"
	_ "github.com/ethereum/go-ethereum/swarm/log"
	_ "github.com/ethereum/go-ethereum/swarm/storage"
)
