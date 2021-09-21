package models

import (
	"github.com/crawlab-team/crawlab-core/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Plugin struct {
	Id            primitive.ObjectID         `json:"_id" bson:"_id"`
	Name          string                     `json:"name" bson:"name"`
	Description   string                     `json:"description" bson:"description"`
	Proto         string                     `json:"proto" bson:"proto"`
	Active        bool                       `json:"active" bson:"active"`
	Endpoint      string                     `json:"endpoint" bson:"endpoint"`
	Cmd           string                     `json:"cmd" bson:"cmd"`
	EventKey      entity.PluginEventKey      `json:"event_key" bson:"event_key"`
	InstallType   string                     `json:"install_type" bson:"install_type"`
	InstallUrl    string                     `json:"install_url" bson:"install_url"`
	DeployMode    string                     `json:"deploy_mode" bson:"deploy_mode"`
	UIComponents  []entity.PluginUIComponent `json:"ui_components" bson:"ui_components"`
	UISidebarNavs []entity.PluginUINav       `json:"ui_sidebar_navs" bson:"ui_sidebar_navs"`
	UIAssets      []entity.PluginUIAsset     `json:"ui_assets" bson:"ui_assets"`
	Status        string                     `json:"status" bson:"status"`
	Pid           int                        `json:"pid,omitempty" bson:"pid,omitempty"`
	Error         string                     `json:"error,omitempty" bson:"error,omitempty"`
	Restart       bool                       `json:"restart" bson:"restart"`
}

func (p *Plugin) GetId() (id primitive.ObjectID) {
	return p.Id
}

func (p *Plugin) SetId(id primitive.ObjectID) {
	p.Id = id
}

func (p *Plugin) GetName() (name string) {
	return p.Name
}

func (p *Plugin) SetName(name string) {
	p.Name = name
}

func (p *Plugin) GetInstallUrl() (url string) {
	return p.InstallUrl
}

func (p *Plugin) SetInstallUrl(url string) {
	p.InstallUrl = url
}
