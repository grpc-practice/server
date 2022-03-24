package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"test/services"
)

func main() {
	gwzux := runtime.NewServeMux()
	grpcEndPoint := "localhost:8081"
	opt := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwzux, grpcEndPoint, opt)
	if err != nil {
		log.Fatal(err)
	}
	err = services.RegisterOrderServiceHandlerFromEndpoint(context.Background(), gwzux, grpcEndPoint, opt)
	if err != nil {
		log.Fatal(err)
	}
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwzux,
	}
	httpServer.ListenAndServe()
}
