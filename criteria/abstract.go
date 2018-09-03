package criteria

import (
	query "github.com/Masterminds/squirrel"
	"github.com/nathanburkett/graphql-go/selected"
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

func ExtractFieldsFromSelectedFields(fields []selected.SelectedField) []string {
	var extracted []string
	for i := 0; i < len(fields); i++ {
		extracted = append(extracted, fields[i].Name)
	}
	return extracted
}
