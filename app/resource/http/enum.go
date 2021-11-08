package http

import (
	"net/http"

	"github.com/ericyaoxr/cmdb/app/resource"
	"github.com/ericyaoxr/cmdb/utils"
	"github.com/ericyaoxr/mcube/http/response"
	"github.com/julienschmidt/httprouter"

	tx_region "github.com/ericyaoxr/cmdb/provider/txyun/region"
)

func (h *handler) ListVendor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := []utils.EnumDescribe{
		// {Value: resource.VendorAliYun.String(), Describe: "阿里云"},
		{Value: resource.Vendor_TENCENT.String(), Describe: "腾讯云"},
	}
	response.Success(w, resp)
}

func (h *handler) ListResourceType(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := map[string][]utils.EnumDescribe{
		// resource.VendorAliYun.String(): {
		// 	{Name: "主机", Value: resource.HostResource.String(), Describe: "阿里云ECS"},
		// 	{Name: "关系型数据库", Value: resource.RdsResource.String(), Describe: "阿里云RDS"},
		// },
		resource.Vendor_TENCENT.String(): {
			{Name: "主机", Value: resource.Type_HOST.String(), Describe: "腾讯云CVM"},
			{Name: "关系型数据库", Value: resource.Type_RDS.String(), Describe: "腾讯云CDB"},
		},
	}
	response.Success(w, resp)
}

func (h *handler) ListVendorRegion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := map[string][]utils.EnumDescribe{
		// resource.VendorAliYun.String():  ali_region.Regions,
		resource.Vendor_TENCENT.String(): tx_region.Regions,
	}
	response.Success(w, resp)
}
