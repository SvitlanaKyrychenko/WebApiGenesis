package model

import "github.com/segmentio/ksuid"

type User struct {
	guid  ksuid.KSUID
	gmail string
}

func (user User) Guid() ksuid.KSUID {
	return user.guid
}

func (user User) Gmail() string {
	return user.gmail
}
