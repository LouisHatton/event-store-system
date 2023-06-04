package appconfig

import "github.com/LouisHatton/insight-wave/internal/config/enviroment"

type Enviroment struct {
	CurrentEnv enviroment.Type `env:"ENVIROMENT" envDefault:"other"`
}

type Server struct {
	Port string `env:"PORT" envDefault:"8080"`
}

type TinyBird struct {
	DatasourcesCreateToken string `env:"TINY_BIRD_DATASOURCES_CREATE"`
	DatasourcesDeleteToken string `env:"TINY_BIRD_DATASOURCES_DELETE"`
	Url                    string `env:"TINY_BIRD_URL"`
}
