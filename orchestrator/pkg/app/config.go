package app

const localhost = "localhost"

type Service struct {
	Port   int
	Domain string
}

type Config struct {
	Service
	Subscription Service
	Auth         Service
}

var appConfig Config = Config{
	Subscription: Service{
		Port:   8001,
		Domain: localhost,
	},
	Auth: Service{
		Port:   8000,
		Domain: localhost,
	},
	Service: Service{
		Port:   8003,
		Domain: localhost,
	},
}

func GetConfig() Config {
	return appConfig
}
