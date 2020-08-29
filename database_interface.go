package easydb

type Database interface {
	Table(table string) (Table, error)
	Query(query string) (interface{}, error)
	Errors() Errors
}

type Table interface {
	Writer() Writer
	Reader() Reader
}

type Writer interface {
	Save(data interface{}) (interface{}, error)
	Update(id string, data interface{}) (interface{}, error)
	Delete(id string) error
	Errors() Errors
}

type Reader interface {
	Read(id string) (interface{}, error)
	ReadAll() ([]interface{}, error)
	Filter(filters []Filter) ([]interface{}, error)
	Errors() Errors
}

type Errors interface {
	IsConflict(err error) bool
	IsNotFound(err error) bool
}
