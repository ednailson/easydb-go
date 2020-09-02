package mysql

type Config struct {
	User            string
	Password        string
	Database        string
	ConnMaxLifetime int
	MaxOpenConns    int
	MaxIdleConns    int
	Tables          []Table
}

type Table struct {
	Name    string
	Columns []Column
	FKs     []FK
	Options []string
}

type FK struct {
	Column          string
	TableReference  string
	ColumnReference string
}

type Column struct {
	Name  string
	Type  string
	Rules []string
}
