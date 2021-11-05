package http

import (
	"net/http"

	"github.com/ericyaoxr/cmdb/pkg/resource"
	"github.com/ericyaoxr/cmdb/utils"
	"github.com/ericyaoxr/mcube/http/response"
	"github.com/julienschmidt/httprouter"

	ali_region "github.com/ericyaoxr/cmdb/provider/aliyun/region"
	hw_region "github.com/ericyaoxr/cmdb/provider/huawei/region"
	tx_region "github.com/ericyaoxr/cmdb/provider/txyun/region"
)

func (h *handler) ListVendor(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := []utils.EnumDescribe{
		{Value: resource.VendorAliYun.String(), Describe: "阿里云"},
		{Value: resource.VendorTencent.String(), Describe: "腾讯云"},
		{Value: resource.VendorHuaWei.String(), Describe: "华为云"},
		{Value: resource.VendorVsphere.String(), Describe: "Vsphere"},
	}
	response.Success(w, resp)
}

func (h *handler) ListResourceType(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := map[string][]utils.EnumDescribe{
		resource.VendorAliYun.String(): {
			{Name: "主机", Value: resource.HostResource.String(), Describe: "阿里云ECS"},
			{Name: "关系型数据库", Value: resource.RdsResource.String(), Describe: "阿里云RDS"},
		},
		resource.VendorTencent.String(): {
			{Name: "主机", Value: resource.HostResource.String(), Describe: "腾讯云CVM"},
			{Name: "关系型数据库", Value: resource.RdsResource.String(), Describe: "腾讯云CDB"},
		},
		resource.VendorHuaWei.String(): {
			{Name: "主机", Value: resource.HostResource.String(), Describe: "华为云ECS"},
			{Name: "关系型数据库", Value: resource.RdsResource.String(), Describe: "华为云RDS"},
		},
		resource.VendorVsphere.String(): {
			{Name: "主机", Value: resource.HostResource.String(), Describe: "VMware vm"},
		},
	}
	response.Success(w, resp)
}

func (h *handler) ListVendorRegion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := map[string][]utils.EnumDescribe{
		resource.VendorAliYun.String():  ali_region.Regions,
		resource.VendorTencent.String(): tx_region.Regions,
		resource.VendorHuaWei.String():  hw_region.Regions,
	}
	response.Success(w, resp)
}
