package envs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {

	if err := godotenv.Load(); err != nil {

		panic(err)

	}

}

func Get(key string) string {

	value, ok := os.LookupEnv(key)

	if ok {

		return value

	}

	log.Println(key, ":value does not exist or is empty")

	return ""

}
