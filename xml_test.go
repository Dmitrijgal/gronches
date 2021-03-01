package main

import (
	"fmt"
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
	if (reflect.DeepEqual(data, TemplateList{}) == true) {
		t.Error("TestReadIfEmpty failed, result shouldn`t be empty. \n Before fixing check testfile. File:", fl)
	}
}

var TestTemplate = Template{
	EmailID:   "1",
	JournalID: "1",
	EmailKey:  "1",
	Subject:   "{$manuscriptId} New Submission",
	Body: `Please, do not reply to this email.

	A new article has been submitted to {$journalTitle}.
	
	Submission URL: {$journalUrl}
	
	Title:
	{$articleTitle}
	
	Corresponding author:
	{$correspondingAuthor}
	
	Authors:
	{$otherAuthors}
	
	Abstract:
	{$articleAbstract}`}

func TestXMLAddVal(t *testing.T) {

	want := TestTemplate
	got := TestTemplate

	want.Variables = "{$articleAbstract}, {$articleTitle}, {$correspondingAuthor}, {$journalTitle}, {$journalUrl}, {$manuscriptId}, {$otherAuthors}"
	got.AppendVariables()

	assert.Equal(t, want, got)

}
