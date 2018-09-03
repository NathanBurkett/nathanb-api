package app

import (
	"github.com/nathanburkett/nathanb-api/data"
	"github.com/nathanburkett/graphql-go"
)

type Instance struct {
	DataSource *data.Source
	Schema     *graphql.Schema
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

func (i *Instance) SetSchema(schema *graphql.Schema) *Instance {
	i.Schema = schema

	return i
}
