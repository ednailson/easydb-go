package easydb

type IDatabase interface {
	Table(table string) (ITable, error)
	Query(query string) (interface{}, error)
	Errors() IErrors
}

type ITable interface {
	Writer() IWriter
	Reader() IReader
}

type IWriter interface {
	Save(data interface{}) (interface{}, error)
	Update(id string, data interface{}) (interface{}, error)
	Delete(id string) error
	Errors() IErrors
}

type IReader interface {
	Read(id string) (interface{}, error)
	ReadAll() ([]interface{}, error)
	Filter(filters []Filter) ([]interface{}, error)
	Errors() IErrors
}

type IErrors interface {
	IsConflict(err error) bool
	IsNotFound(err error) bool
}
