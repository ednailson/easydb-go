package arango

import (
	"github.com/arangodb/go-driver"
	"github.com/ednailson/easydb-go"
	"github.com/pkg/errors"
	"strconv"
)

func (t *table) Read(key string) (interface{}, error) {
	var data map[string]interface{}
	_, err := t.coll.ReadDocument(nil, key, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *table) ReadAll() ([]interface{}, error) {
	query := `FOR j IN ` + t.table + ` RETURN j`
	cursor, err := t.db.Query(nil, query, nil)
	if err != nil {
		return nil, err
	}
	return iterateCursor(cursor)
}

func (t *table) Filter(filters easydb.Filters) ([]interface{}, error) {
	query := `FOR u IN @@collection`
	var bindVars = make(map[string]interface{})
	bindVars["@collection"] = t.table
	var i = 0
	for _, value := range filters.Filters {
		v := value.Key
		if bindVars[value.Key] != nil {
			v = value.Key + strconv.Itoa(i)
			i++
		}
		query += ` FILTER ` + `u.` + value.Key + ` ` + value.Operator + ` @` + v
		bindVars[v] = value.Value
	}
	query += ` RETURN u`
	cursor, err := t.db.Query(nil, query, bindVars)
	if err != nil {
		return nil, err
	}
	return iterateCursor(cursor)
}

func iterateCursor(cursor driver.Cursor) ([]interface{}, error) {
	var data []interface{}
	for cursor.HasMore() {
		var document map[string]interface{}
		_, err := cursor.ReadDocument(nil, &document)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read documents")
		}
		data = append(data, document)
	}
	if len(data) == 0 {
		return nil, driver.ArangoError{
			HasError:     true,
			Code:         404,
			ErrorNum:     404,
			ErrorMessage: "not found",
		}
	}
	return data, nil
}
