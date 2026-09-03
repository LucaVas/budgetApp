package pkey

import (
	"fmt"
	"testing"
)

func TestParsePrivateKey(t *testing.T) {
	var tests = []struct {
		key    []byte
		errMsg string
	}{
		{[]byte{}, "failed to parse PEM private key"},
		{[]byte{0x00}, "failed to parse PEM private key"},
	}

	for idx, test := range tests {
		t.Run(fmt.Sprintf("%d", idx), func(t *testing.T) {
			ans, e := parsePrivateKey(test.key)
			if ans != nil && test.errMsg != "" {
				t.Errorf("got %v, want %v", ans, test.errMsg)
			} else {
				if e.Error() != test.errMsg {
					t.Errorf("got %v, want %v", e, test.errMsg)
				}
			}
		})
	}
}
