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

func (t *table) Writer() easydb.IWriter {
	return t
}

func (t *table) Reader() easydb.IReader {
	return t
}

func (t *table) Errors() easydb.IErrors {
	return t.dbDriver
}
