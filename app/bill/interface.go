package bill

import (
	"net/http"
	"strconv"

	"github.com/ericyaoxr/mcube/pb/page"
)

func NewQueryBillRequest() *QueryBillRequest {
	return &QueryBillRequest{
		Page: &page.PageRequest{
			PageSize:   20,
			PageNumber: 1,
		},
	}
}

func NewQueryBillRequestFromHTTP(r *http.Request) *QueryBillRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}
	return &QueryBillRequest{
		Page: &page.PageRequest{
			PageSize:   psUint64,
			PageNumber: pnUint64,
		},
	}
}
