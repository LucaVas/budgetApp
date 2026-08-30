package file

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

var osOpen = os.Open

func ReadCsv(path string) ([]Row, error) {
	p := filepath.Join(path)
	f, err := osOpen(p)
	if err != nil {
		return nil, fmt.Errorf("Error while opening file '%v': %v", path, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Error while reading file '%v': %v", path, err)
	}

	rows := toRows(records)
	return rows, nil
}

type Row struct {
	idx int
	row map[string]string
}

func toRows(records [][]string) []Row {
	rows := []Row{}
	headers := []string{}
	for n, rec := range records {
		if n == 0 {
			headers = rec
			continue
		}
		row := make(map[string]string)
		for idx, el := range rec {
			row[headers[idx]] = el
		}
		rows = append(rows, Row{n, row})
	}

	return rows
}
