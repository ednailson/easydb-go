package easydb

type Filter struct {
	Key      string
	Value    interface{}
	Operator string
}

type Filters struct {
	Filters []Filter
}

func NewFilters() *Filters {
	return &Filters{
		Filters: []Filter{},
	}
}

func (f *Filters) AddFilter(key string, value interface{}, operator string) *Filters {
	f.Filters = append(f.Filters, Filter{
		Key:      key,
		Value:    value,
		Operator: operator,
	})
	return f
}

func (f *Filters) Done() Filters {
	return *f
}
