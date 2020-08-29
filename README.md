# Easydb

A library on go to access easily some databases. The main rule of this library is to follow a unique interface for all
databases that we implemented. So, you will only need to know one library ([easydb](https://github.com/ednailson/easydb-go)) for accessing many databases. 

## Supported databases

* [MongoDB](https://www.mongodb.com/)
* [ArangoDB](https://www.arangodb.com/)

## Getting started

**Download**

```bash
go get github.com/ednailson/easydb-go
```

## Database Interface

The main principle of the library is the [database interface](database_interface.go).

Every implemented database follows this interface after created. 

There is a `New(...params)` function for every database, and all of them there is
their own config struct param. This function will return a [database interface](database_interface.go), 
so after that it will be the same usage for every database.

Check how to use it below.

## Usage

### Creating a database

We will create a mongoDB instance right now but you can created an instance
of any easydb [implemented database](README.md#supported-databases).

```go
config := mongo.Config {
            Host:     "mongodb.service.com.br",
            Port:     27017,
            Database: "easydb_test",
            Username: "root",
            Password: "dummyPass",
}
db, err := mongo.NewDatabase(config)
``` 

### Creating a table/collection

```go
table, err := db.Table("easydb-table-example")
``` 

Now you can write or read from this table

### Writing on a table/collection

```go
writer := table.Writer()
writer.Save(`{"data": "your data"}`)
```

### Reading on a table/collection

```go
reader := table.Reader()
reader.Read("document-id")
```

# Developer

[Junior](https://github.com/ednailson)