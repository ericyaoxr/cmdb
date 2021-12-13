package http

import (
	"github.com/ericyaoxr/mcube/http/router"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	"github.com/ericyaoxr/cmdb/app"
	"github.com/ericyaoxr/cmdb/app/rds"
)

var (
	h = &handler{}
)

type handler struct {
	service rds.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(rds.AppName)
	h.service = app.GetGrpcApp(rds.AppName).(rds.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return rds.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	r.Handle("GET", "/rdss", h.QueryRds)
	r.Handle("POST", "/rdss", h.CreateRds)
	r.Handle("GET", "/rdss/:id", h.DescribeRds)
	r.Handle("DELETE", "/rdss/:id", h.DeleteRds)
	r.Handle("PUT", "/rdss/:id", h.PutRds)
	r.Handle("PATCH", "/rdss/:id", h.PatchRds)
}

func init() {
	app.RegistryHttpApp(h)
}
