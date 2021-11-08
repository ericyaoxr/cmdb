package impl

import (
	"database/sql"

	pkg "github.com/ericyaoxr/cmdb/app"
	"github.com/ericyaoxr/cmdb/app/host"
	"github.com/ericyaoxr/cmdb/app/secret"
	"github.com/ericyaoxr/cmdb/conf"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"
)

var (
	// Service 服务实例
	Service = &service{}
)

type service struct {
	db     *sql.DB
	log    logger.Logger
	host   host.Service
	secret secret.Service
}

func (s *service) Config() error {
	db, err := conf.C().MySQL.GetDB()
	if err != nil {
		return err
	}

	s.log = zap.L().Named("Task")
	s.db = db
	s.host = pkg.Host
	s.secret = pkg.Secret
	return nil
}
