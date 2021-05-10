package service

import (
	"github.com/crawlab-team/crawlab-core/interfaces"
	"time"
)

type Option func(svc interfaces.NodeService)

func WithConfigPath(path string) Option {
	return func(svc interfaces.NodeService) {
		svc.SetConfigPath(path)
	}
}

func WithMonitorInterval(duration time.Duration) Option {
	return func(svc interfaces.NodeService) {
		svc2, ok := svc.(interfaces.NodeMasterService)
		if ok {
			svc2.SetMonitorInterval(duration)
		}
	}
}

func WithHeartbeatInterval(duration time.Duration) Option {
	return func(svc interfaces.NodeService) {
		svc2, ok := svc.(interfaces.NodeWorkerService)
		if ok {
			svc2.SetHeartbeatInterval(duration)
		}
	}
}
