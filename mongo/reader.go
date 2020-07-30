package mongo

import (
	"github.com/ednailson/easydb-go"
	"go.mongodb.org/mongo-driver/bson"
)

func (t *table) Read(id string) (interface{}, error) {
	var result bson.M
	err := t.coll.FindOne(nil, filterId(id)).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (t *table) ReadAll() ([]interface{}, error) {
	cur, err := t.coll.Find(nil, bson.D{})
	if err != nil {
		return nil, err
	}
	var results []interface{}
	defer cur.Close(nil)
	for cur.Next(nil) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			return nil, err
		}
		results = append(results, result)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return results, nil
}

func (t *table) Filter(filters []easydb.Filter) ([]interface{}, error) {
	return nil, nil
}
