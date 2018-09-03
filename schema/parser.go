package schema

import "github.com/nathanburkett/graphql-go"

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) Parse(schema string, resolver interface{}, opts ...graphql.SchemaOpt) (*graphql.Schema, error) {
	return graphql.ParseSchema(schema, resolver, opts...)
}
