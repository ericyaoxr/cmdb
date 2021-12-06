package client_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ericyaoxr/cmdb/app/resource"
	"github.com/ericyaoxr/cmdb/client"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	should := assert.New(t)
	conf := client.NewConfig("localhost:18050")
	c, err := client.NewClient(conf)
	if should.NoError(err) {
		rs, err := c.Resource().Search(context.Background(), resource.NewSearchRequest())
		should.NoError(err)
		fmt.Println(rs)
	}
}
