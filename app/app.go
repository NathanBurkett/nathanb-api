package app

import (
	"github.com/nathanburkett/nathanb-api/data"
)

type Instance struct {
	DataSource *data.Source
}

var application *Instance

func NewInstance() *Instance {
	if application == nil {
		application = &Instance{}
	}

	return application
}

func (i *Instance) SetDataSource(ds *data.Source) *Instance {
	i.DataSource = ds

	return i
}
