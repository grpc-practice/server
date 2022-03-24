package main

import (
	"google.golang.org/grpc"
	"net"
	"test/services"
)

func main() {
	rpcServer := grpc.NewServer()
	services.RegisterProdServiceServer(rpcServer, &services.ProdService{})
	services.RegisterOrderServiceServer(rpcServer, &services.OrdersService{})
	services.RegisterUserServiceServer(rpcServer, &services.UserService{})
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}
