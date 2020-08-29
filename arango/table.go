package arango

import (
	"github.com/arangodb/go-driver"
	"github.com/ednailson/easydb-go"
)

type table struct {
	db       driver.Database
	table    string
	coll     driver.Collection
	dbDriver *dbDriver
}

func (t *table) Writer() easydb.Writer {
	return t
}

func (t *table) Reader() easydb.Reader {
	return t
}

func (t *table) Errors() easydb.Errors {
	return t.dbDriver
}
