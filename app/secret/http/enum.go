package http

import (
	"net/http"

	"github.com/ericyaoxr/cmdb/app/secret"
	"github.com/ericyaoxr/cmdb/utils"
	"github.com/ericyaoxr/mcube/http/response"
	"github.com/julienschmidt/httprouter"
)

func (h *handler) ListCrendentialType(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	resp := []utils.EnumDescribe{
		{Value: secret.CrendentialAPIKey.String(), Describe: "API凭证"},
		{Value: secret.CrendentialPassword.String(), Describe: "用户名密码"},
	}
	response.Success(w, resp)
}
