package prompt

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/abbyck/husk/execcmd"
	"github.com/abbyck/husk/parser"
	"github.com/abbyck/husk/user"
)

func Prompt() {
	var b bytes.Buffer
	reader := bufio.NewReader(os.Stdin)
	for {
		directory := strings.Split(user.FindPath(), "/")
		fmt.Printf("[%s@%s %s]$ ", user.FindUser(), user.FindHost(), directory[len(directory)-1])

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Tokenize and store as exec.Cmd structures
		commands := parser.Tokenize(input)

		// Handle the execution of the input.
		if err = execcmd.ExecInput(&b, commands); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		io.Copy(os.Stdout, &b)
	}
}
