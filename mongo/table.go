package mongo

import (
	"github.com/ednailson/easydb-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type table struct {
	coll   *mongo.Collection
	driver *driver
}

func (t *table) Writer() easydb.IWriter {
	return t
}

func (t *table) Reader() easydb.IReader {
	return t
}

func (t *table) Errors() easydb.IErrors {
	return t.driver
}
