package rds

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/ericyaoxr/cmdb/app/resource"
	"github.com/ericyaoxr/mcube/types/ftime"
)

func NewDefaultRds() *Rds {
	return &Rds{
		Base: &resource.Base{
			ResourceType: resource.Type_RDS,
		},
		Information: &resource.Information{},
		Describe:    &Describe{},
	}
}

func (h *Rds) Put(req *UpdateRdsData) {
	h.Information = req.Information
	h.Describe = req.Describe
	h.Information.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()
}

func (h *Rds) ShortDesc() string {
	return fmt.Sprintf("%s %s", h.Information.Name, h.Information.PrivateIp)
}

func (h *Rds) Patch(req *UpdateRdsData) error {
	err := ObjectPatch(h.Information, req.Information)
	if err != nil {
		return err
	}

	err = ObjectPatch(h.Describe, req.Describe)
	if err != nil {
		return err
	}

	h.Information.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()
	return nil
}

func ObjectPatch(old, new interface{}) error {
	newByte, err := json.Marshal(new)
	if err != nil {
		return err
	}
	return json.Unmarshal(newByte, old)
}

func (h *Rds) GenHash() error {
	hash := sha1.New()

	b, err := json.Marshal(h.Information)
	if err != nil {
		return err
	}
	hash.Write(b)
	h.Base.ResourceHash = fmt.Sprintf("%x", hash.Sum(nil))

	b, err = json.Marshal(h.Describe)
	if err != nil {
		return err
	}
	hash.Reset()
	hash.Write(b)
	h.Base.DescribeHash = fmt.Sprintf("%x", hash.Sum(nil))
	return nil
}

func NewRdsSet() *RdsSet {
	return &RdsSet{
		Items: []*Rds{},
	}
}

func (s *RdsSet) Add(item *Rds) {
	s.Items = append(s.Items, item)
}

func (s *RdsSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}

func (req *DescribeRdsRequest) Where() (string, interface{}) {
	switch req.DescribeBy {
	case DescribeBy_RDS_ID:
		return "id = ?", req.Value
	case DescribeBy_INSTANCE_ID:
		return "instance_id = ?", req.Value
	default:
		return "", nil
	}
}
