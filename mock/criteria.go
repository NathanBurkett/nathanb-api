package mock

import (
	"github.com/nathanburkett/nathanb-api/criteria"
)

type Criteria struct {
	err      error
	ToSqlErr error
}

func (cri *Criteria) From(table string) criteria.AbstractCriteria {
	return cri
}

func (cri *Criteria) Where(pred interface{}, args ...interface{}) criteria.AbstractCriteria {
	return cri
}

func (cri *Criteria) OrderBy(clauses []string) criteria.AbstractCriteria {
	return cri
}

func (cri *Criteria) Limit(limit uint64) criteria.AbstractCriteria {
	return cri
}

func (cri *Criteria) Offset(limit uint64) criteria.AbstractCriteria {
	return cri
}

func (cri *Criteria) Fields(fields ...string) criteria.AbstractCriteria {
	return cri
}

func (cri *Criteria) ToSql() (string, []interface{}, error) {
	return "", []interface{}{""}, cri.ToSqlErr
}

func (cri *Criteria) Error() error {
	return cri.err
}

func (cri *Criteria) SetError(error) {
	panic("implement me")
}

func (cri *Criteria) Interpreter() criteria.ModelInterpreter {
	panic("implement me")
}
