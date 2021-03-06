package impl

import (
	"context"
	"fmt"

	"github.com/ericyaoxr/cmdb/app/bill"
	"github.com/ericyaoxr/cmdb/app/resource"
	"github.com/ericyaoxr/cmdb/app/secret"
	"github.com/ericyaoxr/cmdb/app/task"
	"github.com/ericyaoxr/cmdb/conf"
	billOp "github.com/ericyaoxr/cmdb/provider/txyun/billing"
	txConn "github.com/ericyaoxr/cmdb/provider/txyun/connectivity"
)

func (s *service) syncBill(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager bill.Pager
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

	switch secret.Vendor {
	case resource.Vendor_TENCENT:
		s.log.Debugf("sync txyun billing ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		operater := billOp.NewBillingOperater(client.BillingClient())
		req := billOp.NewPageQueryRequest()
		req.Rate = int(secret.RequestRate)
		req.Month = "2021-10"
		pager = operater.PageQuery(req)
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

			// 调用bill服务保存数据
			for i := range p.Data.Items {
				target := p.Data.Items[i]
				h, err := s.bill.SaveBill(ctx, target)
				if err != nil {
					s.log.Warnf("save bill error, %s", err)
					t.AddDetailFailed(target.InstanceName, err.Error())
				} else {
					s.log.Debugf("save bill %s to db", h.InstanceName)
					t.AddDetailSucceed(target.InstanceName, "")
				}
			}
		}
	}
}
