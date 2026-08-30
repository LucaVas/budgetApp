package cli

import (
	"errors"
	"fmt"
	"strings"
)

type Bank int

const (
	Luminor Bank = iota
	Revolut
)

var bankName = map[string]Bank{
	"luminor": Luminor,
	"revolut": Revolut,
}

const msg = "Use <path_to_statement:bank_name> for each file."

func GetArgs(args []string) (map[string]string, error) {
	m := make(map[string]string)

	if len(args) < 1 {
		return m, errors.New("Not enough arguments. " + msg)
	}

	for _, arg := range args {
		sp := strings.Split(arg, ":")
		if (len(sp)) != 2 {
			return m, errors.New("Invalid argument. " + msg)
		}

		for _, s := range sp {
			if len(strings.TrimSpace(s)) == 0 {
				return m, errors.New("Blank path or bank name. " + msg)
			}
		}

		path := sp[0]
		b := sp[1]
		_, ok := bankName[b]

		if !ok {
			return m, fmt.Errorf("Bank '%v' not supported. ", b)
		}

		m[b] = path
	}

	return m, nil
}
