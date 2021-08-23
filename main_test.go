package main

import "testing"

//another way of writing test
var tests = []struct {
	name     string
	dividend float32
	divisor  float32
	expected float32
	isErr    bool
}{
	{"valid-test", 100, 10, 10, false},
	{"nonvalid-test", 100, 1, 10, false},
}

func TestDivision(t *testing.T) {
	for _, tt := range tests {
		got, err := divide(tt.dividend, tt.divisor)
		if tt.isErr {
			if err == nil {
				t.Error("expected an error but did not get")
			} else {
				if err != nil {
					t.Error("did not expect and error but got one", err.Error())
				}
			}

			if got != tt.expected {
				t.Errorf("expected %f but got %f", tt.expected, got)
			}
		}
	}
}

//simple method of testing
//writing many functions
func TestDivide(t *testing.T) {
	_, err := divide(10, 1)
	if err != nil {
		t.Error("Got an error while running divide test")
	}
}

//go test  **** in command to run the test
//go test -v ***to cheeck each test data
//go test -cover **** to check the coverage of test
//go test -coverprofile=coverage.out && go tool cover -html=coverage.out  ** to check what is covered. but it did not work so we will check it out
