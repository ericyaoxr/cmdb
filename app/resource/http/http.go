package http

import (
	"github.com/ericyaoxr/mcube/http/router"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	"github.com/ericyaoxr/cmdb/app"
	"github.com/ericyaoxr/cmdb/app/resource"
)

var (
	h = &handler{}
)

type handler struct {
	service resource.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(resource.AppName)
	h.service = app.GetGrpcApp(resource.AppName).(resource.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return resource.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	r.Handle("GET", "/search", h.SearchResource)
	r.Handle("GET", "/vendors", h.ListVendor)
	r.Handle("GET", "/regions", h.ListVendorRegion)
	r.Handle("GET", "/resource_types", h.ListResourceType)
}

func init() {
	app.RegistryHttpApp(h)
}
