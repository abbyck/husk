package execcmd

import (
	"bytes"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func ExecInput(output_buffer *bytes.Buffer, input []string) (err error) {
	var error_buffer bytes.Buffer
	pipe_stack := make([]*io.PipeWriter, len(input)-1)
	buffer := []*exec.Cmd{}
	i := 0

	for j := 0; j < len(input); j++ {
		args := strings.Split(input[j], " ")

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

		cmd := exec.Command(args[0], args[1:]...)

		buffer = append(buffer, cmd)
	}

	for ; i < len(input)-1; i++ {
		stdin_pipe, stdout_pipe := io.Pipe()
		buffer[i].Stdout = stdout_pipe
		buffer[i].Stderr = &error_buffer
		buffer[i+1].Stdin = stdin_pipe
		pipe_stack[i] = stdout_pipe
	}
	buffer[i].Stdout = os.Stdout
	buffer[i].Stderr = &error_buffer
	if err := Call(buffer, pipe_stack); err != nil {
		log.Fatalln((error_buffer.String()), err)
	}

	return err
}
