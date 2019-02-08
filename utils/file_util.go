package utils

import (
	"io/ioutil"
	"os"
)

func WriteFile(filepath string, b []byte) error {
	return ioutil.WriteFile(filepath, b, 0660)
}

func GetFile(filepath string) (*os.File, error) {
	return os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
}

func DeleteFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}
	err = os.Remove(filepath)
	return err
}
