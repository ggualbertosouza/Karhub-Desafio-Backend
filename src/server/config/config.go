package serverConfig

type EnvConfig struct {
	App   AppConfig
	Db    DbConfig
	Mocks SpotifyMock
}

type AppConfig struct {
	Port        int
	Host        string
	Environment string
}

type DbConfig struct {
	Host     string
	Port     int
	Name     string
	User     string
	Password string
	SSLMode  string
}

type SpotifyMock struct {
	Path string
}
