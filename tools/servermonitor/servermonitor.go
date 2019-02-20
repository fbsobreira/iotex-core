// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

// This is a testing tool to check health of servers
// To use, run "make build" and " ./bin/servermonitor"

package servermonitor

import (
	"fmt"

	"github.com/iotexproject/iotex-core/protogen/iotexapi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Client is the blockchain API client.
type Client struct {
	api iotexapi.APIServiceClient
}

// New creates a new Client.
func New(serverAddr string) (*Client, error) {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &Client{
		api: iotexapi.NewAPIServiceClient(conn),
	}, nil
}

// ServerMonitor is to track server's health by checking it's current height
func (c *Client) ServerMonitor(ctx context.Context, request *iotexapi.GetBlockMetasRequest, height uint64) (bool, error) {
	resp, err := c.api.GetBlockMetas(ctx, request)
	if err != nil {
		return false, fmt.Errorf("\n error in GetBlockMetas: %v", err)
	}
	// blkPb := resp.BlkMetas[0]
	for _, blkPb := range resp.BlkMetas {
		if blkPb.Height < height {
			return false, fmt.Errorf("server health status is negative")
		}
	}
	return true, nil
}
