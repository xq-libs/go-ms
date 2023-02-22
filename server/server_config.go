package server

type Config struct {
	Name string `ini:"name"`
	Mode string `ini:"mode"`
	Port string `ini:"port"`
	Host string `ini:"host"`
}
