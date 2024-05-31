package cutils_fs

import (
	"errors"
	"os"
)

func DirExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名的文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
