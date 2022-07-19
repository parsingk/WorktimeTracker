package main

import (
	"github.com/lxn/walk"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var (
	gIcon     *walk.Icon
	iconMutex sync.Mutex
)

func LoadIcon(icoBuf []byte, icoName string) {
	icoFile := filepath.Join(os.TempDir(), icoName)

	var err error

	if _, err = os.Stat(icoName); os.IsNotExist(err) {
		if err = ioutil.WriteFile(icoFile, icoBuf, 0644); err != nil {
			return
		}
	}

	iconMutex.Lock()
	gIcon, _ = walk.NewIconFromFile(icoFile)
	iconMutex.Unlock()
}

func LoadIconFromFile(icoPath string) {
	iconMutex.Lock()
	defer iconMutex.Unlock()
	gIcon, _ = walk.NewIconFromFile(icoPath)
}

func GetIcon() *walk.Icon {
	iconMutex.Lock()
	defer iconMutex.Unlock()

	return gIcon
}
