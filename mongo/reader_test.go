package mongo

import (
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
	defer removeDocument(id)
	t.Run("read an existent document", func(t *testing.T) {
		read, err := reader.Read(id)
		Expect(err).ToNot(HaveOccurred())
		var data Data
		decode(read, &data)
		Expect(data).To(BeEquivalentTo(mockData()))
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
		Expect(data[0]).To(BeEquivalentTo(mockData()))
	})
	t.Run("filter existent documents", func(t *testing.T) {
		read, err := reader.Filter(easydb.NewFilters().AddFilter("name", mockData().Name, Equals).Done())
		Expect(err).ToNot(HaveOccurred())
		var data []Data
		decode(read, &data)
		Expect(len(data)).To(BeEquivalentTo(1))
		Expect(data[0]).To(BeEquivalentTo(mockData()))
		read, err = reader.Filter(easydb.NewFilters().AddFilter("name", mockData2().Name, Equals).Done())
		Expect(err).ToNot(HaveOccurred())
		Expect(len(read)).To(BeEquivalentTo(0))
	})
}

func initReaderTest() (string, easydb.IReader) {
	db, err := NewDatabase(MockConfig())
	Expect(err).ToNot(HaveOccurred())
	table, err := db.Table(testCollName)
	Expect(err).ToNot(HaveOccurred())
	coll := mockColl()
	_, err = coll.DeleteMany(nil, bson.D{})
	Expect(err).ToNot(HaveOccurred())
	resultInsert, err := coll.InsertOne(nil, mockData())
	Expect(err).ToNot(HaveOccurred())
	return strings.ReplaceAll(strings.Split(fmt.Sprintf("%v", resultInsert.InsertedID), `"`)[1], `"`, ""), table.Reader()
}
