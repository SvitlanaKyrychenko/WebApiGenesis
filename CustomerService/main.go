package main

import (
	"Bitcoin/CustomerService/model"
	"Bitcoin/CustomerService/server"
	"Bitcoin/CustomerService/storage"
	grpcModel "Bitcoin/GRPCMessage"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	portAuth = flag.Int("portAuth", 10006, "The server portAuth")
	portReg  = flag.Int("portReg", 10007, "The server portReg")
)

func main() {
	startAuthenticationServer()
	startRegistrationServer()
}

func startAuthenticationServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *portAuth))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcModel.RegisterAuthenticatorServer(grpcServer, newAuthServer())
	grpcServer.Serve(lis)
}

func startRegistrationServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *portReg))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcModel.RegisterRegistrarServer(grpcServer, newRegServer())
	grpcServer.Serve(lis)
}

func newAuthServer() grpcModel.AuthenticatorServer {
	var convertor model.Convertor = model.JSONGConvertor{}
	var storage storage.Storage = storage.FileStorage{Convertor: convertor}
	server := &server.AuthenticationServer{Storage: storage}
	return server
}

func newRegServer() grpcModel.RegistrarServer {
	var convertor model.Convertor = model.JSONGConvertor{}
	var storage storage.Storage = storage.FileStorage{Convertor: convertor}
	server := &server.Registration{Storage: storage}
	return server
}
