package analyzer

import (
	"os"
)

func ReadFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func ReadAndCount(path string) (int, int, error) {
	data, err := ReadFile(path)
	if err != nil {
		return 0, 0, err
	}
	w, l := Count(data)
	return w, l, nil
}