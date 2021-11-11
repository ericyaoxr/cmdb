package http

import (
	"net/http"

	"github.com/ericyaoxr/mcube/http/context"
	"github.com/ericyaoxr/mcube/http/request"
	"github.com/ericyaoxr/mcube/http/response"

	"github.com/ericyaoxr/cmdb/app/application"
)

func (h *handler) QueryApplication(w http.ResponseWriter, r *http.Request) {
	query := application.NewQueryApplicationRequestFromHTTP(r)
	set, err := h.service.QueryApplication(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) CreateApplication(w http.ResponseWriter, r *http.Request) {
	ins := application.NewDefaultApplication()
	if err := request.GetDataFromRequest(r, ins); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.SaveApplication(r.Context(), ins)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) DescribeApplication(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := application.NewDescribeApplicationRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DescribeApplication(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) DeleteApplication(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := application.NewDeleteApplicationRequestWithID(ctx.PS.ByName("id"))
	set, err := h.service.DeleteApplication(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}

func (h *handler) PutApplication(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := application.NewUpdateApplicationRequest(ctx.PS.ByName("id"))

	if err := request.GetDataFromRequest(r, req.UpdateApplicationData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateApplication(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}

func (h *handler) PatchApplication(w http.ResponseWriter, r *http.Request) {
	ctx := context.GetContext(r)
	req := application.NewUpdateApplicationRequest(ctx.PS.ByName("id"))
	req.UpdateMode = application.UpdateMode_PATCH

	if err := request.GetDataFromRequest(r, req.UpdateApplicationData); err != nil {
		response.Failed(w, err)
		return
	}

	ins, err := h.service.UpdateApplication(r.Context(), req)
	if err != nil {
		response.Failed(w, err)
		return
	}

	response.Success(w, ins)
}
