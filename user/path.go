package user

import (
	"log"
	"os"
	"strings"
)

func FindPath() string {
	directory, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	split := strings.Split(directory, "/")
	return split[len(split)-1]
}
