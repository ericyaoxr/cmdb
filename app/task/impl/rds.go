package impl

import (
	"context"
	"fmt"

	"github.com/ericyaoxr/cmdb/app/rds"
	"github.com/ericyaoxr/cmdb/app/resource"
	"github.com/ericyaoxr/cmdb/app/secret"
	"github.com/ericyaoxr/cmdb/app/task"
	"github.com/ericyaoxr/cmdb/conf"

	cdbOp "github.com/ericyaoxr/cmdb/provider/txyun/cdb"
	txConn "github.com/ericyaoxr/cmdb/provider/txyun/connectivity"
)

func (s *service) syncRds(ctx context.Context, secret *secret.Secret, t *task.Task, cb SyncTaskCallback) {
	var (
		pager rds.Pager
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
		s.log.Debugf("sync txyun cdb ...")
		client := txConn.NewTencentCloudClient(secret.ApiKey, secret.ApiSecret, t.Region)
		operater := cdbOp.NewCDBOperater(client.CdbClient())
		pager = operater.PageQuery()
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

			// 调用rds服务保存数据
			for i := range p.Data.Items {
				target := p.Data.Items[i]
				h, err := s.rds.SaveRds(ctx, target)
				if err != nil {
					s.log.Warnf("save rds error, %s", err)
					t.AddDetailFailed(target.Information.Name, err.Error())
				} else {
					s.log.Debugf("save rds %s to db", h.ShortDesc())
					t.AddDetailSucceed(target.Information.Name, "")
				}
			}
		}
	}
}
