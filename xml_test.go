package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
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

	FindVariables(string(s))
}

func TestRemoveDuplicates(t *testing.T) {
	have := []string{"a", "a"}
	want := []string{"a"}
	got := RemoveDuplicates(have)
	assert.Equal(t, want, got)
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

func TestVariableDublicatesTemp(t *testing.T) {
	var tests = []struct {
		have string
		want []string
	}{
		{"Hi {$.text}", []string{""}},
		{"Hi {$text}, {$text}", []string{"{$text}"}},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, FindVariables(test.have))
	}
}

func TestFindVariables(t *testing.T) {
	var tests = []struct {
		have string
		want []string
	}{
		{"Hi {$.text}", []string{""}},
		{"Hi {$te..xt}", []string{""}},
		{"Hi {$te.xt}", []string{"{$te.xt}"}},
		{"Hi {$text.}", []string{""}},
		{"Hi {$text}, its me {$toster}", []string{"{$text}", "{$toster}"}},
		{"Hi {$text}, {} its me {$toster}", []string{"{$text}", "{$toster}"}},
		{"Hi {$text}, {{its me {$toster}", []string{"{$text}", "{$toster}"}},
		{"Hi {$text}, {${its me {$toster}", []string{"{$text}", "{$toster}"}},
		{"Hi {$text},{$$$} {$its me {$as{$toster}. Your bread is in another castle{$A!", []string{"{$text}", "{$toster}"}},
		{"{test} {$a..} {$aaaaa%} Hi {$author}, your journal: {test}{$journal}{test} kinda boring. We need more {$genreAction} with {$characters.main}! {test} {tes} {te} {t} {}",
			[]string{"{$author}", "{$characters.main}", "{$genreAction}", "{$journal}"}},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, FindVariables(test.have))
	}
}

func TestFindVariablesFromFile(t *testing.T) {
	fl := "testdata/templates.xml"
	f, err := os.Open(fl)
	if err != nil {
		t.Error("Test failed, cant read test file: ", fl)
	}
	text, err := ioutil.ReadAll(f)
	if err != nil {
		t.Error("Test failed, error reading file: ", fl)
	}
	got := FindVariables(string(text))
	want := []string{"{$articleAbstract}", "{$articleAuthors}", "{$articleTitle}", "{$authorFullName}",
		"{$correspondingAuthor}", "{$journalTitle}", "{$journalUrl}", "{$manuscriptId}", "{$otherAuthors}", "{$submissionTitle}"}
	assert.Equal(t, want, got)

}
