package arango

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/ednailson/easydb-go"
	"github.com/pkg/errors"
	"strconv"
)

type dbDriver struct {
	db          driver.Database
	indexFields map[string][]string
}

func NewDatabase(config Config) (easydb.IDatabase, error) {
	dbConn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.Host + ":" + strconv.Itoa(config.Port)},
		TLSConfig: &tls.Config{},
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     dbConn,
		Authentication: driver.BasicAuthentication(config.User, config.Password)})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get a database client")
	}
	dbExists, err := client.DatabaseExists(nil, config.Database)
	if err != nil {
		return nil, errors.Wrap(err, "failed to check if database exists")
	}
	var db driver.Database
	if dbExists {
		db, err = client.Database(nil, config.Database)
		if err != nil {
			return nil, errors.Wrap(err, "failed to get database")
		}
	} else {
		db, err = client.CreateDatabase(nil, config.Database, nil)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create database")
		}
	}
	fields := make(map[string][]string)
	for _, f := range config.Collections {
		fields[f.Name] = f.IndexFields
	}
	return &dbDriver{
		db:          db,
		indexFields: fields,
	}, nil
}

func (d *dbDriver) Table(tableName string) (easydb.ITable, error) {
	coll, err := initCollection(d.db, tableName)
	if err != nil {
		return nil, errors.Wrap(err, "failed to init collection")
	}
	if fields, ok := d.indexFields[tableName]; ok && len(fields) > 0 {
		for _, field := range fields {
			if _, _, err := coll.EnsurePersistentIndex(nil, []string{field}, &driver.EnsurePersistentIndexOptions{Unique: true, Name: field}); err != nil {
				return nil, errors.Wrap(err, "failed to set up unique fields")
			}
		}
	}
	return &table{
		db:       d.db,
		table:    tableName,
		coll:     coll,
		dbDriver: d,
	}, nil
}

func initCollection(db driver.Database, collection string) (driver.Collection, error) {
	exist, err := db.CollectionExists(nil, collection)
	if err != nil {
		return nil, err
	}
	if !exist {
		return db.CreateCollection(nil, collection, nil)
	}
	return db.Collection(nil, collection)
}

func (d *dbDriver) Query(query string) (interface{}, error) {
	cursor, err := d.db.Query(nil, query, nil)
	if err != nil {
		return nil, err
	}
	return iterateCursor(cursor)
}

func (d *dbDriver) Errors() easydb.IErrors {
	return d
}

func (d *dbDriver) IsConflict(err error) bool {
	return driver.IsConflict(err)
}

func (d *dbDriver) IsNotFound(err error) bool {
	return driver.IsNotFound(err)
}
