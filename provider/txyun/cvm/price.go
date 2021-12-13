package cvm

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {

	credential := common.NewCredential(
		"SecretId",
		"SecretKey",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	client, _ := cvm.NewClient(credential, "ap-guangzhou", cpf)

	request := cvm.NewInquiryPriceRenewInstancesRequest()

	request.InstanceIds = common.StringPtrs([]string{"ins-ivqfp5jg"})
	request.InstanceChargePrepaid = &cvm.InstanceChargePrepaid{
		Period: common.Int64Ptr(12),
	}

	response, err := client.InquiryPriceRenewInstances(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", response.ToJsonString())
}
