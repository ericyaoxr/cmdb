package http

import (
	"net/http"

	"github.com/ericyaoxr/mcube/http/context"
	"github.com/ericyaoxr/mcube/http/request"
	"github.com/ericyaoxr/mcube/http/response"

	"github.com/ericyaoxr/cmdb/app/rds"
)

func (h *handler) QueryRds(w http.ResponseWriter, r *http.Request) {
	query := rds.NewQueryRdsRequestFromHTTP(r)
	set, err := h.service.QueryRds(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateRds(w http.ResponseWriter, r *http.Request) {
	ins := rds.NewDefaultRds()
	if err := request.GetDataFromRequest(r, ins); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.SaveRds(r.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeRds(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := rds.NewDescribeRdsRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DescribeRds(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteRds(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := rds.NewDeleteRdsRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DeleteRds(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PutRds(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := rds.NewUpdateRdsRequest(ctx.PS.ByName("id"))

	if err := request.GetDataFromRequest(r, req.UpdateRdsData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateRds(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) PatchRds(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := rds.NewUpdateRdsRequest(ctx.PS.ByName("id"))
	req.UpdateMode = rds.UpdateMode_PATCH

	if err := request.GetDataFromRequest(r, req.UpdateRdsData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateRds(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
