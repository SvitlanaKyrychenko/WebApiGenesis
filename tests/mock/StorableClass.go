package mock

import "github.com/segmentio/ksuid"

type StorableClass struct {
	Guid    ksuid.KSUID `json:"guid"`
	Message string      `json:"message"`
}

func (class StorableClass) Name() string {
	return "mockClass"
}
func (class StorableClass) GetGuid() ksuid.KSUID {
	return class.Guid
}
