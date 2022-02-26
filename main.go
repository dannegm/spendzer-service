package main

import (
	_ "github.com/sakirsensoy/genv/dotenv/autoload"

	"github.com/kpango/glg"

	"github.com/dannegm/spendzer-service/api"
	"github.com/dannegm/spendzer-service/config"
)

func main() {
	if config.App.Debug {
		glg.Get().SetLevel(glg.DEBG)
	} else {
		glg.Get().SetLevel(glg.FATAL)
	}

	api.Setup(api.ServerConfig{
		Port:  config.Server.Port,
		Debug: config.App.Debug,
	})
}
