package application

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"

	"github.com/ericyaoxr/mcube/types/ftime"
)

func NewDefaultApplication() *Application {
	return &Application{
		Status: "已上线",
	}
}

func (h *Application) Put(req *UpdateApplicationData) {
	h.UpdateAt = ftime.Now().Timestamp()
	h.GenHash()
}

func (h *Application) Patch(req *UpdateApplicationData) error {
	err := ObjectPatch(h, req)
	if err != nil {
		return err
	}

	err = ObjectPatch(h, req)
	if err != nil {
		return err
	}

	h.UpdateAt = ftime.Now().Timestamp()
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

func (h *Application) GenHash() error {
	hash := sha1.New()

	b, err := json.Marshal(h)
	if err != nil {
		return err
	}
	hash.Write(b)
	h.Id = fmt.Sprintf("%x", hash.Sum(nil))

	b, err = json.Marshal(h)
	if err != nil {
		return err
	}
	hash.Reset()
	hash.Write(b)
	h.Id = fmt.Sprintf("%x", hash.Sum(nil))
	return nil
}

func NewApplicationSet() *ApplicationSet {
	return &ApplicationSet{
		Items: []*Application{},
	}
}

func (s *ApplicationSet) Add(item *Application) {
	s.Items = append(s.Items, item)
}

func (s *ApplicationSet) ToJsonString() string {
	b, _ := json.Marshal(s)
	return string(b)
}
