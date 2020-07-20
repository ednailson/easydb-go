package arango

import (
	"github.com/arangodb/go-driver"
	"github.com/ednailson/easydb-go"
	. "github.com/onsi/gomega"
	"testing"
)

func TestWriter(t *testing.T) {
	RegisterTestingT(t)
	db, err := NewDatabase(MockDBConfig())
	Expect(err).ToNot(HaveOccurred())
	table, err := db.Table(MockDBConfig().Collections[0].Name)
	Expect(err).ToNot(HaveOccurred())
	coll := MockCollection(MockDBConfig())
	t.Run("create a document", func(t *testing.T) {
		RegisterTestingT(t)
		CreateADocument(coll, table.Writer())
	})
	t.Run("update a document", func(t *testing.T) {
		RegisterTestingT(t)
		UpdateADocument(coll, table.Writer())
	})
	t.Run("update a nonexistent document", func(t *testing.T) {
		RegisterTestingT(t)
		UpdateANonexistentDocument(table.Writer())
	})
	t.Run("remove a document", func(t *testing.T) {
		RegisterTestingT(t)
		RemoveADocument(coll, table.Writer())
	})
	t.Run("remove a nonexistent document", func(t *testing.T) {
		RegisterTestingT(t)
		RemoveANonexistentDocument(table.Writer())
	})
}

func CreateADocument(coll driver.Collection, writer easydb.IWriter) {
	key, err := writer.Save(getUserMock())
	Expect(err).ToNot(HaveOccurred())
	Expect(key).ToNot(BeEquivalentTo(""))
	defer RemoveDocument(coll, key.(string))
	var u user
	_, err = coll.ReadDocument(nil, key.(string), &u)
	Expect(err).ToNot(HaveOccurred())
	assertUsers(getUserMock(), u)
}

func UpdateADocument(coll driver.Collection, writer easydb.IWriter) {
	document, err := coll.CreateDocument(nil, getUserMock())
	Expect(err).ToNot(HaveOccurred())
	defer RemoveDocument(coll, document.Key)
	key, err := writer.Update(document.Key, getUserMock2())
	Expect(err).ToNot(HaveOccurred())
	Expect(key).To(BeEquivalentTo(document.Key))
	var u user
	_, err = coll.ReadDocument(nil, document.Key, &u)
	Expect(err).ToNot(HaveOccurred())
	assertUsers(u, getUserMock2())
}

func UpdateANonexistentDocument(writer easydb.IWriter) {
	key, err := writer.Update("nonexistent key", getUserMock())
	Expect(err).To(HaveOccurred())
	Expect(key).ToNot(BeEquivalentTo(""))
}

func RemoveADocument(coll driver.Collection, writer easydb.IWriter) {
	document, err := coll.CreateDocument(nil, getUserMock())
	Expect(err).ToNot(HaveOccurred())
	err = writer.Delete(document.Key)
	Expect(err).ToNot(HaveOccurred())
	var u user
	_, err = coll.ReadDocument(nil, document.Key, &u)
	Expect(err).To(HaveOccurred())
}

func RemoveANonexistentDocument(writer easydb.IWriter) {
	err := writer.Delete("nonexistent key")
	Expect(err).To(HaveOccurred())
}
