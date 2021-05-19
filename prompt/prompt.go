package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abbyck/husk/execcmd"
)

func Prompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		directory := strings.Split(execcmd.FindPath(), "/")
		fmt.Printf("[%s@%s %s]$ ", execcmd.FindUser(), execcmd.FindHost(), directory[len(directory)-1])

		// Read the keyboad input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		// Handle the execution of the input.
		if err = execcmd.ExecInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
