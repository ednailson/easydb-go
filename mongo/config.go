package mongo

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}
