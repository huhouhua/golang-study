package main

import (
	"golang-study/grpc/api"
	context2 "golang.org/x/net/context"
	"sync"
)

// 实现gprc定义的服务
var _ api.PersonalInfoServer = &personalInfoService{}

type personalInfoService struct {
	sync.Mutex
	persons map[int64]*api.PersonalInformation
}

func (p *personalInfoService) Register(ctx context2.Context, information *api.PersonalInformation) (*api.Empty, error) {

}

func (p *personalInfoService) Get(ctx context2.Context, request *api.PersonalInformationRequest) (*api.PersonalInformation, error) {

}

func (p *personalInfoService) GetAll(ctx context2.Context, empty *api.Empty) (*api.PesrsonalInformationList, error) {
	//TODO implement me
	panic("implement me")
}

func (p *personalInfoService) Remove(ctx context2.Context, request *api.PersonalInformationRequest) (*api.Empty, error) {
	//TODO implement me
	panic("implement me")
}

//if request.Id <= 0 {
//return nil, errors.New("Id不能小于0！")
//}
//if p.persons[request.Id] == nil {
//return nil, errors.New(fmt.Sprint("id 为%s不存在"))
//}
//return p.persons[request.Id], nil

//p.Lock()
//defer p.Unlock()
//p.persons[information.Id] = information
//return &api.Empty{}, nil
