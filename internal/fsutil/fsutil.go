package fsutil

import (
	"errors"
	"os"
)

func EnsureDir(path string) error {
	return os.MkdirAll(path, 0o755)
}

func WriteFileIfMissing(path string, content []byte, perm os.FileMode) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, perm)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil
		}

		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	return err
}
