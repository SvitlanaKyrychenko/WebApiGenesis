package util

import (
	"WebApiGenesis/CustomerService/server"
	"WebApiGenesis/CustomerService/storage"
)

func PrepareRegService() server.Registrar {
	var storage storage.Storage = PrepareMockStorage()
	return server.Registration{Storage: storage}
}
