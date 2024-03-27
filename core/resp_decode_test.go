package core_test

import (
	"fmt"
	"testing"

	"github.com/vzendara/zre-db/core"
)

func TestDecodeSimpleString(t *testing.T) {
	cases := map[string]string{
		"+OK\r\n":  "OK",
		"+BYE\r\n": "BYE",
	}
	for k, v := range cases {
		value, _ := core.Decode([]byte(k))
		if value != v {
			t.Fail()
		}
	}
}
