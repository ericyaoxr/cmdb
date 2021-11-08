package impl

import (
	"context"
	"fmt"

	"github.com/ericyaoxr/cmdb/app/host"
	"github.com/ericyaoxr/cmdb/app/resource"
	"github.com/ericyaoxr/cmdb/app/secret"
	"github.com/ericyaoxr/cmdb/app/task"
	"github.com/ericyaoxr/cmdb/conf"

	txConn "github.com/ericyaoxr/cmdb/provider/txyun/connectivity"
	cvmOp "github.com/ericyaoxr/cmdb/provider/txyun/cvm"
)

type SyncTaskCallback func(*task.Task)

func (s *service) syncHost(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager host.Pager
	)

	// 处理任务状态
	t.Run()
	defer func() {
		t.Completed()
		cb(t)
	}()

	// 解密secret
	err := secret.DecryptAPISecret(conf.C().App.EncryptKey)
	if err != nil {
		s.log.Warnf("decrypt api secret error, %s", err)
	}

	hs := host.NewHostSet()
	switch secret.Vendor {
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun host ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		operater := cvmOp.NewCVMOperater(client.CvmClient())
		pager = operater.PageQuery()
	// case resource.Vendor_ALIYUN:
	// 	s.log.Debugf("sync aliyun host ...")
	// 	client := aliConn.NewAliCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
	// 	ec, err := client.EcsClient()
	// 	if err != nil {
	// 		t.Failed(err.Error())
	// 		return
	// 	}
	// 	operater := ecsOp.NewEcsOperater(ec)
	// 	req := ecsOp.NewPageQueryRequest()
	// 	req.Rate = int(secret.RequestRate)
	// 	pager = operater.PageQuery(req)
	default:
		t.Failed(fmt.Sprintf("unsuport vendor %s", secret.Vendor))
		return
	}

	// 分页查询数据
	if pager != nil {
		hasNext := true
		for hasNext {
			p := pager.Next()
			hasNext = p.HasNext

			if p.Err != nil {
				t.Failed(fmt.Sprintf("sync error, %s", p.Err))
				return
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
			t.AddDetailFailed(target.Information.Name, err.Error())
		} else {
			t.AddDetailSucceed(target.Information.Name, "")
		}
	}
}
