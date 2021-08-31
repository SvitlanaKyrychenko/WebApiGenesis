package util

import (
	"WebApiGenesis/services"
	"WebApiGenesis/storage"
)

func PrepareRegService() services.Registrar {
	var storage storage.Storage = PrepareMockStorage()
	return services.Registration{Storage: storage}
}
