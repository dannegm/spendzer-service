package config

import (
	"time"

	"github.com/sakirsensoy/genv"
)

type appConfig struct {
	Debug  bool
	Secret string
}

var App = &appConfig{
	Debug:  genv.Key("APP_DEBUG").Default(false).Bool(),
	Secret: genv.Key("APP_SECRET").String(),
}

type serverConfig struct {
	Port int
}

var Server = &serverConfig{
	Port: genv.Key("SERVER_PORT").Default(3000).Int(),
}

type mongoConfig struct {
	URI      string
	Database string
	Timeout  time.Duration
}

var Mongo = &mongoConfig{
	URI:      genv.Key("MONGO_URI").String(),
	Database: genv.Key("MONGO_DATABASE").Default("spendzer").String(),
	Timeout:  time.Duration(genv.Key("MONGO_TIMEOUT").Default(10).Int()),
}
