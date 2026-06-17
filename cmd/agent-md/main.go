package main

import (
	"fmt"
	"os"

	"github.com/dmparrow/agent-md/internal/initcmd"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("usage: agent-md init")
	}

	if args[0] != "init" {
		return fmt.Errorf("usage: agent-md init")
	}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	return initcmd.Init(wd)
}
