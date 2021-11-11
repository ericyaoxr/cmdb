package application

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

const (
	AppName = "Application"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewQueryApplicationRequestFromHTTP(r *http.Request) *QueryApplicationRequest {
	qs := r.URL.Query()

	ps := qs.Get("page_size")
	pn := qs.Get("page_number")
	kw := qs.Get("keywords")

	psUint64, _ := strconv.ParseUint(ps, 10, 64)
	pnUint64, _ := strconv.ParseUint(pn, 10, 64)

	if psUint64 == 0 {
		psUint64 = 20
	}
	if pnUint64 == 0 {
		pnUint64 = 1
	}
	return &QueryApplicationRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

func (req *QueryApplicationRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}

func NewDescribeApplicationRequestWithID(id string) *DescribeApplicationRequest {
	return &DescribeApplicationRequest{
		Id: id,
	}
}

func NewDeleteApplicationRequestWithID(id string) *DeleteApplicationRequest {
	return &DeleteApplicationRequest{Id: id}
}

func NewUpdateApplicationRequest(id string) *UpdateApplicationRequest {
	return &UpdateApplicationRequest{
		Id:                    id,
		UpdateMode:            UpdateMode_PUT,
		UpdateApplicationData: &UpdateApplicationData{},
	}
}

func (req *UpdateApplicationRequest) Validate() error {
	return validate.Struct(req)
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewApplicationSet(),
	}
}

type PagerResult struct {
	Data    *ApplicationSet
	Err     error
	HasNext bool
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}
