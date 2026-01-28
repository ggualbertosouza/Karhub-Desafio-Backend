package serverConfig

type EnvConfig struct {
	App AppConfig
}

type AppConfig struct {
	Port        int
	Host        string
	Environment string
}
