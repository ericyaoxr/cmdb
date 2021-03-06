package config

import (
	"crypto/sha1"
	"encoding/json"

	"github.com/ericyaoxr/cmdb/app/application"
	"github.com/ericyaoxr/mcube/types/ftime"
)

func NewDefaultConfig() *Config {
	return &Config{
		Base:     &application.Base{},
		Describe: &Describe{},
	}
}

func (h *Config) Put(req *UpdateConfigData) {
	h.Describe = req.Describe
	h.Describe.UpdateAt = ftime.Now().Timestamp()
}

func (h *Config) Patch(req *UpdateConfigData) error {
	err := ObjectPatch(h.Describe, req.Describe)
	if err != nil {
		return err
	}
	h.Describe.UpdateAt = ftime.Now().Timestamp()
	return nil
}

func ObjectPatch(old, new interface{}) error {
	newByte, err := json.Marshal(new)
	if err != nil {
		return err
	}
	return json.Unmarshal(newByte, old)
}

func (h *Config) GenHash() error {
	hash := sha1.New()

	b, err := json.Marshal(h)
	if err != nil {
		return err
	}
	hash.Write(b)

	b, err = json.Marshal(h)
	if err != nil {
		return err
	}
	hash.Reset()
	hash.Write(b)
	return nil
}

func NewConfigSet() *ConfigSet {
	return &ConfigSet{
		Items: []*Config{},
	}
}

func (s *ConfigSet) Add(item *Config) {
	s.Items = append(s.Items, item)
}

func (s *ConfigSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
