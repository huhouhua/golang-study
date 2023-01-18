package main

import (
	"context"
	"golang-study/grpc/api"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	startGRPCServer(ctx)
}

func startGRPCServer(ctx context.Context) {
	//启动一个端口
	lis, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		log.Fatalf("监听错误：%v", err)
	}
	//创建grpc服务
	server := grpc.NewServer([]grpc.ServerOption{}...)

	//注册服务
	api.RegisterPersonalInfoServer(server, &personalInfoService{
		persons: map[int64]*api.PersonalInformation{},
	})
	go func() {
		select {
		case <-ctx.Done():
			server.Stop()
		}
	}()
	//启动服务
	if err := server.Serve(lis); err != nil {
		log.Fatalf("启动服务错误: %v", err)
	}

}
