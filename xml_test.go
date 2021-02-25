package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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
	fl := "testdata/templates.xml"
	f, err := os.Open(fl)
	if err != nil {
		t.Error("Test failed, cant read test file ", fl)
	}
	defer f.Close()

	data, err := ReadXML(f)
	if err != nil {
		t.Error("Error reading file")
	}

	want := data
	want.Row[0].Variables = "{$articleAbstract}, {$articleTitle}, {$correspondingAuthor}, {$journalTitle}, {$journalUrl}, {$manuscriptId}, {$otherAuthors}"
	want.Row[1].Variables = "{$articleAbstract}, {$articleAuthors}, {$articleTitle}, {$authorFullName}, {$journalTitle}, {$journalUrl}, {$manuscriptId}, {$submissionTitle}"

	got := data

	for i, row := range got.Row {
		got.Row[i] = row.AppendVariables()
	}

	//FIle is written mainly for self-check
	file, _ := xml.MarshalIndent(data, "", "	")
	_ = ioutil.WriteFile("testdata/templateWithVariableTest.xml", file, 0644)

	assert.Equal(t, want, got)

}
