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
		t.Error("error: ", err)
	}
	if (reflect.DeepEqual(data, List{}) == true) {
		t.Error("TestReadIfEmpty failed, result shouldn`t be empty. \n Before fixing check testfile. File:", fl)
	}
}
func TestXMLAddVal(t *testing.T) {

	var data List

	err := xml.Unmarshal(sourceTest, &data)
	if err != nil {
		fmt.Println(err)
	}
	want := data
	want.Template[0].Variables = "{$articleAbstract}, {$articleTitle}, {$correspondingAuthor}, {$journalTitle}, {$journalUrl}, {$manuscriptId}, {$otherAuthors}"
	want.Template[1].Variables = "{$articleAbstract}, {$articleAuthors}, {$articleTitle}, {$authorFullName}, {$journalTitle}, {$journalUrl}, {$manuscriptId}, {$submissionTitle}"

	got := data

	got.Template[0].AppendVariables()
	got.Template[1].AppendVariables()

	//FIle is written mainly for self-check
	file, _ := xml.MarshalIndent(data, "", "	")
	_ = ioutil.WriteFile("testdata/templateWithVariableTest.xml", file, 0644)

	assert.Equal(t, want, got)

}

var sourceTest = []byte(`<DATA>

<ROW>
	<email_id>1</email_id>
	<journal_id>1</journal_id>
	<email_key>1</email_key>
	<subject>{$manuscriptId} New Submission</subject>
	<body>Please, do not reply to this email.

A new article has been submitted to {$journalTitle}.

Submission URL: {$journalUrl}

Title:
{$articleTitle}

Corresponding author:
{$correspondingAuthor}

Authors:
{$otherAuthors}

Abstract:
{$articleAbstract}</body>
</ROW>

<ROW>
	<email_id>2</email_id>
	<journal_id>1</journal_id>
	<email_key>2</email_key>
	<subject>{$manuscriptId} Submission Acknowledgment</subject>
	<body>Please, do not reply to this email.

Dear {$authorFullName},

Thank you for your submission of the article "{$submissionTitle}"
(manuscript ID {$manuscriptId}) to {$journalTitle}.

Sincerely,
Sonia Petrone
Editor of Statistical Science
sonia.petrone@unibocconi.it

Submission URL: {$journalUrl}

Title:
{$articleTitle}

Authors:
{$articleAuthors}

Abstract:
{$articleAbstract}</body>
</ROW>
</DATA>
`)
