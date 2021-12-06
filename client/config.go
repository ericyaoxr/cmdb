package client

// NewDefaultConfig todo
func NewConfig(address string) *Config {
	return &Config{
		address:        address,
	}
}

// Config 客户端配置
type Config struct {
	address string
}

// SetAddress todo
func (c *Config) SetAddress(addr string) {
	c.address = addr
}

// Address 地址
func (c *Config) Address() string {
	return c.address
}
