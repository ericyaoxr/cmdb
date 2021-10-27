package cdb

import (
	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/infraboard/cmdb/pkg/rds"
)

func NewCDBOperater(client *cdb.Client) *CDBOperater {
	return &CDBOperater{
		client: client,
	}
}

type CDBOperater struct {
	client *cdb.Client
}

func (o *CDBOperater) transferSet(items []*cdb.InstanceInfo) *rds.RdsSet {
	set := rds.NewRdsSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *CDBOperater) transferOne(ins *cdb.InstanceInfo) *rds.Rds {
	return nil
}
