package http

import (
	"github.com/ericyaoxr/mcube/http/router"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	"github.com/ericyaoxr/cmdb/app"
	"github.com/ericyaoxr/cmdb/app/config"
)

var (
	h = &handler{}
)

type handler struct {
	service config.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(config.AppName)
	h.service = app.GetGrpcApp(config.AppName).(config.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return config.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	r.Handle("GET", "/configs", h.QueryConfig)
	r.Handle("POST", "/configs", h.CreateConfig)
	r.Handle("GET", "/configs/:id", h.DescribeConfig)
	r.Handle("DELETE", "/configs/:id", h.DeleteConfig)
	r.Handle("PUT", "/configs/:id", h.PutConfig)
	r.Handle("PATCH", "/configs/:id", h.PatchConfig)
}

func init() {
	app.RegistryHttpApp(h)
}
