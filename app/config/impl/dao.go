package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ericyaoxr/cmdb/app/config"
)

func (s *service) save(ctx context.Context, conf *config.Config) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事务
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
	stmt, err = tx.Prepare(insertConfigSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	desc := conf.Describe
	_, err = stmt.Exec(
		desc.Id, desc.ApplicationId, desc.Name, desc.Host, desc.Port, desc.Env, desc.Type, desc.Source, desc.CreateAt, desc.UpdateAt,
	)
	if err != nil {
		return fmt.Errorf("save config info error, %s", err)
	}

	return tx.Commit()
}

func (s *service) delete(ctx context.Context, req *config.DeleteConfigRequest) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事物
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

	stmt, err = tx.Prepare(deleteConfigSQL)
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
