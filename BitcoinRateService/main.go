package main

import (
	"WebApiGenesis/BitcoinRateService/bitcoin"
	grpcModel "WebApiGenesis/GRPCMessage/model"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	portBtsRate = flag.Int("portBtsRate", 10005, "The server portAuth")
)

func main() {
	startBtsRateServer()
}

func startBtsRateServer() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *portBtsRate))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	grpcModel.RegisterConvertorServer(grpcServer, newConvetorServer())
	grpcServer.Serve(lis)
}

func newConvetorServer() grpcModel.ConvertorServer {
	server := &bitcoin.ConversionCryptonator{}
	return server
}
