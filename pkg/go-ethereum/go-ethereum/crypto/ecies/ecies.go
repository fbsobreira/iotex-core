//go:binary-only-package
package ecies

import (
	_ "crypto"
	_ "crypto/aes"
	_ "crypto/cipher"
	_ "crypto/ecdsa"
	_ "crypto/elliptic"
	_ "crypto/hmac"
	_ "crypto/sha256"
	_ "crypto/sha512"
	_ "crypto/subtle"
	_ "fmt"
	_ "hash"
	_ "io"
	_ "math/big"

	_ "github.com/ethereum/go-ethereum/crypto"
)
