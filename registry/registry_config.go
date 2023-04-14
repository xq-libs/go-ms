package registry

type Config struct {
	Client   *ClientConfig
	Instance *InstanceConfig
}

type ClientConfig struct {
	Url             string `ini:"url"`
	RefreshInterval int    `ini:"refreshInterval"`
	TryTimes        int    `ini:"tryTimes"`
}

type InstanceConfig struct {
	Group          string `ini:"group"`
	Service        string `ini:"service"`
	Host           string `ini:"host"`
	Port           int    `ini:"port"`
	HomePageUrl    string `ini:"homepageUrl"`
	HealthCheckUrl string `ini:"healthCheckUrl"`
}
