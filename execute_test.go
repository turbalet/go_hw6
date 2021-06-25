package main

import (
	"errors"
	"reflect"
	"testing"
)

var (
	tasks1 = []func() error{
		getNil,
		getError,
		getNil,
		getError,
		getNil,
		getError,
	}

	tasks2 = []func() error{
		getNil,
		getError,
		getError,
		getError,
		getError,
		getError,
	}

	tasks3 = []func() error{
		getNil,
		getNil,
		getNil,
		getNil,
		getNil,
		getNil,
	}
)

func TestExecute(t *testing.T) {
	expectedErr := errors.New("number of errors is more than E")
	testTable := []struct {
		data     []func() error
		expected error
	}{
		{tasks1, nil},
		{tasks2, expectedErr},
		{tasks3, nil},
	}

	for _, testCase := range testTable {
		err := Execute(testCase.data, 3)
		if !reflect.DeepEqual(testCase.expected, err) {
			t.Errorf("Incorrect result. Expected %v, got %v", testCase.expected, err)
		}
	}
}

func getError() error {
	return errors.New("")
}

func getNil() error {
	return nil
}

func TestExecuteChan(t *testing.T) {

	expectedErr := errors.New("number of errors is more than E")
	testTable := []struct {
		data     []func() error
		expected error
	}{
		{tasks1, nil},
		{tasks2, expectedErr},
		{tasks3, nil},
	}

	for _, testCase := range testTable {
		err := ExecuteChan(testCase.data, 3)
		if !reflect.DeepEqual(testCase.expected, err) {
			t.Errorf("Incorrect result. Expected %v, got %v", testCase.expected, err)
		}
	}

}
