package model

import "github.com/segmentio/ksuid"

type DBUser struct{
	Guid ksuid.KSUID `json:"guid"`
	Password string `json:"password"`
	Email string `json:"gmail"`
}

func (User DBUser) Name() string {
	return "DBUser"
}
func (user DBUser) GetGuid() ksuid.KSUID {
	return user.Guid
}