package http

import (
	"net/http"

	"github.com/ericyaoxr/mcube/http/context"
	"github.com/ericyaoxr/mcube/http/request"
	"github.com/ericyaoxr/mcube/http/response"

	"github.com/ericyaoxr/cmdb/app/config"
)

func (h *handler) QueryConfig(w http.ResponseWriter, r *http.Request) {
	query := config.NewQueryConfigRequestFromHTTP(r)
	set, err := h.service.QueryConfig(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateConfig(w http.ResponseWriter, r *http.Request) {
	ins := config.NewDefaultConfig()
	if err := request.GetDataFromRequest(r, ins); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.SaveConfig(r.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeConfig(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := config.NewDescribeConfigRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DescribeConfig(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteConfig(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := config.NewDeleteConfigRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DeleteConfig(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PutConfig(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := config.NewUpdateConfigRequest(ctx.PS.ByName("id"))

	if err := request.GetDataFromRequest(r, req.UpdateConfigData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateConfig(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) PatchConfig(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := config.NewUpdateConfigRequest(ctx.PS.ByName("id"))
	req.UpdateMode = config.UpdateMode_PATCH

	if err := request.GetDataFromRequest(r, req.UpdateConfigData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateConfig(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
