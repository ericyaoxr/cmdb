package cdb

import (
	"time"

	cdb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cdb/v20170320"

	"github.com/ericyaoxr/cmdb/app/rds"
	"github.com/ericyaoxr/cmdb/app/resource"
	"github.com/ericyaoxr/cmdb/utils"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"
)

func NewCDBOperater(client *cdb.Client) *CDBOperater {
	return &CDBOperater{
		client: client,
		log:    zap.L().Named("Tx CDB"),
	}
}

type CDBOperater struct {
	client *cdb.Client
	log    logger.Logger
}

func (o *CDBOperater) transferSet(items []*cdb.InstanceInfo) *rds.RdsSet {
	set := rds.NewRdsSet()
	for i := range items {
		set.Add(o.transferOne(items[i]))
	}
	return set
}

func (o *CDBOperater) transferOne(ins *cdb.InstanceInfo) *rds.Rds {
	r := rds.NewDefaultRds()
	r.Base.Vendor = resource.Vendor_TENCENT
	r.Base.Region = o.client.GetRegion()
	r.Base.Zone = utils.PtrStrV(ins.Zone)
	r.Base.CreateAt = o.parseTime(utils.PtrStrV(ins.CreateTime))
	r.Base.InstanceId = utils.PtrStrV(ins.InstanceId)

	r.Information.ExpireAt = o.parseTime(utils.PtrStrV(ins.DeadlineTime))
	r.Information.Type = utils.PtrStrV(ins.DeviceType)
	r.Information.Name = utils.PtrStrV(ins.InstanceName)
	r.Information.Status = rds.Status_name[int32(*ins.Status)]
	// r.Information.Tags = utils.PtrStrV(ins.TagList)
	r.Information.PublicIp = utils.StringPtrStrv(ins.WanDomain)
	r.Information.PrivateIp = utils.StringPtrStrv(ins.Vip)
	r.Information.PayType = rds.PayType_name[int32(*ins.PayType)]

	r.Describe.Cpu = utils.PtrInt64(ins.Cpu)
	r.Describe.Memory = utils.PtrInt64(ins.Memory)
	r.Describe.Volume = utils.PtrInt64(ins.Volume)
	r.Describe.EngineVersion = utils.PtrStrV(ins.EngineVersion)
	r.Describe.InstanceNodes = utils.PtrInt64(ins.InstanceNodes)
	r.Describe.InstanceType = utils.PtrInt64(ins.InstanceType)
	r.Describe.Port = utils.PtrInt64(ins.Vport)
	if ins.MasterInfo != nil {
		r.Describe.Master = utils.PtrStrV(ins.MasterInfo.InstanceId)
	}
	if ins.SlaveInfo != nil {
		r.Describe.Slave = utils.PtrStrV(ins.SlaveInfo.First.Vip)
	}
	r.Describe.DeployMode = utils.PtrInt64(ins.DeployMode)
	r.Describe.Qps = utils.PtrInt64(ins.Qps)
	r.Describe.WanPort = utils.PtrInt64(ins.WanPort)

	return r
}

func transferTags(tags []*cdb.TagInfoItem) map[string]string {
	return nil
}

func (o *CDBOperater) parseTime(t string) int64 {
	if t != "" {
		ts, err := time.Parse("2006-01-02 15:04:05", t)
		if err != nil {
			o.log.Errorf("parse time %s error, %s", t, err)
			return 0
		}

		return ts.UnixNano() / 1000000
	} else {
		return 0
	}
}
