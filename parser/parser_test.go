// ******************************************************************************
// Matronator © 2024.                                                          *
// ******************************************************************************

package parser

import (
	"fmt"
	"testing"
)

func TestParseStringTableDriven(t *testing.T) {
	var tests = []struct {
		str, arg string
		want     string
	}{
		{"Hello, <% nickname %>!", "World", "Hello, World!"},
		{"Hello, <% nickname|upper %>!", "World", "Hello, WORLD!"},
		{"Hello, <% nickname|truncate:2 %>!", "World", "Hello, Wo!"},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("'%s' [arg: %s]", tt.str, tt.arg)
		t.Run(testname, func(t *testing.T) {
			args := Argument{
				"nickname": tt.arg,
			}
			ans := ParseString(tt.str, args)

			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
