package app

type AppConfig struct {
	Name string `ini:"name"`
}

type ServerConfig struct {
	Mode string `ini:"mode"`
	Port string `ini:"port"`
	Host string `ini:"host"`
}
