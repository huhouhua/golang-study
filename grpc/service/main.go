package main

import (
	"context"
	"golang-study/grpc/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

}

func startGRPCServer(ctx context.Context) {
	//启动一个端口
	lis, err := net.Listen("tcp", "")
	if err != nil {
		log.Fatalf("监听错误：%v", err)
	}
	//创建grpc服务
	server := grpc.NewServer([]grpc.ServerOption{}...)

	//注册服务
	api.RegisterGetPersonalInfoServer(server, &personalInfoService{})

}
