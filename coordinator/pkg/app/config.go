package app

const localhost = "localhost"

type CommConfig struct {
	Port   int
	Domain string
}

type GRPC struct {
	Subscription CommConfig
	Auth         CommConfig
}

type Config struct {
	GRPC
	REST CommConfig
}

var appConfig Config = Config{
	GRPC: GRPC{
		Subscription: CommConfig{
			Port:   50002,
			Domain: localhost,
		},
		Auth: CommConfig{
			Port:   50000,
			Domain: localhost,
		},
	},
	REST: CommConfig{
		Port:   8080,
		Domain: localhost,
	},
}

func GetConfig() Config {
	return appConfig
}
