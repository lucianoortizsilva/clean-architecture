package configuration

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT_HTTP = 0
)

func LoadEnvHttp() {
	var err error

	err = godotenv.Load("adapters/.env")

	if err != nil {
		log.Fatal(err)
	}

	PORT_HTTP, err = strconv.Atoi(os.Getenv("PORT_HTTP"))
	if err != nil {
		PORT_HTTP = 9000
	}

}
