package pkg

import (
	"github.com/ericyaoxr/cmdb/pkg/host"
	"github.com/ericyaoxr/cmdb/pkg/resource"
	"github.com/ericyaoxr/cmdb/pkg/secret"
	"github.com/ericyaoxr/cmdb/pkg/task"
)

var (
	Host     host.Service
	Secret   secret.Service
	Resource resource.Service
	Task     task.Service
)
