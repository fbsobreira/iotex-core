//go:binary-only-package
package build

import (
	_ "archive/tar"
	_ "archive/zip"
	_ "bytes"
	_ "compress/gzip"
	_ "context"
	_ "flag"
	_ "fmt"
	_ "io"
	_ "io/ioutil"
	_ "log"
	_ "net/url"
	_ "os"
	_ "os/exec"
	_ "path"
	_ "path/filepath"
	_ "runtime"
	_ "strings"
	_ "text/template"

	_ "github.com/Azure/azure-storage-blob-go/2018-03-28/azblob"
	_ "golang.org/x/crypto/openpgp"
)
