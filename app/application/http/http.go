package http

import (
	"github.com/ericyaoxr/mcube/http/router"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	"github.com/ericyaoxr/cmdb/app"
	"github.com/ericyaoxr/cmdb/app/application"
)

var (
	h = &handler{}
)

type handler struct {
	service application.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(application.AppName)
	h.service = app.GetGrpcApp(application.AppName).(application.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return application.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	r.Handle("GET", "/applications", h.QueryApplication)
	r.Handle("POST", "/applications", h.CreateApplication)
	r.Handle("GET", "/applications/:id", h.DescribeApplication)
	r.Handle("DELETE", "/applications/:id", h.DeleteApplication)
	r.Handle("PUT", "/applications/:id", h.PutApplication)
	r.Handle("PATCH", "/applications/:id", h.PatchApplication)
}

func init() {
	app.RegistryHttpApp(h)
}
