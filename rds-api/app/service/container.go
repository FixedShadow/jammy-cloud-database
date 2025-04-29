package service

type ContainerService struct{}

type IContainerService interface {
	ContainerCreate() error
}
