package cdb_test

import (
	"fmt"
	"os"
	"testing"

	op "github.com/ericyaoxr/cmdb/provider/txyun/cdb"
	"github.com/ericyaoxr/cmdb/provider/txyun/connectivity"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
)

var (
	operater *op.CDBOperater
)

func TestQuery(t *testing.T) {
	pager := operater.PageQuery()

	hasNext := true
	for hasNext {
		p := pager.Next()
		hasNext = p.HasNext
		fmt.Println(p.Data)
	}
}

func init() {
	var secretID, secretKey string
	if secretID = os.Getenv("TX_CLOUD_SECRET_ID"); secretID == "" {
		panic("empty TX_CLOUD_SECRET_ID")
	}

	if secretKey = os.Getenv("TX_CLOUD_SECRET_KEY"); secretKey == "" {
		panic("empty TX_CLOUD_SECRET_KEY")
	}

	client := connectivity.NewTencentCloudClient(secretID, secretKey, regions.Guangzhou)
	operater = op.NewCDBOperater(client.CdbClient())
}
