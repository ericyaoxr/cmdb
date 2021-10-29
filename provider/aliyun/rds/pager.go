package rds

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/rds"
	"github.com/ericyaoxr/mcube/logger"
	"github.com/ericyaoxr/mcube/logger/zap"

	cmdbRds "github.com/ericyaoxr/cmdb/pkg/rds"
)

func newPager(pageSize int, operater *RdsOperater) *pager {
	req := rds.CreateDescribeDBInstancesRequest()
	req.PageSize = requests.NewInteger(pageSize)

	return &pager{
		size:     pageSize,
		number:   1,
		operater: operater,
		req:      req,
		log:      zap.L().Named("Pagger"),
	}
}

type pager struct {
	size     int
	number   int
	total    int64
	operater *RdsOperater
	req      *rds.DescribeDBInstancesRequest
	log      logger.Logger
}

func (p *pager) Next() *cmdbRds.PagerResult {
	result := cmdbRds.NewPagerResult()

	resp, err := p.operater.Query(p.nextReq())
	if err != nil {
		result.Err = err
		return result
	}

	p.total = int64(resp.Total)

	result.Data = resp
	result.HasNext = p.hasNext()

	p.number++
	return result
}

func (p *pager) WithLogger(log logger.Logger) {
	p.log = log
}

func (p *pager) nextReq() *rds.DescribeDBInstancesRequest {
	p.log.Debug("请求第%d页数据", p.number)
	p.req.PageNumber = requests.NewInteger(p.number)
	return p.req
}

func (p *pager) hasNext() bool {
	return int64(p.number*p.size) < p.total
}