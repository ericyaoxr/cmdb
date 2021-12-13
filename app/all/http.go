package all

import (
	_ "github.com/ericyaoxr/cmdb/app/application/http"
	_ "github.com/ericyaoxr/cmdb/app/bill/http"
	_ "github.com/ericyaoxr/cmdb/app/config/http"
	_ "github.com/ericyaoxr/cmdb/app/host/http"
	_ "github.com/ericyaoxr/cmdb/app/rds/http"
	_ "github.com/ericyaoxr/cmdb/app/resource/http"
	_ "github.com/ericyaoxr/cmdb/app/secret/http"
	_ "github.com/ericyaoxr/cmdb/app/task/http"
)
