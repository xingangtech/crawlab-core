package client

import (
	"github.com/crawlab-team/crawlab-core/constants"
	"github.com/crawlab-team/crawlab-core/interfaces"
	"time"
)

type ModelNodeDelegate struct {
	n interfaces.Node
	interfaces.GrpcClientModelDelegate
}

func (d *ModelNodeDelegate) UpdateStatus(active bool, activeTs time.Time, status string) (err error) {
	d.n.SetActive(active)
	d.n.SetActiveTs(activeTs)
	d.n.SetStatus(status)
	return d.Save()
}

func (d *ModelNodeDelegate) UpdateStatusOnline() (err error) {
	return d.UpdateStatus(true, time.Now(), constants.NodeStatusOnline)
}

func (d *ModelNodeDelegate) UpdateStatusOffline() (err error) {
	return d.UpdateStatus(false, time.Now(), constants.NodeStatusOffline)
}

func NewModelNodeDelegate(n interfaces.Node) interfaces.ModelNodeDelegate {
	return &ModelNodeDelegate{
		n:                       n,
		GrpcClientModelDelegate: NewModelDelegate(n),
	}
}
