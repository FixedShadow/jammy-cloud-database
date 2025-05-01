package service

type InstanceService struct {
}

type IInstanceService interface {
	InitInstance() error
}

func NewInstanceService() IInstanceService {
	return &InstanceService{}
}

func (i *InstanceService) InitInstance() error {
	return nil
}
