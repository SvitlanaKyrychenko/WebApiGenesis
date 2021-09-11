package app

import (
	grpcModel "WebApiGenesis/GRPCMessage/model"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

type BtsRate struct {
	BtsRateServer *string
}

func (btsRate BtsRate) BTSRateHandler(response http.ResponseWriter, request *http.Request) {
	if rate, err := createBtsRateConnection(btsRate.BtsRateServer); err == nil {
		fmt.Fprintf(response, "BTS: 1\nUAH: "+fmt.Sprintf("%.2f", rate))
	}

}

func createBtsRateConnection(serverAddr *string) (float32, error) {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := grpcModel.NewConvertorClient(conn)
	var input grpcModel.BtsRateInput = grpcModel.BtsRateInput{}
	btsRate, err := client.ToUAH(context.Background(), &input)
	if err != nil {
		return -1.0, err
	}
	return btsRate.Rate, nil
}
