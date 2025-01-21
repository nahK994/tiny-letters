package app

const localhost = "localhost"

type Service struct {
	Port   int
	Domain string
}

type GRPC struct {
	Subscription Service
	Auth         Service
}

type Config struct {
	GRPC
	REST Service
}

var appConfig Config = Config{
	GRPC: GRPC{
		Subscription: Service{
			Port:   50002,
			Domain: localhost,
		},
		Auth: Service{
			Port:   50000,
			Domain: localhost,
		},
	},
	REST: Service{
		Port:   8080,
		Domain: localhost,
	},
}

func GetConfig() Config {
	return appConfig
}
