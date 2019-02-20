//go:binary-only-package
package jsre

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/robertkrimen/otto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/internal/jsre/deps"
)
