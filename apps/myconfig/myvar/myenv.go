package myvar

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

var ENV_DB_CONNECT string
var ENV_PORT string

func GetMyENV(EnvKey string, defaultValue ...string) string {
	value := os.Getenv(EnvKey)
	log.Info().Msgf("[ENV] %s: %s", EnvKey, value)
	if len(value) == 0 && len(defaultValue) != 0 {
		value = defaultValue[0]
	} else if len(value) == 0 && len(defaultValue) == 0 {
		value = ""
	}
	return value
}

func SetEnv() {

	godotenv.Load(".env")

	ENV_DB_CONNECT = GetMyENV("DB_CONNECT")
	ENV_PORT = GetMyENV("PORT")
}
