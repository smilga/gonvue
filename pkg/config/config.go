package config

import (
	"flag"
)

type AppMode int

const (
	Development AppMode = iota
	Production
)

var dev bool
var jsPath string
var appName string
var port string

func init() {
	flag.BoolVar(&dev, "development", false, "Set to run application in development mode")
	flag.StringVar(&jsPath, "js", "/static/app.js", "Set assets path")
	flag.StringVar(&appName, "name", "Demo App", "Application name")
	flag.StringVar(&port, "port", "80", "Server port")
	flag.Parse()
}

type Config struct {
	Name   string
	Mode   AppMode
	JSPath string
	Port   string
}

func New() *Config {
	mode := Production
	if dev {
		mode = Development
	}

	return &Config{appName, mode, jsPath, port}
}
