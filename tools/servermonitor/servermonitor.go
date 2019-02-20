// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// This is a testing tool to check health of servers
// To use, run "make build" and " ./bin/servermonitor"

package main

import (
	"fmt"

	"github.com/iotexproject/iotex-core/protogen/iotexapi"

	"github.com/iotexproject/iotex-core/api"

	_ "go.uber.org/automaxprocs"
)

const (
	adminNumber = 2
)

func servermonitor(svr *api.Server, height uint64) error {
	dummyIn := new(iotexapi.GetBlockMetasRequest)
	resp, err := svr.GetBlockMetas(nil, dummyIn)
	if err != nil {
		fmt.Println("error in server monitoring in GetBlockMetas")
		return err
	}
	for _, blkmeta := range resp.BlkMetas {
		if blkmeta.Height < height {
			fmt.Printf("server %T health is down", svr)
			return err
		}
	}
	return nil
}

func main() {
	// TODO -
}
