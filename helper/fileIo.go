package helper

import (
	"encoding/json"
	"errors"
	"os"
)

func CheckFileExist(dir string) bool {
	_, err := os.Stat(dir)
	return !errors.Is(err, os.ErrNotExist)
}

func EnsureFileExist(dir string) error {
	if !CheckFileExist(dir) {
		file, err := os.Create(dir)
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.Write([]byte("[]")); err != nil {
			return err
		}
	}
	return nil
}

func WriteIntoJson[T any](dir string, v T) error {
	byte, err := json.MarshalIndent(&v, "", " ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(dir, byte, 0644); err != nil {
		return nil
	}
	return nil
}
