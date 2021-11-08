package http

import (
	"github.com/ericyaoxr/mcube/http/router"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	"github.com/ericyaoxr/cmdb/app"
	"github.com/ericyaoxr/cmdb/app/host"
)

var (
	h = &handler{}
)

type handler struct {
	service host.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(host.AppName)
	h.service = app.GetGrpcApp(host.AppName).(host.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return host.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	r.Handle("GET", "/hosts", h.QueryHost)
	r.Handle("POST", "/hosts", h.CreateHost)
	r.Handle("GET", "/hosts/:id", h.DescribeHost)
	r.Handle("DELETE", "/hosts/:id", h.DeleteHost)
	r.Handle("PUT", "/hosts/:id", h.PutHost)
	r.Handle("PATCH", "/hosts/:id", h.PatchHost)
}

func init() {
	app.RegistryHttpApp(h)
}
