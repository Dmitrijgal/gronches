package src

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
	have := "Hi {$author}, your journal: {$journal} kinda boring. We need more {$genreAction} with {$characters.main}!"
	want := "{$author}, {$journal}, {$genreAction}, {$characters.main}"
	got := FindVarsCharBrute(have)
	if got != want {
		t.Errorf("TestFindVars got %s, expected %s", got, want)
	}
}
