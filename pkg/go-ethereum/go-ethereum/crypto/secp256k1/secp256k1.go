//go:binary-only-package
package secp256k1

import "C"

import (
	_ "crypto/elliptic"
	_ "errors"
	_ "math/big"
	_ "unsafe"
)
