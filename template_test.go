package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
