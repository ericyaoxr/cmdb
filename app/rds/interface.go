package rds

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
)

// use a single instance of Validate, it caches struct info
var (
	validate = validator.New()
)

func NewQueryRdsRequestFromHTTP(r *http.Request) *QueryRdsRequest {
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
	return &QueryRdsRequest{
		PageSize:   psUint64,
		PageNumber: pnUint64,
		Keywords:   kw,
	}
}

func NewQueryRdsRequest() *QueryRdsRequest {
	return &QueryRdsRequest{
		PageSize:   20,
		PageNumber: 1,
	}
}

func (req *QueryRdsRequest) OffSet() int64 {
	return int64(req.PageSize) * int64(req.PageNumber-1)
}

func NewDescribeRdsRequestWithID(id string) *DescribeRdsRequest {
	return &DescribeRdsRequest{
		DescribeBy: DescribeBy_RDS_ID,
		Value:      id,
	}
}

func NewDeleteRdsRequestWithID(id string) *DeleteRdsRequest {
	return &DeleteRdsRequest{Id: id}
}

func NewUpdateRdsRequest(id string) *UpdateRdsRequest {
	return &UpdateRdsRequest{
		Id:            id,
		UpdateMode:    UpdateMode_PUT,
		UpdateRdsData: &UpdateRdsData{},
	}
}

func (req *UpdateRdsRequest) Validate() error {
	return validate.Struct(req)
}

func NewPagerResult() *PagerResult {
	return &PagerResult{
		Data: NewRdsSet(),
	}
}

type PagerResult struct {
	Data    *RdsSet
	Err     error
	HasNext bool
}

// 分页迭代器
type Pager interface {
	Next() *PagerResult
}
