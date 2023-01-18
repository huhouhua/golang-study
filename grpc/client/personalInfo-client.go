package main

import (
	"context"
	"golang-study/grpc/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	//连接服务
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接grpc服务失败！%v", err)
	}
	//创建客户端服务
	c := api.NewPersonalInfoClient(conn)

	personal := &api.PersonalInformation{
		Name: "张三",
		Id:   1,
		Sex:  "男",
		Tall: 1,
	}
	//调用注册方法
	_, err = c.Register(context.TODO(), personal)
	if err != nil {
		log.Fatalf("注册失败 %v", err)
	}

	//调用获取方法
	ret, err := c.Get(context.TODO(), &api.PersonalInformationRequest{Id: personal.Id})
	if err != nil {
		log.Fatalf("获取失败 %v", err)
	}
	log.Printf("信息: %#v \n", ret)

	//调用删除方法
	_, err = c.Remove(context.TODO(), &api.PersonalInformationRequest{Id: personal.Id})
	if err != nil {
		log.Fatalf("删除失败 %v", err)
	}
}
