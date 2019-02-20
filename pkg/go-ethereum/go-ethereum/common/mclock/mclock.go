//go:binary-only-package
package mclock

import (
	_ "sync"
	_ "time"

	_ "github.com/aristanetworks/goarista/monotime"
)
