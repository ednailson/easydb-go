package easydb

type IDatabase interface {
	Table(table string) (ITable, error)
	Query(query string) (interface{}, error)
}

type ITable interface {
	Writer() IWriter
	Reader() IReader
}

type IWriter interface {
	Save(data interface{}) (interface{}, error)
	Update(id string, data interface{}) (interface{}, error)
	Delete(id string) error
	IsConflict(err error) bool
}

type IReader interface {
	Read(id string) (interface{}, error)
	ReadAll() ([]interface{}, error)
	Filter(filters []Filter) ([]interface{}, error)
	IsNotFound(err error) bool
}
