package cdb

import (
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/ericyaoxr/cmdb/app/rds"
	"github.com/ericyaoxr/cmdb/utils"
)

func (o *CDBOperater) Query(req *cdb.DescribeDBInstancesRequest) (*rds.RdsSet, error) {
	resp, err := o.client.DescribeDBInstances(req)
	if err != nil {
		return nil, err
	}

	set := o.transferSet(resp.Response.Items)
	set.Total = utils.PtrInt64(resp.Response.TotalCount)

	return set, nil
}

func (o *CDBOperater) PageQuery() rds.Pager {
	return newPager(20, o)
}
