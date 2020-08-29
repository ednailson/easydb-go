package arango

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/ednailson/easydb-go"
	. "github.com/onsi/gomega"
	"strconv"
)

const dbNameTest = "testingDB"
const collUserTest = "userTestingCollection"
const dbPassTest = "dummyPass"
const dbUserTest = "root"
const dbPortTest = 8529
const dbHostTest = "http://arangodb.service.internal.com.br"

func mockDBConfig() Config {
	return Config{
		Host:     dbHostTest,
		Port:     dbPortTest,
		User:     dbUserTest,
		Password: dbPassTest,
		Database: dbNameTest,
		Collections: []Collection{
			{
				Name:        collUserTest,
				IndexFields: []string{"username", "email"},
			},
		},
	}
}

func removeDocument(coll driver.Collection, key string) {
	_, err := coll.RemoveDocument(nil, key)
	Expect(err).ToNot(HaveOccurred())
}

func mockCollection(config Config) driver.Collection {
	client := mockClient(config)
	db, err := client.Database(nil, config.Database)
	Expect(err).ToNot(HaveOccurred())
	coll, err := db.Collection(nil, config.Collections[0].Name)
	Expect(err).ToNot(HaveOccurred())
	err = coll.Truncate(nil)
	Expect(err).ToNot(HaveOccurred())
	return coll
}

func mockClient(config Config) driver.Client {
	dbConn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{config.Host + ":" + strconv.Itoa(config.Port)},
		TLSConfig: &tls.Config{},
	})
	Expect(err).ToNot(HaveOccurred())
	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     dbConn,
		Authentication: driver.BasicAuthentication(config.User, config.Password)})
	Expect(err).ToNot(HaveOccurred())
	return client
}

func mockTestingDB() easydb.IDatabase {
	db, err := NewDatabase(Config{
		Host:     dbHostTest,
		Port:     dbPortTest,
		User:     dbUserTest,
		Password: dbPassTest,
		Database: dbNameTest,
		Collections: []Collection{
			{
				Name:        collUserTest,
				IndexFields: []string{"username", "email"},
			},
		},
	})
	Expect(err).ToNot(HaveOccurred())
	return db
}

func mockUserTable() easydb.ITable {
	db := mockTestingDB()
	table, err := db.Table(collUserTest)
	Expect(err).ToNot(HaveOccurred())
	return table
}

func truncateUserCollection(coll driver.Collection) {
	err := coll.Truncate(driver.WithWaitForSync(nil))
	if driver.IsPreconditionFailed(err) {
		err := coll.Truncate(driver.WithWaitForSync(nil))
		Expect(err).ToNot(HaveOccurred())
	}
}
