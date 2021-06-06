package execcmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func ExecInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supported.
		if len(args) < 2 {
			return errors.New("path required")
		}
		// Change the directory and return the error.
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// Pass the program and the arguments separately.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	// Execute the command and return the error.
	return cmd.Run()
}

func FindHost() string {
	host, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
	}
	return host
}

func FindPath() string {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	split := strings.Split(directory, "/")
	return split[len(split)-1]
}

func FindUser() string {
	user, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return user.Username
}
