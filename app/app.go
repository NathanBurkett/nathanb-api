package app

import (
	"github.com/nathanburkett/graphql-go"
	"github.com/nathanburkett/nathanb-api/data"
)

type Instance struct {
	dataSource *data.Source
	schema     *graphql.Schema
	rootDir    string
}

var application *Instance

func NewInstance() *Instance {
	if application == nil {
		application = &Instance{}
	}

	return application
}

func (i *Instance) DataSource() *data.Source {
	return i.dataSource
}

func (i *Instance) SetDataSource(ds *data.Source) *Instance {
	i.dataSource = ds

	return i
}

func (i *Instance) Schema() *graphql.Schema {
	return i.schema
}

func (i *Instance) SetSchema(schema *graphql.Schema) *Instance {
	i.schema = schema

	return i
}

func (i *Instance) SetRootDir(root string) *Instance {
	i.rootDir = root

	return i
}

func (i *Instance) RootDir() string {
	return i.rootDir
}
