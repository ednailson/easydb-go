package arango

func (t *table) Save(data interface{}) (interface{}, error) {
	document, err := t.coll.CreateDocument(nil, data)
	if err != nil {
		return nil, err
	}
	return document.Key, nil
}

func (t *table) Update(id string, data interface{}) (interface{}, error) {
	updatedDoc, err := t.coll.UpdateDocument(nil, id, data)
	if err != nil {
		return nil, err
	}
	return updatedDoc.Key, nil
}

func (t *table) Delete(id string) error {
	_, err := t.coll.RemoveDocument(nil, id)
	if err != nil {
		return err
	}
	return nil
}
