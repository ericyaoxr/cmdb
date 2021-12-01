package http

import (
	"github.com/ericyaoxr/cmdb/app/bill"
	"github.com/ericyaoxr/mcube/http/router"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	"github.com/ericyaoxr/mcube/app"
)

var (
	h = &handler{}
)

type handler struct {
	service bill.ServiceServer
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named(bill.AppName)
	h.service = app.GetGrpcApp(bill.AppName).(bill.ServiceServer)
	return nil
}

func (h *handler) Name() string {
	return bill.AppName
}

func (h *handler) Registry(r router.SubRouter) {
	hr := r.ResourceRouter("bill")
	hr.Permission(true)
}

func init() {
	app.RegistryHttpApp(h)
}
