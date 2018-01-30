package main

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	cases := []struct {
		name  string
		input string
		exp   string
	}{
		{
			name:  "simple input",
			input: "this is a test",
			exp:   "/*****************\n* this is a test *\n*****************/",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			act := generate(c.input)
			equals(t, c.exp, act)
		})
	}
}

func equals(tb testing.TB, exp interface{}, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		tb.Fatalf("\n\texp: %#[1]v (%[1]T)\n\tgot: %#[2]v (%[2]T)", exp, act)
	}
}
