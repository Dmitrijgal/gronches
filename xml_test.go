package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func BenchmarkReadCsv(t *testing.B) {
	fl := "testdata/templates.xml"
	f, err := os.Open(fl)
	if err != nil {
		t.Error("Test failed, cant read test file ", fl)
	}
	defer f.Close()

	_, _ = ReadXML(f)

}

func BenchmarkFindVars(t *testing.B) {
	fl := "testdata/templates.xml"
	f, err := os.Open(fl)
	if err != nil {
		t.Error("Test failed, cant read test file ", fl)
	}
	defer f.Close()
	s, _ := ioutil.ReadAll(f)

	FindVarsCharBrute(string(s))

}

func TestReadIfEmpty(t *testing.T) {
	fl := "testdata/templates.xml"
	f, err := os.Open(fl)
	if err != nil {
		t.Error("Test failed, cant read test file ", fl)
	}
	defer f.Close()

	data, err := ReadXML(f)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	if (reflect.DeepEqual(data, Data{}) == true) {
		t.Error("TestReadIfEmpty failed, result shouldn`t be empty. \n Before fixing check testfile. File:", fl)
	}
}

func TestFindVarsCharBrute(t *testing.T) {
	have := "{test} Hi {$author}, your journal: {test}{$journal}{test} kinda boring. We need more {$genreAction} with {$characters.main}! {test} {tes} {te} {t} {}"
	want := []string{"{$author}", "{$journal}", "{$genreAction}", "{$characters.main}"}
	got := FindVarsCharBrute(have)
	if reflect.DeepEqual(got, want) {
		t.Errorf("TestFindVars got:\n|%s| \n expected: \n |%s|", got, want)
	}
}

/* func TestFindVars(t *testing.T) {
	var tests = []struct {
		have string
		want []string
	}{
		{"Hi {$name}", []string{"{$name}"}},
		{"Hi {$name}", []string{"{$name}"}},
		{"Hi {$name}", []string{"{$name}"}},
		{"Hi {$name}", []string{"{$name}"}},
		{"Hi {$name}", []string{"{$name}"}},
		{"Hi {$name}", []string{"{$name}"}},
		{"Hi {$name}", []string{"{$name}"}},
		{"Hi {$name}", []string{"{$name}"}},
		{"Hi {$name}", []string{"{$name}"}},
	}
	for _, test := range tests {
		got := FindVarsCharBrute(test.have)

		if got != test.want {
			t.Errorf("TestFindVars got:\n|%s| \n expected: \n |%s|", got, want)
		}
	}

}
*/
