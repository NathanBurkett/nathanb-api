package criteria

import (
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/graphql-go/selected"
	"github.com/nathanburkett/nathanb-api/data_object"
)

type Criteria struct {
	builder     query.SelectBuilder
	modelType   data_object.Model
	interpreter ModelInterpreter
	err         error
}

type PaginationArgs struct {
	Limit *uint64
	Page  *uint64
	OrderBy *[]string
	Where *map[string][]WhereClause
}

func New(model data_object.Model, args interface{}, fields []selected.SelectedField) AbstractCriteria {
	cri := &Criteria{
		builder: query.SelectBuilder{}.PlaceholderFormat(query.Question),
	}

	if cri.determineModelInterpreter(model); cri.err != nil {
		return cri
	}

	if cri.interpreter.handleArgs(cri, args); cri.err != nil {
		return cri
	}

	if cri.extractFields(fields); cri.err != nil {
		return cri
	}

	cri.From(model.Table())

	return cri
}

func (c *Criteria) determineModelInterpreter(model data_object.Model) {
	factory := ModelInterpreterFactory{}
	interpreter, err := factory.Create(model)
	if err != nil {
		c.err = err
		return
	}

	c.interpreter = interpreter
}

func (c *Criteria) extractFields(selectedFields []selected.SelectedField) {
	if c.Error() != nil {
		return
	}

	fields := c.extractFieldsFromSelectedFields(selectedFields)

	var columns []string

	for i := 0; i < len(fields); i++ {
		if c.Error() != nil {
			continue
		}

		column, skip, err := c.interpreter.handleField(fields[i])
		if err != nil {
			c.err = err
			return
		}

		if skip {
			continue
		}

		columns = append(columns, column)
	}

	c.builder.Columns(columns...)
}

func (c *Criteria) extractFieldsFromSelectedFields(selectedFields []selected.SelectedField) []string {
	var extracted []string
	for i := 0; i < len(selectedFields); i++ {
		extracted = append(extracted, selectedFields[i].Name)
	}
	return extracted
}

func (c *Criteria) From(table string) AbstractCriteria {
	c.builder = c.builder.From(table)
	return c
}

func (c *Criteria) Where(pred interface{}, args ...interface{}) AbstractCriteria {
	c.builder = c.builder.Where(pred, args...)
	return c
}

func (c *Criteria) OrderBy(clauses []string) AbstractCriteria {
	c.builder = c.builder.OrderBy(clauses...)
	return c
}

func (c *Criteria) Limit(limit uint64) AbstractCriteria {
	c.builder = c.builder.Limit(limit)
	return c
}

func (c *Criteria) Offset(limit uint64) AbstractCriteria {
	c.builder = c.builder.Offset(limit)
	return c
}

func (c *Criteria) Fields(fields ...string) AbstractCriteria {
	c.builder = c.builder.Columns(fields...)
	return c
}

func (c *Criteria) ToSql() (string, []interface{}, error) {
	return c.builder.ToSql()
}

func (c *Criteria) Error() error {
	return c.err
}

func (c *Criteria) SetError(err error) {
	c.err = err
}
