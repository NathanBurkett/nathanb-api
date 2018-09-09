package criteria

import (
	"fmt"
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/nathanb-api/data_object"
)

type Criteria struct {
	builder  query.SelectBuilder
	modelType data_object.Model
	interpretation InterpretationHandler
	err error
}

type PaginationArgs struct {
	Limit *uint64
	Page  *uint64
	OrderBy *[]string
	Where *map[string][]WhereClause
}

type InterpretationHandler interface {
	handleArgs(AbstractCriteria, interface{})
	handleField(string) (string, bool, error)
}

func New(model data_object.Model, args interface{}, fields []string) AbstractCriteria {
	cri := &Criteria{
		builder: query.SelectBuilder{}.PlaceholderFormat(query.Question),
	}

	if cri.determineModelInterpretation(model); cri.err != nil {
		return cri
	}

	if cri.interpretation.handleArgs(cri, args); cri.err != nil {
		return cri
	}

	if cri.extractFields(fields); cri.err != nil {
		return cri
	}

	cri.From(model.Table())

	return cri
}

func (c *Criteria) determineModelInterpretation(model data_object.Model) {
	switch T := model.(type) {
	case data_object.Category:
		c.interpretation = categoryInterpretation{}
	case data_object.Classification:
		c.interpretation = classificationInterpretation{}
	case data_object.ContentBlock:
		c.interpretation = contentBlockInterpretation{}
	case data_object.Media:
		c.interpretation = mediaInterpretation{}
	case data_object.Profile:
		c.interpretation = profileInterpretation{}
	case data_object.Publication:
		c.interpretation = publicationInterpretation{}
	case data_object.User:
		c.interpretation = userInterpretation{}
	default:
		c.err = fmt.Errorf("unknown data object type: %s", T)
	}
}

func (c *Criteria) extractFields(fields []string) {
	if c.Error() != nil {
		return
	}

	var columns []string

	for i := 0; i < len(fields); i++ {
		if c.Error() != nil {
			continue
		}

		column, skip, err := c.interpretation.handleField(fields[i])
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
