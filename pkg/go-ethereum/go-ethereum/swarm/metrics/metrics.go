//go:binary-only-package
package metrics

import (
	_ "time"

	_ "github.com/ethereum/go-ethereum/cmd/utils"
	_ "github.com/ethereum/go-ethereum/metrics"
	_ "github.com/ethereum/go-ethereum/metrics/influxdb"
	_ "github.com/ethereum/go-ethereum/swarm/log"
	_ "gopkg.in/urfave/cli.v1"
)
