package file

import (
	"errors"
	"os"
	"testing"
)

func MockOsOpen(name string) (*os.File, error) {
	if name == "invalid/path" {
		return nil, errors.New("Error!")
	}

	f := os.File{}
	return &f, nil
}

func TestReadCsvInvalidPath(t *testing.T) {
	path := "invalid/path"
	_, err := ReadCsv(path)
	if err == nil {
		t.Errorf("Err is %v; want 'Error while opening file 'invalid/path': Error!", err)
	}
}
