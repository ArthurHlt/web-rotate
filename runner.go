package main

import (
	"os/exec"
	"strings"
)

func defaultUserData(cmd *exec.Cmd) error {
	flagIndex := -1
	for i, arg := range cmd.Args {
		if strings.Contains(arg, "user-data-dir") {
			flagIndex = i
		}
	}
	if flagIndex == -1 {
		return nil
	}
	args := make([]string, len(cmd.Args)-2)
	n := 0
	for i, arg := range cmd.Args {
		if i == flagIndex || i == flagIndex+1 {
			continue
		}
		args[n] = arg
		n++
	}
	cmd.Args = args
	return nil
}
