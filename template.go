package main

import (
	"encoding/xml"
	"strings"
)

// Template is single template structure.
// Represents a sinle template row.
type Template struct {
	XMLName   xml.Name `xml:"ROW"`
	EmailID   string   `xml:"email_id"`
	JournalID string   `xml:"journal_id"`
	EmailKey  string   `xml:"email_key"`
	Subject   string   `xml:"subject"`
	Body      string   `xml:"body"`
	Variables string   `xml:"variables,omitempty"`
}

//AppendVariables is function which receives row, and returns it with variables.
//Variables are parsed from subject and body.
func (t *Template) AppendVariables() {
	t.Variables = strings.Join(FindVariables(t.Subject+" "+t.Body), ", ")
}
