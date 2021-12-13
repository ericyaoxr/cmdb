package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ericyaoxr/mcube/exception"
	"github.com/ericyaoxr/mcube/sqlbuilder"
	"github.com/ericyaoxr/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/ericyaoxr/cmdb/app/application"
)

const (
	insertApplicationSQL = `INSERT INTO application (
		id,name,repo,branch,module,topic,job,
		description,create_at,update_at,
		status,project_id
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?);`
	updateApplicationSQL = `UPDATE application SET 
	name=?,repo=?,branch=?,module=?,topic=?,
	job=?,description=?,create_at=?,update_at=?,
	status=?,project_id=?
	WHERE id = ?`

	queryApplicationSQL  = `SELECT * FROM application`
	deleteApplicationSQL = `DELETE FROM application WHERE id = ?;`
)

func (s *service) SaveApplication(ctx context.Context, h *application.Application) (
	*application.Application, error) {
	h.Base.Id = xid.New().String()
	h.Base.CreateAt = ftime.Now().Timestamp()
	h.Base.UpdateAt = ftime.Now().Timestamp()

	if err := s.save(ctx, h); err != nil {
		return nil, err
	}

	return h, nil
}

func (s *service) QueryApplication(ctx context.Context, req *application.QueryApplicationRequest) (
	*application.ApplicationSet, error) {
	query := sqlbuilder.NewQuery(queryApplicationSQL)

	if req.Keywords != "" {
		query.Where("id = ? OR name LIKE ? OR Repo = ? OR Branch = ? OR Module LIKE ? OR Topic LIKE ? OR Job LIKE ?",
			req.Keywords,
			"%"+req.Keywords+"%",
			req.Keywords,
			req.Keywords,
			req.Keywords+"%",
			"%"+req.Keywords+"%",
			req.Keywords,
		)
	}

	querySQL, args := query.Order("project_id").Desc().Limit(req.OffSet(), uint(req.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query application error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := application.NewApplicationSet()
	for rows.Next() {
		ins := application.NewDefaultApplication()
		app := ins.Base
		err := rows.Scan(
			&app.Id, &app.Name, &app.Repo, &app.Branch, &app.Module, &app.Topic, &app.Job,
			&app.Description, &app.CreateAt, &app.UpdateAt, &app.Status, &app.ProjectId,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query application error, %s", err.Error())
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

func (s *service) UpdateApplication(ctx context.Context, req *application.UpdateApplicationRequest) (
	*application.Application, error) {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 检测参数合法性
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update application error, %s", err)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("start tx error, %s", err)
	}

	// 查询出该条应用的数据
	ins, err := s.DescribeApplication(ctx, application.NewDescribeApplicationRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	switch req.UpdateMode {
	case application.UpdateMode_PATCH:
		ins.Patch(req.UpdateApplicationData)
	default:
		ins.Put(req.UpdateApplicationData)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	// 避免SQL注入, 请使用Prepare
	stmt, err = tx.Prepare(updateApplicationSQL)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	app := ins.Base
	_, err = stmt.Exec(
		app.Name, app.Repo, app.Branch, app.Module, app.Topic,
		app.Job, app.Description, app.CreateAt, app.UpdateAt, app.Status, app.ProjectId,
		app.Id,
	)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DescribeApplication(ctx context.Context, req *application.DescribeApplicationRequest) (
	*application.Application, error) {
	query := sqlbuilder.NewQuery(queryApplicationSQL)
	querySQL, args := query.Where("id = ?", req.Id).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query application error, %s", err.Error())
	}
	defer queryStmt.Close()

	ins := application.NewDefaultApplication()
	app := ins.Base
	err = queryStmt.QueryRow(args...).Scan(
		&app.Id, &app.Name, &app.Repo, &app.Branch, &app.Module, &app.Topic,
		&app.Job, &app.Description, &app.CreateAt, &app.UpdateAt, &app.Status, &app.ProjectId,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%#v not found", req)
		}
		return nil, exception.NewInternalServerError("describe application error, %s", err.Error())
	}

	return ins, nil
}

func (s *service) DeleteApplication(ctx context.Context, req *application.DeleteApplicationRequest) (
	*application.Application, error) {
	ins, err := s.DescribeApplication(ctx, application.NewDescribeApplicationRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.delete(ctx, req); err != nil {
		return nil, err
	}

	return ins, nil
}
