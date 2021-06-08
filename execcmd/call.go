package execcmd

import (
	"io"
	"os/exec"
)

func Call(stack []*exec.Cmd, pipes []*io.PipeWriter) (err error) {
	if stack[0].Process == nil {
		if err = stack[0].Start(); err != nil {
			return err
		}
	}
	if len(stack) > 1 {
		if err = stack[1].Start(); err != nil {
			return err
		}
		defer func() {
			if err == nil {
				pipes[0].Close()
				err = Call(stack[1:], pipes[1:])
			} else {
				stack[1].Wait()
			}
		}()
	}
	return stack[0].Wait()
}
