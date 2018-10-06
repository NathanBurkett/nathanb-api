package mock

import (
	"github.com/nathanburkett/nathanb-api/criteria"
)

type Criteria struct {
	ToSqlErr error
}

func (cri *Criteria) From(table string) criteria.AbstractCriteria {
	panic("implement me")
}

func (cri *Criteria) Where(pred interface{}, args ...interface{}) criteria.AbstractCriteria {
	panic("implement me")
}

func (cri *Criteria) OrderBy(clauses []string) criteria.AbstractCriteria {
	panic("implement me")
}

func (cri *Criteria) Limit(limit uint64) criteria.AbstractCriteria {
	panic("implement me")
}

func (cri *Criteria) Offset(limit uint64) criteria.AbstractCriteria {
	panic("implement me")
}

func (cri *Criteria) Fields(fields ...string) criteria.AbstractCriteria {
	panic("implement me")
}

func (cri *Criteria) ToSql() (string, []interface{}, error) {
	return "", []interface{}{""}, cri.ToSqlErr
}

func (cri *Criteria) Error() error {
	panic("implement me")
}

func (cri *Criteria) SetError(error) {
	panic("implement me")
}

func (cri *Criteria) Interpreter() criteria.ModelInterpreter {
	panic("implement me")
}
