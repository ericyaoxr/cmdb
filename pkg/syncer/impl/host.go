package impl

import (
	"context"
	"fmt"

	"github.com/ericyaoxr/cmdb/conf"
	"github.com/ericyaoxr/cmdb/pkg/host"
	"github.com/ericyaoxr/cmdb/pkg/resource"
	"github.com/ericyaoxr/cmdb/pkg/syncer"
	"github.com/ericyaoxr/mcube/exception"

	aliConn "github.com/ericyaoxr/cmdb/provider/aliyun/connectivity"
	ecsOp "github.com/ericyaoxr/cmdb/provider/aliyun/ecs"
	txConn "github.com/ericyaoxr/cmdb/provider/txyun/connectivity"
	cvmOp "github.com/ericyaoxr/cmdb/provider/txyun/cvm"
)

func (s *service) syncHost(ctx context.Context, secret *syncer.Secret, region string) (
	*syncer.SyncReponse, error) {
	var (
		pager host.Pager
	)

	// 解密secret
	err := secret.DecryptAPISecret(conf.C().App.EncryptKey)
	if err != nil {
		s.log.Warnf("decrypt api secret error, %s", err)
	}

	hs := host.NewHostSet()
	switch secret.Vendor {
	case resource.VendorAliYun:
		s.log.Debugf("sync aliyun host ...")
		client := aliConn.NewAliCloudClient(secret.APIKey, secret.APISecret, region)
		ec, err := client.EcsClient()
		if err != nil {
			return nil, err
		}
		operater := ecsOp.NewEcsOperater(ec)
		req := ecsOp.NewPageQueryRequest()
		req.Rate = secret.RequestRate
		pager = operater.PageQuery(req)
	case resource.VendorTencent:
		s.log.Debugf("sync txyun host ...")
		client := txConn.NewTencentCloudClient(secret.APIKey, secret.APISecret, region)
		operater := cvmOp.NewCVMOperater(client.CvmClient())
		pager = operater.PageQuery()
	default:
		return nil, exception.NewBadRequest("unsuport vendor %s", secret.Vendor)
	}

	set := syncer.NewSyncReponse()
	// 分页查询数据
	if pager != nil {
		hasNext := true
		for hasNext {
			p := pager.Next()
			hasNext = p.HasNext

			if p.Err != nil {
				return nil, fmt.Errorf("sync error, %s", p.Err)
			}

			// 调用host服务保持数据
			for i := range p.Data.Items {
				hs.Add(p.Data.Items[i])
			}
		}
	}

	// 调用host服务保持数据
	for i := range hs.Items {
		target := hs.Items[i]
		_, err := s.host.SaveHost(ctx, target)
		if err != nil {
			set.AddFailed(target.Name, err.Error())
		} else {
			set.AddSucceed(target.Name, "")
		}
	}

	return set, nil
}
