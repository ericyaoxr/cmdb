package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ericyaoxr/mcube/exception"
	"github.com/ericyaoxr/mcube/sqlbuilder"
	"github.com/ericyaoxr/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/ericyaoxr/cmdb/app/config"
)

const (
	insertConfigSQL = `INSERT INTO config (
		id,application_id,name,host,port,env,type,source,create_at,update_at
	) VALUES (?,?,?,?,?,?,?,?,?,?);`
	updateConfigSQL = `UPDATE config SET 
		application_id=?,name=?,host=?,port=?,env=?,type=?,source=?,create_at=?,update_at=?
	WHERE id = ?`

	queryConfigSQL  = `SELECT c.id,c.application_id,c.name,c.host,c.port,c.env,c.type,c.source,c.create_at,c.update_at,a.name,a.repo, a.branch, a.module,a.topic, a.job, a.description, a.status FROM config as c LEFT JOIN application as a ON c.application_id=a.id`
	deleteConfigSQL = `DELETE FROM config WHERE id = ?;`
)

func (s *service) SaveConfig(ctx context.Context, conf *config.Config) (
	*config.Config, error) {
	desc := conf.Describe
	desc.Id = xid.New().String()
	desc.CreateAt = ftime.Now().Timestamp()
	desc.UpdateAt = ftime.Now().Timestamp()

	if err := s.save(ctx, conf); err != nil {
		return nil, err
	}

	return conf, nil
}

func (s *service) QueryConfig(ctx context.Context, req *config.QueryConfigRequest) (
	*config.ConfigSet, error) {
	query := sqlbuilder.NewQuery(queryConfigSQL)

	if req.Keywords != "" {
		query.Where("c.name LIKE ? OR c.id = ? OR c.env = ? OR c.type LIKE ? OR c.host LIKE ? OR a.name LIKE ? OR a.topic LIKE ?",
			"%"+req.Keywords+"%",
			req.Keywords,
			req.Keywords,
			req.Keywords,
			req.Keywords+"%",
			"%"+req.Keywords+"%",
			"%"+req.Keywords+"%",
		)
	}

	querySQL, args := query.Order("c.create_at").Desc().Limit(req.OffSet(), uint(req.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query config error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := config.NewConfigSet()
	for rows.Next() {
		ins := config.NewDefaultConfig()
		app := ins.Base
		desc := ins.Describe
		err := rows.Scan(
			&desc.Id, &desc.ApplicationId, &desc.Name, &desc.Host, &desc.Port, &desc.Env, &desc.Type, &desc.Source, &desc.CreateAt, &desc.UpdateAt,
			&app.Name, &app.Repo, &app.Branch, &app.Module, &app.Topic, &app.Job, &app.Description, &app.Status,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query host error, %s", err.Error())
		}
		set.Add(ins)
	}

	// 获取total SELECT COUNT(*) FROMT t Where ....
	countSQL, args := query.BuildCount()
	countStmt, err := s.db.Prepare(countSQL)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	defer countStmt.Close()
	err = countStmt.QueryRow(args...).Scan(&set.Total)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}

	return set, nil
}

func (s *service) UpdateConfig(ctx context.Context, req *config.UpdateConfigRequest) (
	*config.Config, error) {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 检测参数合法性
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update config error, %s", err)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("start tx error, %s", err)
	}

	// 查询出该条实例的数据
	conf, err := s.DescribeConfig(ctx, config.NewDescribeConfigRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case config.UpdateMode_PATCH:
		conf.Patch(req.UpdateConfigData)
	default:
		conf.Put(req.UpdateConfigData)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(updateConfigSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	desc := conf.Describe
	_, err = stmt.Exec(
		desc.ApplicationId, desc.Name, desc.Host, desc.Port, desc.Env, desc.Type, desc.Source, desc.CreateAt, desc.UpdateAt,
		desc.Id,
	)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return conf, nil
}

func (s *service) DescribeConfig(ctx context.Context, req *config.DescribeConfigRequest) (
	*config.Config, error) {
	query := sqlbuilder.NewQuery(queryConfigSQL)
	querySQL, args := query.Where("c.id = ?", req.Id).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query config error, %s", err.Error())
	}
	defer queryStmt.Close()

	conf := config.NewDefaultConfig()
	app := conf.Base
	desc := conf.Describe
	err = queryStmt.QueryRow(args...).Scan(
		&desc.Id, &desc.ApplicationId, &desc.Name, &desc.Host, &desc.Port, &desc.Env, &desc.Type, &desc.Source, &desc.CreateAt, &desc.UpdateAt,
		&app.Name, &app.Repo, &app.Branch, &app.Module, &app.Topic, &app.Job, &app.Description, &app.Status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%#v not found", req)
		}
		return nil, exception.NewInternalServerError("appribe config error, %s", err.Error())
	}

	return conf, nil
}

func (s *service) DeleteConfig(ctx context.Context, req *config.DeleteConfigRequest) (
	*config.Config, error) {
	conf, err := s.DescribeConfig(ctx, config.NewDescribeConfigRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.delete(ctx, req); err != nil {
		return nil, err
	}

	return conf, nil
}
