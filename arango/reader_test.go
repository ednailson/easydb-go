package arango

import (
	"encoding/json"
	"github.com/arangodb/go-driver"
	"github.com/ednailson/easydb-go"
	. "github.com/onsi/gomega"
	"testing"
)

func TestReader(t *testing.T) {
	RegisterTestingT(t)
	db, err := NewDatabase(MockDBConfig())
	Expect(err).ToNot(HaveOccurred())
	table, err := db.Table(MockDBConfig().Collections[0].Name)
	Expect(err).ToNot(HaveOccurred())
	coll := MockCollection(MockDBConfig())
	t.Run("Test read a document", func(t *testing.T) {
		RegisterTestingT(t)
		ReadADocument(coll, table.Reader())
	})
	t.Run("Test read a nonexistent document", func(t *testing.T) {
		RegisterTestingT(t)
		ReadANonexistentDocument(table.Reader())
	})
	t.Run("Test read all documents", func(t *testing.T) {
		RegisterTestingT(t)
		ReadAllDocuments(coll, table.Reader())
	})
	t.Run("Test read all empty", func(t *testing.T) {
		RegisterTestingT(t)
		ReadAllEmpty(table.Reader())
	})
	t.Run("Test filters", func(t *testing.T) {
		RegisterTestingT(t)
		ReadFilters(coll, table.Reader())
	})
}

func ReadADocument(coll driver.Collection, reader easydb.IReader) {
	documentData, err := coll.CreateDocument(nil, getUserMock())
	Expect(err).ToNot(HaveOccurred())
	defer RemoveDocument(coll, documentData.Key)
	data, err := reader.Read(documentData.Key)
	Expect(err).ToNot(HaveOccurred())
	userReceived := getUserFromRead(data)
	Expect(userReceived).ToNot(BeEquivalentTo(""))
	assertUsers(userReceived, getUserMock())
}

func ReadANonexistentDocument(reader easydb.IReader) {
	data, err := reader.Read("wrong key")
	Expect(err).To(HaveOccurred())
	Expect(data).To(BeNil())
}

func ReadAllDocuments(coll driver.Collection, reader easydb.IReader) {
	documentData, err := coll.CreateDocument(nil, getUserMock())
	Expect(err).ToNot(HaveOccurred())
	defer RemoveDocument(coll, documentData.Key)
	documentData, err = coll.CreateDocument(nil, getUserMock2())
	Expect(err).ToNot(HaveOccurred())
	defer RemoveDocument(coll, documentData.Key)
	data, err := reader.ReadAll()
	Expect(err).ToNot(HaveOccurred())
	Expect(len(data)).To(BeEquivalentTo(2))
	assertUsers(getUserFromRead(data[0]), getUserMock())
	assertUsers(getUserFromRead(data[1]), getUserMock2())
}

func ReadAllEmpty(reader easydb.IReader) {
	data, err := reader.ReadAll()
	Expect(err).To(HaveOccurred())
	Expect(data).To(BeNil())
	Expect(reader.Errors().IsNotFound(err)).To(BeTrue())
}

func ReadFilters(coll driver.Collection, reader easydb.IReader) {
	data, err := reader.Filter(nil)
	Expect(data).To(BeNil())
	Expect(err).To(HaveOccurred())
	Expect(reader.Errors().IsNotFound(err)).To(BeTrue())
	documentData, err := coll.CreateDocument(nil, getUserMock())
	Expect(err).ToNot(HaveOccurred())
	defer RemoveDocument(coll, documentData.Key)
	documentData, err = coll.CreateDocument(nil, getUserMock2())
	Expect(err).ToNot(HaveOccurred())
	defer RemoveDocument(coll, documentData.Key)
	data, err = reader.Filter([]easydb.Filter{
		{
			Key:      "email",
			Value:    "email@email.com",
			Operator: Equals,
		},
	})
	Expect(err).ToNot(HaveOccurred())
	Expect(len(data)).To(BeEquivalentTo(1))
	assertUsers(getUserFromRead(data[0]), getUserMock())
}

type user struct {
	Key      string `json:"_key,omitempty"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func assertUsers(user1, user2 user) {
	user1.Key = ""
	user2.Key = ""
	Expect(user1).To(BeEquivalentTo(user2))
}

func getUserFromRead(data interface{}) user {
	j, err := json.Marshal(data)
	Expect(err).ToNot(HaveOccurred())
	var u user
	err = json.Unmarshal(j, &u)
	Expect(err).ToNot(HaveOccurred())
	return u
}

func getUserMock() user {
	return user{
		Name:     "name",
		Username: "email",
		Email:    "email@email.com",
		Password: "password",
	}
}

func getUserMock2() user {
	return user{
		Name:     "name2",
		Username: "email2",
		Email:    "email2@email2.com",
		Password: "pass",
	}
}
