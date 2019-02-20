//go:binary-only-package
package http

import (
	_ "bufio"
	_ "bytes"
	_ "context"
	_ "encoding/hex"
	_ "encoding/json"
	_ "fmt"
	_ "html/template"
	_ "io"
	_ "io/ioutil"
	_ "mime"
	_ "mime/multipart"
	_ "net/http"
	_ "os"
	_ "path"
	_ "regexp"
	_ "runtime/debug"
	_ "strconv"
	_ "strings"
	_ "time"

	_ "github.com/pborman/uuid"
	_ "github.com/rs/cors"

	_ "github.com/ethereum/go-ethereum/common"
	_ "github.com/ethereum/go-ethereum/log"
	_ "github.com/ethereum/go-ethereum/metrics"
	_ "github.com/ethereum/go-ethereum/swarm/api"
	_ "github.com/ethereum/go-ethereum/swarm/log"
	_ "github.com/ethereum/go-ethereum/swarm/sctx"
	_ "github.com/ethereum/go-ethereum/swarm/spancontext"
	_ "github.com/ethereum/go-ethereum/swarm/storage"
	_ "github.com/ethereum/go-ethereum/swarm/storage/mru"
)
