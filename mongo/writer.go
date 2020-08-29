package mongo

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

func (t *table) Save(data interface{}) (interface{}, error) {
	result, err := t.coll.InsertOne(nil, data)
	if err != nil {
		return nil, err
	}
	return strings.ReplaceAll(strings.Split(fmt.Sprintf("%v", result.InsertedID), `"`)[1], `"`, ""), nil
}

func (t *table) Update(id string, data interface{}) (interface{}, error) {
	updated, err := t.coll.UpdateOne(nil, filterId(id), bson.M{"$set": data})
	if err != nil {
		return nil, err
	}
	if updated.MatchedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}
	return id, nil
}

func (t *table) Delete(id string) error {
	_, err := t.coll.DeleteOne(nil, filterId(id))
	return err
}

func filterId(id string) bson.M {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return bson.M{"_id": objectId}
}
