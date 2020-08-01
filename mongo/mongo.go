package mongo

import (
	"fmt"
	"github.com/ednailson/easydb-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type driver struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewDatabase(config Config) (easydb.IDatabase, error) {
	client, err := mongo.Connect(nil, options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%d", config.Host, config.Port)).
		SetAuth(options.Credential{
			Username: config.Username,
			Password: config.Password,
		}),
	)
	if err != nil {
		return nil, err
	}
	return &driver{client: client, db: client.Database(config.Database)}, nil
}

func (d *driver) Table(tableName string) (easydb.ITable, error) {
	return &table{coll: d.db.Collection(tableName), driver: d}, nil
}

func (d *driver) Query(query string) (interface{}, error) {
	return nil, nil
}

func (d *driver) Errors() easydb.IErrors {
	return d
}

func (d *driver) IsConflict(err error) bool {
	return false
}

func (d *driver) IsNotFound(err error) bool {
	return err == mongo.ErrNoDocuments
}
