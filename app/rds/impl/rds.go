package impl

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/ericyaoxr/mcube/exception"
	"github.com/ericyaoxr/mcube/sqlbuilder"
	"github.com/ericyaoxr/mcube/types/ftime"
	"github.com/rs/xid"

	"github.com/ericyaoxr/cmdb/app/rds"
)

const (
	insertResourceSQL = `INSERT INTO resource (
		id,vendor,region,zone,create_at,expire_at,category,type,instance_id,
		name,description,status,update_at,sync_at,sync_accout,public_ip,
		private_ip,pay_type,describe_hash,resource_hash
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);`
	insertRdsSQL = `INSERT INTO rds (
		resource_id,cpu,memory,volume,engine_version,instance_nodes,instance_type,
		vport,master,slave,
		deploy_mode,qps,wan_port
	) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?);`
	updateResourceSQL = `UPDATE resource SET 
		expire_at=?,category=?,type=?,name=?,description=?,
		status=?,update_at=?,sync_at=?,sync_accout=?,
		public_ip=?,private_ip=?,pay_type=?,describe_hash=?,resource_hash=?
	WHERE id = ?`
	updateRdsSQL = `UPDATE rds SET 
		cpu=?,memory=?,volume=?,engine_version=?,instance_nodes=?,instance_type=?,
		vport=?,master=?,
		slave=?,deploy_mode=?,qps=?,wan_port=?
	WHERE resource_id = ?`

	queryRdsSQL       = `SELECT * FROM resource as r LEFT JOIN rds h ON r.id=h.resource_id`
	deleteRdsSQL      = `DELETE FROM rds WHERE resource_id = ?;`
	deleteResourceSQL = `DELETE FROM resource WHERE id = ?;`
)

func (s *service) SaveRds(ctx context.Context, h *rds.Rds) (
	*rds.Rds, error) {
	h.Base.Id = xid.New().String()
	h.Describe.ResourceId = h.Base.Id
	h.Base.SyncAt = ftime.Now().Timestamp()

	if err := s.save(ctx, h); err != nil {
		return nil, err
	}
	return h, nil
}

