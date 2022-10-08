package config

import "os"

var DatabaseURI string
var Port string

func init() {
	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
	DatabaseURI = os.Getenv("DATABASE_URI")
}
