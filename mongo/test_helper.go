package mongo

import (
	"fmt"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const testCollName = "easydb_test"

func MockClient() *mongo.Client {
	client, err := mongo.Connect(nil, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", MockConfig().Host, MockConfig().Port)).SetAuth(options.Credential{
		Username: "root",
		Password: "dummyPass",
	}))
	Expect(err).ToNot(HaveOccurred())
	return client
}

func MockColl() *mongo.Collection {
	return MockClient().Database(MockConfig().Database).Collection(testCollName)
}

func MockConfig() Config {
	return Config{
		Host:     "mongodb.service.com.br",
		Port:     27017,
		Database: "easydb_test",
		Username: "root",
		Password: "dummyPass",
	}
}

func MockData() Data {
	return Data{
		Name: "mock_data",
		Data: Name{
			Name: "mock_data_data",
		},
	}
}

type Data struct {
	Name string `json:"name"`
	Data Name   `json:"data"`
}

type Name struct {
	Name string `json:"name"`
}