func (s *service) QueryRds(ctx context.Context, req *rds.QueryRdsRequest) (
	*rds.RdsSet, error) {
	query := sqlbuilder.NewQuery(queryRdsSQL)

	if req.Keywords != "" {
		query.Where("r.name LIKE ? OR r.id = ? OR r.instance_id = ? OR r.private_ip LIKE ? OR r.public_ip LIKE ?",
			"%"+req.Keywords+"%",
			req.Keywords,
			req.Keywords,
			req.Keywords+"%",
			req.Keywords+"%",
		)
	}

	querySQL, args := query.Order("sync_at").Desc().Limit(req.OffSet(), uint(req.PageSize)).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query rds error, %s", err.Error())
	}
	defer queryStmt.Close()

	rows, err := queryStmt.Query(args...)
	if err != nil {
		return nil, exception.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	set := rds.NewRdsSet()
	var (
		publicIPList, privateIPList string
	)
	for rows.Next() {
		ins := rds.NewDefaultRds()
		base := ins.Base
		info := ins.Information
		desc := ins.Describe
		err := rows.Scan(
			&base.Id, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
			&info.Category, &info.Type, &base.InstanceId, &info.Name, &info.Description,
			&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
			&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash, &desc.ResourceId,
			&desc.Cpu, &desc.Memory, &desc.Volume, &desc.EngineVersion, &desc.InstanceNodes, &desc.InstanceType,
			&desc.Port, &desc.Master, &desc.DeployMode, &desc.Qps,
			&desc.WanPort,
		)
		if err != nil {
			return nil, exception.NewInternalServerError("query rds error, %s", err.Error())
		}
		info.LoadPrivateIPString(privateIPList)
		info.LoadPublicIPString(publicIPList)
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

func (s *service) UpdateRds(ctx context.Context, req *rds.UpdateRdsRequest) (
	*rds.Rds, error) {
	var (
		stmt *sql.Stmt
		err  error
	)

	// 检测参数合法性
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate update rds error, %s", err)
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("start tx error, %s", err)
	}

	// 查询出该条实例的数据
	ins, err := s.DescribeRds(ctx, rds.NewDescribeRdsRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	oldRH, oldDH := ins.Base.ResourceHash, ins.Base.DescribeHash

	switch req.UpdateMode {
	case rds.UpdateMode_PATCH:
		ins.Patch(req.UpdateRdsData)
	default:
		ins.Put(req.UpdateRdsData)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
	}()

	if oldRH != ins.Base.ResourceHash {
		// 避免SQL注入, 请使用Prepare
		stmt, err = tx.Prepare(updateResourceSQL)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		base := ins.Base
		info := ins.Information
		_, err = stmt.Exec(
			info.ExpireAt, info.Category, info.Type, info.Name, info.Description,
			info.Status, info.UpdateAt, base.SyncAt, info.SyncAccount,
			info.PublicIPToString(), info.PrivateIPToString(), info.PayType, base.DescribeHash, base.ResourceHash,
			ins.Describe.ResourceId,
		)
		if err != nil {
			return nil, err
		}
	} else {
		s.log.Debug("resource data hash not changed, needn't update")
	}

	if oldDH != ins.Base.DescribeHash {
		// 避免SQL注入, 请使用Prepare
		stmt, err = tx.Prepare(updateRdsSQL)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		// base := ins.Base
		desc := ins.Describe
		_, err = stmt.Exec(
			desc.Cpu, desc.Memory, desc.Volume, desc.EngineVersion, desc.InstanceNodes, desc.InstanceType,
			desc.Port, desc.Master,
			desc.Slave, desc.DeployMode, desc.Qps,
			desc.WanPort,
		)
		if err != nil {
			return nil, err
		}
	} else {
		s.log.Debug("describe data hash not changed, needn't update")
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return ins, nil
}

func (s *service) DescribeRds(ctx context.Context, req *rds.DescribeRdsRequest) (
	*rds.Rds, error) {
	query := sqlbuilder.NewQuery(queryRdsSQL)
	cond, val := req.Where()
	querySQL, args := query.Where(cond, val).BuildQuery()
	s.log.Debugf("sql: %s", querySQL)

	queryStmt, err := s.db.Prepare(querySQL)
	if err != nil {
		return nil, exception.NewInternalServerError("prepare query rds error, %s", err.Error())
	}
	defer queryStmt.Close()

	ins := rds.NewDefaultRds()
	var (
		publicIPList, privateIPList string
	)
	base := ins.Base
	info := ins.Information
	desc := ins.Describe
	err = queryStmt.QueryRow(args...).Scan(
		&base.Id, &base.Vendor, &base.Region, &base.Zone, &base.CreateAt, &info.ExpireAt,
		&info.Category, &info.Type, &base.InstanceId, &info.Name, &info.Description,
		&info.Status, &info.UpdateAt, &base.SyncAt, &info.SyncAccount,
		&publicIPList, &privateIPList, &info.PayType, &base.DescribeHash, &base.ResourceHash, &desc.ResourceId,
		&desc.Cpu, &desc.Memory, &desc.Volume, &desc.EngineVersion, &desc.InstanceNodes, &desc.InstanceType,
		&desc.Port, &desc.Master, &desc.Slave, &desc.DeployMode,
		&desc.Qps, &desc.WanPort,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.NewNotFound("%#v not found", req)
		}
		return nil, exception.NewInternalServerError("describe rds error, %s", err.Error())
	}

	info.LoadPrivateIPString(privateIPList)
	info.LoadPublicIPString(publicIPList)

	return ins, nil
}

func (s *service) DeleteRds(ctx context.Context, req *rds.DeleteRdsRequest) (
	*rds.Rds, error) {
	ins, err := s.DescribeRds(ctx, rds.NewDescribeRdsRequestWithID(req.Id))
	if err != nil {
		return nil, err
	}

	if err := s.delete(ctx, req); err != nil {
		return nil, err
	}

	return ins, nil
}
