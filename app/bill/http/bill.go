package http

import (
	"net/http"

	"github.com/ericyaoxr/cmdb/app/bill"
	"github.com/ericyaoxr/mcube/http/response"
)

func (h *handler) QueryBill(w http.ResponseWriter, r *http.Request) {
	query := bill.NewQueryBillRequestFromHTTP(r)
	set, err := h.service.QueryBill(r.Context(), query)
	if err != nil {
		response.Failed(w, err)
		return
	}
	response.Success(w, set)
}
