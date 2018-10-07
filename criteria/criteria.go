package criteria

import (
	"fmt"
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

func New(model data_object.Model, args interface{}, fields []selected.SelectedField) AbstractCriteria {
	builder := query.SelectBuilder{}

	cri := &Criteria{
		builder:   builder.PlaceholderFormat(query.Question),
		modelType: model,
	}

	if cri.determineModelInterpreter(model); cri.Error() != nil {
		return cri
	}

	if cri.interpreter.handleArgs(cri, args); cri.Error() != nil {
		return cri
	}

	if cri.extractFields(fields); cri.Error() != nil {
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
	fields := c.extractFieldsFromSelectedFields(selectedFields)

	var columns []string

	for i := 0; i < len(fields); i++ {
		column, skip, err := c.interpreter.handleField(fields[i])
		if err != nil {
			table := c.modelType.Table()
			c.SetError(fmt.Errorf("%s table %v", table, err))
			return
		}

		if skip {
			continue
		}

		columns = append(columns, column)
	}

	c.Fields(columns...)
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
	if c.err != nil {
		return "", []interface{}(nil), c.err
	}

	return c.builder.ToSql()
}

func (c *Criteria) Error() error {
	return c.err
}

func (c *Criteria) SetError(err error) {
	c.err = err
}

func (c *Criteria) Interpreter() ModelInterpreter {
	return c.interpreter
}
