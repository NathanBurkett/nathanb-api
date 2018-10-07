package criteria

import (
	query "github.com/Masterminds/squirrel"
)

const FieldId = "id"
const FieldCreatedAt = "createdAt"
const FieldUpdatedAt = "updatedAt"
const FieldDeletedAt = "deletedAt"

const QualifierEq = "="
const QualifierLT = "<"
const QualifierLTE = "<="
const QualifierGT = ">"
const QualifierGTE = ">="
const QualifierNotEq = "!="

const DirDesc = "DESC"
const DirAsc = "ASC"

const DefaultLimit = 10
const DefaultPage = 1

type WhereClause struct {
	qualifier *string
	value *interface{}
}

func interpretWhereClauses(key string, clauses []WhereClause) []interface{} {
	var retClauses []interface{}

	for i := 0; i < len(clauses); i++ {
		wc := clauses[i]
		q := *wc.qualifier

		switch q {
		case QualifierEq:
			retClauses = append(retClauses, query.Eq{key:  *wc.value})
			break
		case QualifierLT:
			retClauses = append(retClauses, query.Lt{key:  *wc.value})
			break
		case QualifierLTE:
			retClauses = append(retClauses, query.LtOrEq{key:  *wc.value})
			break
		case QualifierGT:
			retClauses = append(retClauses, query.Gt{key:  *wc.value})
			break
		case QualifierGTE:
			retClauses = append(retClauses, query.GtOrEq{key:  *wc.value})
			break
		case QualifierNotEq:
			retClauses = append(retClauses, query.NotEq{key:  *wc.value})
			break
		default:

			retClauses = append(retClauses, query.Eq{key:  *wc.value})
		}
	}

	return retClauses
}

type PaginationArgs struct {
	Limit *uint64
	Page  *uint64
	OrderBy *[]string
	Where *map[string][]WhereClause
}

func interpretPaginationArgs(c AbstractCriteria, args PaginationArgs) {
	c.Limit(*args.Limit)
	if *args.Page > 1 {
		offset := (*args.Limit * *args.Page) - *args.Limit
		c.Offset(offset)
	}

	if *args.OrderBy != nil && len(*args.OrderBy) > 0 {
		c.OrderBy(*args.OrderBy)
	}

	if args.Where == nil {
		return
	}

	for key, clauses := range *args.Where {
		if c.Error() != nil {
			return
		}

		column, shouldSkip, err := c.Interpreter().handleField(key)
		if err != nil {
			c.SetError(err)
			return
		}

		if shouldSkip {
			continue
		}

		// TODO nested loop. Refactor this
		newClauses := interpretWhereClauses(column, clauses)
		for i := 0; i < len(newClauses); i++ {
			c.Where(newClauses[i])
		}
	}
}

func checkDefaultPaginationArgs(args PaginationArgs) PaginationArgs {
	if args.Limit == nil {
		l := uint64(DefaultLimit)
		args.Limit = &l
	}

	if args.Page == nil {
		l := uint64(DefaultPage)
		args.Page = &l
	}

	return args
}

type AbstractCriteria interface {
	From(table string) AbstractCriteria
	Where(pred interface{}, args ...interface{}) AbstractCriteria
	OrderBy(clauses []string) AbstractCriteria
	Limit(limit uint64) AbstractCriteria
	Offset(limit uint64) AbstractCriteria
	Fields(fields ...string) AbstractCriteria
	ToSql() (string, []interface{}, error)
	Error() error
	SetError(error)
	Interpreter() ModelInterpreter
}
