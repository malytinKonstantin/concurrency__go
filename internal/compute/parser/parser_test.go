package parser

import (
	"reflect"
	"testing"
)

func TestSimpleParser_Parse(t *testing.T) {
	parser := NewParser()

	tests := []struct {
		input       string
		expectedCmd *Command
		expectErr   bool
	}{
		{
			input: "SET key value",
			expectedCmd: &Command{
				Type: SET,
				Args: []string{"key", "value"},
			},
			expectErr: false,
		},
		{
			input: "GET key",
			expectedCmd: &Command{
				Type: GET,
				Args: []string{"key"},
			},
			expectErr: false,
		},
		{
			input:       "DEL key",
			expectedCmd: &Command{Type: DEL, Args: []string{"key"}},
			expectErr:   false,
		},
		{
			input:       "UNKNOWN cmd",
			expectedCmd: nil,
			expectErr:   true,
		},
		{
			input:       "",
			expectedCmd: nil,
			expectErr:   true,
		},
	}

	for _, test := range tests {
		cmd, err := parser.Parse(test.input)
		if test.expectErr && err == nil {
			t.Errorf("Expected error for input '%s', got nil", test.input)
		}
		if !test.expectErr && err != nil {
			t.Errorf("Did not expect error for input '%s', got '%v'", test.input, err)
		}
		if !reflect.DeepEqual(cmd, test.expectedCmd) {
			t.Errorf("For input '%s', expected command '%v', got '%v'", test.input, test.expectedCmd, cmd)
		}
	}
}
