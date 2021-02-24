package main

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"strings"
)

func main() {

}

// Data is structure for xml file XML data. It countains Row structures.
type Data struct {
	XMLName xml.Name `xml:"DATA"`
	Row     []Row    `xml:"ROW"`
}

// Row is single row structure.
type Row struct {
	XMLName   xml.Name `xml:"ROW"`
	EmailID   string   `xml:"email_id"`
	JournalID string   `xml:"journal_id"`
	EmailKey  string   `xml:"email_key"`
	Subject   string   `xml:"subject"`
	Body      string   `xml:"body"`
	Variables string   `xml:"variables"`
}

// ReadXML is function to read xml file.
func ReadXML(r io.Reader) (template Data, err error) {
	xmlFileData, err := ioutil.ReadAll(r)
	if err != nil {
		return template, err
	}
	xml.Unmarshal(xmlFileData, &template)
	return template, nil
}

//AppendVariables is function which receives row, and returns it with variables.
//Variables are parsed from subject and body.
func (r Row) AppendVariables() Row {
	r.Variables = strings.Join(FindVariables(r.Subject+" "+r.Body), ", ")
	return r
}
