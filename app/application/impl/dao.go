package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ericyaoxr/cmdb/app/application"
)

func (s *service) save(ctx context.Context, ins *application.Application) error {
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
	stmt, err = tx.Prepare(insertApplicationSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	app := ins.Base
	_, err = stmt.Exec(
		&app.Id, &app.Name, &app.Repo, &app.Branch, &app.Module, &app.Topic, &app.Job,
		&app.Description, &app.CreateAt, &app.UpdateAt, &app.Status,
	)
	if err != nil {
		return fmt.Errorf("save application resource info error, %s", err)
	}

	return tx.Commit()
}

func (s *service) delete(ctx context.Context, req *application.DeleteApplicationRequest) error {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 开启一个事物
	// 文档请参考: http://cngolib.com/database-sql.html#db-begintx
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

	stmt, err = tx.Prepare(deleteApplicationSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Id)
	if err != nil {
		return err
	}

	stmt, err = s.db.Prepare(deleteApplicationSQL)
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
