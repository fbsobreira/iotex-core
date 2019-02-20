//go:binary-only-package
package client

import (
	_ "archive/tar"
	_ "bytes"
	_ "encoding/json"
	_ "errors"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "mime"
	_ "mime/multipart"
	_ "net/http"
	_ "net/textproto"
	_ "os"
	_ "path/filepath"
	_ "regexp"
	_ "strconv"
	_ "strings"

	_ "github.com/ethereum/go-ethereum/swarm/api"
	_ "github.com/ethereum/go-ethereum/swarm/storage/mru"
)
