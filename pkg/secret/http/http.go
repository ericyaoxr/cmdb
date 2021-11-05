package http

import (
	"fmt"

	"github.com/julienschmidt/httprouter"

	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	"github.com/ericyaoxr/cmdb/pkg"
	"github.com/ericyaoxr/cmdb/pkg/secret"
)

var (
	api = &handler{}
)

type handler struct {
	service secret.Service
	log     logger.Logger
}

func (h *handler) Config() error {
	h.log = zap.L().Named("Syncer")
	if pkg.Secret == nil {
		return fmt.Errorf("dependence service secret not ready")
	}
	h.service = pkg.Secret
	return nil
}

func RegistAPI(r *httprouter.Router) {
	api.Config()
	r.POST("/secrets", api.CreateSecret)
	r.GET("/secrets", api.QuerySecret)
	r.GET("/secrets/:id", api.DescribeSecret)
	r.DELETE("/secrets/:id", api.DeleteSecret)
	r.GET("/crendential_types", api.ListCrendentialType)
}
