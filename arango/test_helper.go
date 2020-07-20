package arango

import (
	"crypto/tls"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/ednailson/easydb-go"
	. "github.com/onsi/gomega"
	"strconv"
)

const DBNameTest = "testingDB"
const CollUserTest = "userTestingCollection"
const DBPassTest = "dummyPass"
const DBUserTest = "root"
const DBPortTest = 8529
const DBHostTest = "http://arangodb.service.internal.com.br"

func MockDBConfig() Config {
	return Config{
		Host:     DBHostTest,
		Port:     DBPortTest,
		User:     DBUserTest,
		Password: DBPassTest,
		Database: DBNameTest,
		Collections: []Collection{
			{
				Name:        CollUserTest,
				IndexFields: []string{"username", "email"},
			},
		},
	}
}

func RemoveDocument(coll driver.Collection, key string) {
	_, err := coll.RemoveDocument(nil, key)
	Expect(err).ToNot(HaveOccurred())
}

func MockCollection(config Config) driver.Collection {
	client := MockClient(config)
	db, err := client.Database(nil, config.Database)
	Expect(err).ToNot(HaveOccurred())
	coll, err := db.Collection(nil, config.Collections[0].Name)
	Expect(err).ToNot(HaveOccurred())
	err = coll.Truncate(nil)
	Expect(err).ToNot(HaveOccurred())
	return coll
}

func MockClient(config Config) driver.Client {
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

func MockTestingDB() easydb.IDatabase {
	db, err := NewDatabase(Config{
		Host:     DBHostTest,
		Port:     DBPortTest,
		User:     DBUserTest,
		Password: DBPassTest,
		Database: DBNameTest,
		Collections: []Collection{
			{
				Name:        CollUserTest,
				IndexFields: []string{"username", "email"},
			},
		},
	})
	Expect(err).ToNot(HaveOccurred())
	return db
}

func MockUserTable() easydb.ITable {
	db := MockTestingDB()
	table, err := db.Table(CollUserTest)
	Expect(err).ToNot(HaveOccurred())
	return table
}

func TruncateUserCollection(coll driver.Collection) {
	err := coll.Truncate(driver.WithWaitForSync(nil))
	if driver.IsPreconditionFailed(err) {
		err := coll.Truncate(driver.WithWaitForSync(nil))
		Expect(err).ToNot(HaveOccurred())
	}
}
