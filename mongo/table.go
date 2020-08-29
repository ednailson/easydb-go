package mongo

import (
	"github.com/ednailson/easydb-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type table struct {
	coll   *mongo.Collection
	driver *driver
}

func (t *table) Writer() easydb.Writer {
	return t
}

func (t *table) Reader() easydb.Reader {
	return t
}

func (t *table) Errors() easydb.Errors {
	return t.driver
}
