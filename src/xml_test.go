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
	if (reflect.DeepEqual(data, XMLData{}) == true) {
		t.Error("TestReadIfEmpty failed, result shouldn`t be empty. \n Check that test file isnt empty and exists. File:", fl)
	}
}
