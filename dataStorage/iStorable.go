package dataStorage

import "github.com/segmentio/ksuid"

type IsStorable interface {
	GetGuid() ksuid.KSUID
	Name() string
}

type ClassStorable struct{
	Guid ksuid.KSUID
	NameClass string
}

func (class ClassStorable) Name() string {
	return class.NameClass
}
func (class ClassStorable) GetGuid() ksuid.KSUID {
	return class.Guid
}

