package mock

import (
	"WebApiGenesis/model"
	"github.com/segmentio/ksuid"
	"sync"
)

type VirtualStorage struct {
	m       sync.Mutex
	storage map[ksuid.KSUID][]byte
}

type Storage struct {
	virtualStorage VirtualStorage
	Convertor      model.Convertor
}

func (storage *Storage) Init() {
	storage.virtualStorage.storage = make(map[ksuid.KSUID][]byte)
}

func (storage Storage) AddOrUpdateAsync(obj model.FileStorable) error {
	class, err := storage.Convertor.ConvertByte(obj)
	if err != nil {
		return err
	}
	storage.virtualStorage.m.Lock()
	storage.virtualStorage.storage[obj.GetGuid()] = class
	storage.virtualStorage.m.Unlock()
	return nil
}

func (storage Storage) GetAsync(obj model.FileStorable) ([]byte, error) {
	storage.virtualStorage.m.Lock()
	class, contains := storage.virtualStorage.storage[obj.GetGuid()]
	storage.virtualStorage.m.Unlock()
	if contains {
		return class, nil
	}
	return nil, nil
}

func (storage Storage) GetALLAsync(obj model.FileStorable) (map[ksuid.KSUID][]byte, error) {
	result := make(map[ksuid.KSUID][]byte)
	storage.virtualStorage.m.Lock()
	result = storage.virtualStorage.storage
	storage.virtualStorage.m.Unlock()
	return result, nil
}

func (storage Storage) Contains(obj model.FileStorable) bool {
	storage.virtualStorage.m.Lock()
	_, contains := storage.virtualStorage.storage[obj.GetGuid()]
	storage.virtualStorage.m.Unlock()
	return contains
}
