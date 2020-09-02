package mysql

import (
	"database/sql"
	"fmt"
	"github.com/ednailson/easydb-go"
	"github.com/huandu/go-sqlbuilder"
	"time"
)

type driver struct {
	db *sql.DB
}

func NewDatabase(config Config) (easydb.Database, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", config.User, config.Password, config.Database))
	if err != nil {
		return nil, err
	}
	db.SetConnMaxLifetime(time.Minute * time.Duration(config.ConnMaxLifetime))
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	for _, table := range config.Tables {
		ctb := sqlbuilder.NewCreateTableBuilder()
		ctb.CreateTable(fmt.Sprintf("%s.%s", config.Database, table.Name)).IfNotExists()
		for _, column := range table.Columns {
			rules := ""
			for _, rule := range column.Rules {
				rules = fmt.Sprintf("%s %s", rules, rule)
			}
			ctb.Define(column.Name, column.Type, rules)
		}
		for _, fk := range table.FKs {
			ctb.Option(fmt.Sprintf(
				"FOREIGN KEY %s REFERENCES %s.%s (%s)",
				fk.Column,
				config.Database,
				fk.ColumnReference,
				fk.ColumnReference))
		}
		for _, option := range table.Options {
			ctb.Option(option)
		}
		_, err = db.Query(ctb.String())
		if err != nil {
			return nil, err
		}
	}
	return &driver{db: db}, nil
}

func (d *driver) Table(tableName string) (easydb.Table, error) {
	return &table{
		name:   tableName,
		driver: d,
	}, nil
}

func (d *driver) Query(query string) (interface{}, error) {
	rows, err := d.db.Query(query)
	if err != nil {
		return nil, err
	}
	var results []interface{}
	for rows.Next() {
		var result interface{}
		err = rows.Scan(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	return results, nil
}

func (d *driver) Errors() easydb.Errors {
	return d
}

func (d *driver) IsConflict(err error) bool {
	return false
}

func (d *driver) IsNotFound(err error) bool {
	return false
}
