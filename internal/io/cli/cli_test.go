package cli

import (
	"maps"
	"testing"
)


func TestGetArgsLength(t *testing.T) {
	args := []string{}
	_, err := GetArgs(args)
	if (err == nil) {
		t.Errorf("Err is %v; want 'Not enough arguments. Use <path_to_statement:bank_name> for each file.'", err)
	}
}

func TestGetArgsFormat(t *testing.T) {
	args := []string{"hello:world:!"}
	_, err := GetArgs(args)
	if (err == nil) {
		t.Errorf("Err is %v; want 'Invalid arguments. Use <path_to_statement:bank_name> for each file.'", err)
	}
}

func TestGetArgsPartLength(t *testing.T) {
	args := []string{"hello:"}
	_, err := GetArgs(args)
	if (err == nil) {
		t.Errorf("Err is %v; want 'Blank path or bank name. Use <path_to_statement:bank_name> for each file.'", err)
	}
}

func TestGetArgsInvalidBank(t *testing.T) {
	args := []string{"path:invalidBank"}
	_, err := GetArgs(args)
	if (err == nil) {
		t.Errorf("Err is %v; want 'Bank invalidBank not supported.", err)
	}
}

func TestGetArgsValid(t *testing.T) {
	args := []string{"path1:luminor", "path2:revolut"}
	exp := map[string]string{
		"luminor": "path1",
		"revolut": "path2",
	}
	got, err := GetArgs(args)
	if (err != nil) {
		t.Errorf("Err is %v; but nil wanted", err)
	}

	if !maps.Equal(got, exp) {
		t.Errorf("Got %v; wanted %v ", got, exp)
	}

}