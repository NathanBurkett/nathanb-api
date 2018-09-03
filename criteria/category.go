package criteria

import (
	"github.com/nathanburkett/nathanb-api/data_object"
	query "github.com/Masterminds/squirrel"
	"fmt"
	"github.com/satori/go.uuid"
)

const FieldCategoryId = FieldId
const FieldCategoryTitle = "title"
const FieldCategorySlug = "slug"
const FieldCategoryPublications = "publications"
const FieldCategoryCreatedAt = FieldCreatedAt
const FieldCategoryUpdatedAt = FieldUpdatedAt
const FieldCategoryDeletedAt = FieldDeletedAt

type FirstCategoryArgs struct {
	ID    *uuid.UUID
	Title *string
	Slug  *string
}

type categoryInterpretation struct {}

func (ci categoryInterpretation) handleArgs(c *Criteria, args interface{}) {
	if c.err != nil {
		return
	}

	switch T := args.(type) {
	case FirstCategoryArgs:
		ci.interpretFirstCategoryArgs(c, args)
		break
	case PaginationArgs:
		ci.interpretAllCategoryArgs(c, args)
		break
	default:
		c.err = fmt.Errorf("unknown category argument type: %s", T)
	}
}

func (ci categoryInterpretation) interpretFirstCategoryArgs(c *Criteria, args interface{}) {
	firstArgs := args.(FirstCategoryArgs)

	if firstArgs.ID != nil {
		c.Where(query.Eq{data_object.FieldCategoryId: firstArgs.ID})
	}

	if firstArgs.Title != nil {
		c.Where(query.Eq{data_object.FieldCategoryTitle: firstArgs.Title})
	}

	if firstArgs.Slug != nil {
		c.Where(query.Eq{data_object.FieldCategorySlug: firstArgs.Slug})
	}
}

func (ci categoryInterpretation) interpretAllCategoryArgs(c *Criteria, args interface{}) {
	allArgs := args.(PaginationArgs)

	if allArgs.Limit == nil {
		l := uint64(10)
		allArgs.Limit = &l
	}

	if allArgs.Page == nil {
		l := uint64(1)
		allArgs.Page = &l
	}

	if allArgs.OrderBy != nil {
		allArgs.OrderBy = &[]string{"slug DESC"}
	}

	c.Limit(*allArgs.Limit)
	c.Offset(*allArgs.Limit * *allArgs.Page)
	c.OrderBy(*allArgs.OrderBy)

	if allArgs.Where != nil {
		for key, clauses := range *allArgs.Where {
			if c.err != nil {
				return
			}

			column, shouldSkip, err := ci.handleField(key)
			if err != nil {
				c.err = err
				return
			}

			if shouldSkip {
				continue
			}

			newClauses := interpretWhereClauses(column, clauses)
			for i := 0; i < len(newClauses); i++ {
				c.Where(newClauses[i])
			}

		}
	}
}

func (ci categoryInterpretation) handleField(field string) (string, bool, error) {
	var (
		column string
		err error
		skip bool
	)

	switch field {
	case FieldCategoryId:
		column = data_object.FieldCategoryId
	case FieldCategoryTitle:
		column = data_object.FieldCategoryTitle
	case FieldCategorySlug:
		column = data_object.FieldCategorySlug
	case FieldCategoryPublications:
		skip = true
	case FieldCategoryCreatedAt:
		column = data_object.FieldCategoryCreatedAt
	case FieldCategoryUpdatedAt:
		column = data_object.FieldCategoryUpdatedAt
	case FieldCategoryDeletedAt:
		column = data_object.FieldCategoryDeletedAt
	default:
		err = fmt.Errorf("unknown category field: %s", field)
	}

	return column, skip, err
}
