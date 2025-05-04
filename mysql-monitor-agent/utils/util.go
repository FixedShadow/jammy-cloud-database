package utils

import (
	"os"
	"path/filepath"
)

var workingPath string

func GetWorkingPath() string {
	var err error
	if workingPath == "" {
		workingPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {

			return ""
		}
	}
	return workingPath
}
