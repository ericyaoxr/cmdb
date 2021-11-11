package config

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

const (
	AppName = "Config"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewQueryConfigRequestFromHTTP(r *http.Request) *QueryConfigRequest {
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
	return &QueryConfigRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

func (req *QueryConfigRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}

func NewDescribeConfigRequestWithID(id string) *DescribeConfigRequest {
	return &DescribeConfigRequest{
		Id: id,
	}
}

func NewDeleteConfigRequestWithID(id string) *DeleteConfigRequest {
	return &DeleteConfigRequest{Id: id}
}

func NewUpdateConfigRequest(id string) *UpdateConfigRequest {
	return &UpdateConfigRequest{
		Id:             id,
		UpdateMode:     UpdateMode_PUT,
		UpdateConfigData: &UpdateConfigData{},
	}
}

func (req *UpdateConfigRequest) Validate() error {
	return validate.Struct(req)
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewConfigSet(),
	}
}

type PagerResult struct {
	Data    *ConfigSet
	Err     error
	HasNext bool
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}
