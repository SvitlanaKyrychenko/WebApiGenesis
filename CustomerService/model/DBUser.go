package model

import (
	"github.com/segmentio/ksuid"
)

type DBUser struct {
	Guid     ksuid.KSUID `json:"guid"`
	Email    string      `json:"gmail"`
	Password string      `json:"password"`
}

func (user DBUser) Name() string {
	return "DBUser"
}
func (user DBUser) GetGuid() ksuid.KSUID {
	return user.Guid
}
