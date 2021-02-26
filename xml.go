package main

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"strings"
)

func main() {

}

// List is structure for xml file XML List. It countains Row structures.
type List struct { // rename
	XMLName  xml.Name   `xml:"DATA"`
	Template []Template `xml:"ROW"`
}

// Template is single template structure.
// Represents a sinle row.
type Template struct { //rename
	XMLName   xml.Name `xml:"ROW"`
	EmailID   string   `xml:"email_id"`
	JournalID string   `xml:"journal_id"`
	EmailKey  string   `xml:"email_key"`
	Subject   string   `xml:"subject"`
	Body      string   `xml:"body"`
	Variables string   `xml:"variables"`
}

// ReadXML is function to read xml file.
func ReadXML(r io.Reader) (template List, err error) {
	templateList, err := ioutil.ReadAll(r)
	if err != nil {
		return template, err
	}
	xml.Unmarshal(templateList, &template)
	return template, nil
}

//AppendVariables is function which receives row, and returns it with variables.
//Variables are parsed from subject and body.
func (r *Template) AppendVariables() {
	r.Variables = strings.Join(FindVariables(r.Subject+" "+r.Body), ", ")
}
