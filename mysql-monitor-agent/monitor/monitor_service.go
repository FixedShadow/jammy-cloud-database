package monitor

import "github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/service"

type Service struct{}

func (s *Service) Init() {

}

func (s *Service) Start() {
	go service.CollectToServer()
}
