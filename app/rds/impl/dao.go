package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ericyaoxr/cmdb/app/rds"
)

func (s *service) save(ctx context.Context, h *rds.Rds) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事务
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	// wiki: https://en.wikipedia.org/wiki/Isolation_(database_systems)#Isolation_levels
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(insertResourceSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 生成描写信息的Hash
	if err := h.GenHash(); err != nil {
		return err
	}

	base := h.Base
	info := h.Information
	_, err = stmt.Exec(
		base.Id, base.Vendor, base.Region, base.Zone, base.CreateAt, info.ExpireAt, info.Category, info.Type, base.InstanceId,
		info.Name, info.Description, info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount, info.PublicIPToString(),
		info.PrivateIPToString(), info.PayType, base.DescribeHash, base.ResourceHash,
	)
	if err != nil {
		return fmt.Errorf("save rds resource info error, %s", err)
	}

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(insertRdsSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	desc := h.Describe
	_, err = stmt.Exec(
		desc.ResourceId, desc.Cpu, desc.Memory, desc.Volume, desc.EngineVersion, desc.InstanceNodes, desc.InstanceType,
		desc.Port, desc.Master, desc.Slave,
		desc.DeployMode, desc.Qps, desc.WanPort,
	)
	if err != nil {
		return fmt.Errorf("save rds resource describe error, %s", err)
	}

	return tx.Commit()
}

func (s *service) delete(ctx context.Context, req *rds.DeleteRdsRequest) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事物
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
	// 关于事物级别可以参考文章: https://zhuanlan.zhihu.com/p/117476959
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		fmt.Printf("start tx error, %s", err)
		return err
	}

	// 执行结果提交或者回滚事务
	// 当使用sql.Tx的操作方式操作数据后，需要我们使用sql.Tx的Commit()方法显式地提交事务，
	// 如果出错，则可以使用sql.Tx中的Rollback()方法回滚事务，保持数据的一致性
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	stmt, err = tx.Prepare(deleteRdsSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Id)
	if err != nil {
		return err
	}

	stmt, err = s.db.Prepare(deleteResourceSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
