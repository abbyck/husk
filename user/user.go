package user

import (
	"log"
	"os/user"
)

func FindUser() string {
	user, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	return user.Username
}
