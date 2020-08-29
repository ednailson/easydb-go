package mongo

import (
	"github.com/ednailson/easydb-go"
	. "github.com/onsi/gomega"
	"go.mongodb.org/mongo-driver/bson"
	"testing"
)

func TestWriter(t *testing.T) {
	RegisterTestingT(t)
	writer := initWriterTest()
	t.Run("create a document", func(t *testing.T) {
		id := insertDocumentWithWriter(writer)
		defer removeDocument(id)
		Expect(findDocument(id)).To(BeEquivalentTo(mockData()))
	})
	t.Run("update a document", func(t *testing.T) {
		id := insertDocumentWithWriter(writer)
		defer removeDocument(id)
		updated, err := writer.Update(id, mockData2())
		Expect(err).ToNot(HaveOccurred())
		Expect(updated).To(BeEquivalentTo(id))
		Expect(findDocument(id)).To(BeEquivalentTo(mockData2()))
	})
	t.Run("update a nonexistent document", func(t *testing.T) {
		updated, err := writer.Update("nonexistent-id", mockData2())
		Expect(err).To(HaveOccurred())
		Expect(writer.Errors().IsNotFound(err)).To(BeTrue())
		Expect(updated).To(BeNil())
	})
	t.Run("delete a document", func(t *testing.T) {
		id := insertDocumentWithWriter(writer)
		err := writer.Delete(id)
		Expect(err).ToNot(HaveOccurred())
		coll := mockColl()
		result := coll.FindOne(nil, filterId(id))
		Expect(result.Err()).To(HaveOccurred())
		Expect(writer.Errors().IsNotFound(result.Err())).To(BeTrue())
	})
	t.Run("delete a nonexistent document", func(t *testing.T) {
		err := writer.Delete("nonexistent-id")
		Expect(err).ToNot(HaveOccurred())
	})
}

func initWriterTest() easydb.Writer {
	db, err := NewDatabase(MockConfig())
	Expect(err).ToNot(HaveOccurred())
	table, err := db.Table(testCollName)
	Expect(err).ToNot(HaveOccurred())
	coll := mockColl()
	_, err = coll.DeleteMany(nil, bson.D{})
	Expect(err).ToNot(HaveOccurred())
	return table.Writer()
}
