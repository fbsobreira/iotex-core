// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package servermonitor

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/iotex-core/protogen/iotexapi"
	"golang.org/x/net/context"
)

func TestServerMonitor(t *testing.T) {
	// TODO - complete it
	// TODO - Run function servermonitor for all given server's list
	// TODO - how would you maintain server's info
	// TODO - how to update height value while blockchain is running

	requireTest := require.New(t)

	// demo code only.. to check it on minicluster
	log.Println("Info: Please run minicluster on your machine first")
	jrpcAddr := "127.0.0.1:14004"

	client, err := New(jrpcAddr)
	msg := fmt.Sprintf("error in connecting to server %s, %v ", jrpcAddr, err)
	requireTest.NoError(err, msg)

	// TODO - how to prepare proper blockchain request
	//blkMetaRequest := &iotexapi.GetBlockMetasRequest{
	//	Lookup: &iotexapi.GetBlockMetasRequest_ByIndex{
	//		ByIndex: &iotexapi.GetBlockMetasByIndexRequest{
	//			Start: 0,
	//			Count: 4,
	//		},
	//	},
	//}
	blkMetaRequest := new(iotexapi.GetBlockMetasRequest)
	result, err1 := client.ServerMonitor(context.Background(), blkMetaRequest, 4)
	msg = fmt.Sprintf("error in getting server health %s, %v ", jrpcAddr, err)
	requireTest.NoError(err1, msg)
	requireTest.EqualValues(true, result)
}
