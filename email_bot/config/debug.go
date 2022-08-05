package config

import (
	"os"

	"github.com/mitchellh/go-ps"
)

func IsDebugging() bool {
	pid := os.Getppid()
	for pid != 0 {
		switch p, err := ps.FindProcess(pid); {
		case err != nil:
			return false
		case p.Executable() == "dlv":
			return true
		default:
			pid = p.PPid()
		}
	}
	return false
}
