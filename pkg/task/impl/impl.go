package impl

import (
	"database/sql"

	"github.com/ericyaoxr/cmdb/conf"
	"github.com/ericyaoxr/cmdb/pkg"
	"github.com/ericyaoxr/cmdb/pkg/host"
	"github.com/ericyaoxr/cmdb/pkg/secret"
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
