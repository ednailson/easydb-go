package arango

type Config struct {
	Host        string       `json:"host"`
	Port        int          `json:"port"`
	User        string       `json:"user"`
	Password    string       `json:"password"`
	Database    string       `json:"database"`
	Collections []Collection `json:"collections"`
}

type Collection struct {
	Name        string
	IndexFields []string
}
