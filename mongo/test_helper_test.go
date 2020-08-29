package mongo

import (
	"encoding/json"
	"fmt"
	"github.com/ednailson/easydb-go"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Name string `json:"name"`
	Data Name   `json:"data"`
}

type Name struct {
	Name string `json:"name"`
}

const testCollName = "easydb_test"

func mockClient() *mongo.Client {
	client, err := mongo.Connect(nil, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", MockConfig().Host, MockConfig().Port)).SetAuth(options.Credential{
		Username: "root",
		Password: "dummyPass",
	}))
	Expect(err).ToNot(HaveOccurred())
	return client
}

func mockColl() *mongo.Collection {
	return mockClient().Database(MockConfig().Database).Collection(testCollName)
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

func mockData() Data {
	return Data{
		Name: "mock_data",
		Data: Name{
			Name: "mock_data_data",
		},
	}
}

func mockData2() Data {
	return Data{
		Name: "mock_data 2",
		Data: Name{
			Name: "mock_data_data 2",
		},
	}
}

func decode(read, data interface{}) {
	body, err := json.Marshal(read)
	Expect(err).ToNot(HaveOccurred())
	err = json.Unmarshal(body, &data)
	Expect(err).ToNot(HaveOccurred())
}

func findDocument(id string) Data {
	coll := mockColl()
	result := coll.FindOne(nil, filterId(id))
	Expect(result.Err()).ToNot(HaveOccurred())
	var testStruct Data
	Expect(result.Decode(&testStruct)).ToNot(HaveOccurred())
	return testStruct
}

func removeDocument(id string) {
	coll := mockColl()
	_, err := coll.DeleteOne(nil, filterId(id))
	Expect(err).ToNot(HaveOccurred())
}

func insertDocumentWithWriter(writer easydb.IWriter) string {
	insertedId, err := writer.Save(mockData())
	Expect(err).ToNot(HaveOccurred())
	Expect(insertedId).ToNot(BeEquivalentTo(""))
	id, ok := insertedId.(string)
	Expect(ok).To(BeTrue())
	return id
}
