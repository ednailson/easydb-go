package mysql

import "github.com/ednailson/easydb-go"

type table struct {
	name   string
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
