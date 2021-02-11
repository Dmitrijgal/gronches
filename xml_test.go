package main

import (
	"fmt"
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
	want := "{$author}, {$journal}, {$genreAction}, {$characters.main}"
	got := FindVarsCharBrute(have)
	if got != want {
		t.Errorf("TestFindVars got:\n|%s| \n expected: \n |%s|", got, want)
	}
}
