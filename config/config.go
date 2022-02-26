package config

import "github.com/sakirsensoy/genv"

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
	URI string
}

var Mongo = &mongoConfig{
	URI: genv.Key("MONGO_URI").String(),
}
