package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

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

func TestXMLAddVal(t *testing.T) {
	//At the moment this "test" just outputs file, which I chekc by myself
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

	for _, row := range data.Row {
		row = row.AppendVariables()
	}

	file, _ := xml.MarshalIndent(data, "", "	")
	_ = ioutil.WriteFile("testdata/templateWithVariable.xml", file, 0644)
}
