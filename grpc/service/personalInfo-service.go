package main

import (
	"errors"
	"fmt"
	"golang-study/grpc/api"
	context2 "golang.org/x/net/context"
	"log"
	"sync"
)

// 实现gprc定义的服务
var _ api.PersonalInfoServer = &personalInfoService{}

type personalInfoService struct {
	sync.Mutex
	persons map[int64]*api.PersonalInformation
}

func (p *personalInfoService) Register(ctx context2.Context, information *api.PersonalInformation) (*api.Empty, error) {
	p.Lock()
	defer p.Unlock()
	p.persons[information.Id] = information
	log.Printf("注册人的名字:%s \n", information.Name)

	return &api.Empty{}, nil
}

func (p *personalInfoService) Get(ctx context2.Context, request *api.PersonalInformationRequest) (*api.PersonalInformation, error) {
	if request.Id <= 0 {
		return nil, errors.New("Id不能小于0！")
	}
	if p.persons[request.Id] == nil {
		return nil, errors.New(fmt.Sprint("id 为%s不存在"))
	}
	return p.persons[request.Id], nil
}

func (p *personalInfoService) GetAll(ctx context2.Context, empty *api.Empty) (*api.PesrsonalInformationList, error) {
	list := &api.PesrsonalInformationList{
		Items: make([]*api.PersonalInformation, len(p.persons)),
	}
	for _, information := range p.persons {
		list.Items = append(list.Items, information)
	}
	return list, nil
}

func (p *personalInfoService) Remove(ctx context2.Context, request *api.PersonalInformationRequest) (*api.Empty, error) {
	if request.Id <= 0 {
		return nil, errors.New("Id不能小于0！")
	}
	if p.persons[request.Id] == nil {
		return nil, errors.New(fmt.Sprint("id 为%s不存在"))
	}
	p.Lock()
	defer p.Unlock()
	log.Printf("删除的名字:%s \n", p.persons[request.Id].Name)
	delete(p.persons, request.Id)
	return &api.Empty{}, nil
}
