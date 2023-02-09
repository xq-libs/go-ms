package database

type DbConfig struct {
	Type     string `ini:"type"`
	Host     string `ini:"host"`
	Port     uint   `ini:"port"`
	Database string `ini:"database"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}
