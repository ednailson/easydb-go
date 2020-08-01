package mongo

import (
	"encoding/json"
	"fmt"
	"github.com/ednailson/easydb-go"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	RegisterTestingT(t)
	id, reader := initReaderTest()
	t.Run("read an existent document", func(t *testing.T) {
		read, err := reader.Read(id)
		Expect(err).ToNot(HaveOccurred())
		var data Data
		decode(read, &data)
		Expect(data).To(BeEquivalentTo(MockData()))
	})
	t.Run("read a nonexistent document", func(t *testing.T) {
		read, err := reader.Read("id")
		Expect(err).To(HaveOccurred())
		Expect(err).To(BeEquivalentTo(mongo.ErrNoDocuments))
		Expect(read).To(BeNil())
	})
	t.Run("read all existent documents", func(t *testing.T) {
		read, err := reader.ReadAll()
		Expect(err).ToNot(HaveOccurred())
		var data []Data
		decode(read, &data)
		Expect(len(data)).To(BeEquivalentTo(1))
		Expect(data[0]).To(BeEquivalentTo(MockData()))
	})
}

func decode(read, data interface{}) {
	body, err := json.Marshal(read)
	Expect(err).ToNot(HaveOccurred())
	err = json.Unmarshal(body, &data)
	Expect(err).ToNot(HaveOccurred())
}

func initReaderTest() (string, easydb.IReader) {
	db, err := NewDatabase(MockConfig())
	Expect(err).ToNot(HaveOccurred())
	table, err := db.Table(testCollName)
	Expect(err).ToNot(HaveOccurred())
	coll := MockColl()
	_, err = coll.DeleteMany(nil, bson.D{})
	Expect(err).ToNot(HaveOccurred())
	resultInsert, err := coll.InsertOne(nil, MockData())
	Expect(err).ToNot(HaveOccurred())
	return strings.ReplaceAll(strings.Split(fmt.Sprintf("%v", resultInsert.InsertedID), `"`)[1], `"`, ""), table.Reader()
}
