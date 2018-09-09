package mock

type Database struct {
	GetErr    error
	SelectErr error
}

func (db *Database) Get(dest interface{}, query string, args ...interface{}) error {
	return db.GetErr
}

func (db *Database) Select(dest interface{}, query string, args ...interface{}) error {
	return db.SelectErr
}
