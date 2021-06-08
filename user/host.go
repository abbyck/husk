package user

import (
	"log"
	"os"
)

func FindHost() string {
	host, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return host
}
