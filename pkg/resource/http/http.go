package http
package http

import (
	"fmt"

	"github.com/julienschmidt/httprouter"

	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	"github.com/ericyaoxr/cmdb/pkg"
	"github.com/ericyaoxr/cmdb/pkg/resource"
)

var (
	api = &handler{}
)

type handler struct {
	service resource.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Resource")
	if pkg.Resource == nil {
		return fmt.Errorf("dependence service resource not ready")
	}
	h.service = pkg.Resource
	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.GET("/search", api.SearchResource)
	r.GET("/vendors", api.ListVendor)
	r.GET("/regions", api.ListVendorRegion)
	r.GET("/resource_types", api.ListResourceType)
}
