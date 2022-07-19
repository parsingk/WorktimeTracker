package main

import (
	"github.com/mitchellh/go-ps"
	"log"
	"os"
)

func IsRunningProcess() bool {
	process, err := ps.FindProcess(os.Getpid())
	if err != nil {
		log.Panic(err)
	}

	pslist, err := ps.Processes()
	if err != nil {
		log.Panic(err)
	}

	var cnt = 0
	for _, psn := range pslist {
		if psn.Executable() == process.Executable() {
			cnt++

			if cnt > 1 {
				return true
			}
		}
	}

	return false
}
