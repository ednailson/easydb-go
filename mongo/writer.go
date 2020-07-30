package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func (t *table) Save(data interface{}) (interface{}, error) {
	result, err := t.coll.InsertOne(nil, data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (t *table) Update(id string, data interface{}) (interface{}, error) {
	updated, err := t.coll.UpdateOne(nil, filterId(id), data)
	if err != nil {
		return nil, err
	}
	if updated.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}
	return updated.UpsertedID, nil
}

func (t *table) Delete(id string) error {
	_, err := t.coll.DeleteOne(nil, filterId(id))
	return err
}

func filterId(id string) bson.M {
	return bson.M{"_id": id}
}
