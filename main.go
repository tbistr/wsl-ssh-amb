package main

import (
	"os"
	"os/exec"
	"strings"
)

const (
	TARGET_PREFIX = "wsl_"
	FLAG_OPTS     = "46AaCfGgKkMNnqsTtVvXxYy"
)

func main() {
	args := append([]string{}, os.Args[1:]...)
	var c *exec.Cmd

	dest := 0
	for i := 0; i < len(args); i++ {
		a := args[i]
		if strings.HasPrefix(a, "-") {
			opt := strings.TrimPrefix(a, "-")
			// These chars are flag opts.
			// If a is k-v opt, then skip v.
			if !strings.Contains(FLAG_OPTS, opt) {
				i++
			}
		} else {
			dest = i
			break
		}
	}

	if strings.HasPrefix(args[dest], TARGET_PREFIX) {
		args[dest] = strings.TrimPrefix(args[dest], TARGET_PREFIX)
		// [-F path] is configFilepath opt.
		// Remove because it is a Windows path expression.
		for i := 0; i < len(args); i++ {
			if args[i] == "-F" {
				args = append(args[0:i], args[i+2:]...)
			}
		}
		c = exec.Command("wsl", append([]string{"ssh"}, args...)...)
	} else {
		c = exec.Command("ssh", args...)
	}

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	c.Run()

	// err.Error() is already displayed in os.Stderr.
	os.Exit(c.ProcessState.ExitCode())
}
