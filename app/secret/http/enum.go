package http

import (
	"net/http"

	"github.com/ericyaoxr/cmdb/app/secret"
	"github.com/ericyaoxr/cmdb/utils"
	"github.com/ericyaoxr/mcube/http/response"
)

func (h *handler) ListCrendentialType(w http.ResponseWriter, r *http.Request) {
	resp := []utils.EnumDescribe{
		{Value: secret.Type_API_KEY.String(), Describe: "API凭证"},
		{Value: secret.Type_PASSWORD.String(), Describe: "用户名密码"},
	}
	response.Success(w, resp)
}
