package manager

import "github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor"

type ServiceManager struct {
	services map[string]Service
}

func NewServiceManager() *ServiceManager {
	servicesMap := make(map[string]Service)
	return &ServiceManager{services: servicesMap}
}

func (sm *ServiceManager) RegisterService() {
	sm.services["monitorService"] = &monitor.Service{}
}

func (sm *ServiceManager) InitService() {
	for _, service := range sm.services {
		service.Init()
	}
}

func (sm *ServiceManager) StartService() {
	for _, service := range sm.services {
		service.Start()
	}
}
