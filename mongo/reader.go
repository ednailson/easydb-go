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
	return result, nil
}

func (t *table) ReadAll() ([]interface{}, error) {
	return t.find(bson.D{})
}

func (t *table) Filter(filters []easydb.Filter) ([]interface{}, error) {
	return t.find(buildFilter(filters))
}

func (t *table) find(filter interface{}) ([]interface{}, error) {
	cur, err := t.coll.Find(nil, filter)
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

func buildFilter(filters []easydb.Filter) bson.M {
	filter := bson.M{}
	for _, value := range filters {
		filter[value.Key] = bson.M{
			value.Operator: value.Value,
		}
	}
	return filter
}
