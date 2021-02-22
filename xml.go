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

//AppendVariablesToXML is function which receives template as xml, and returns it with extra row.
//Row contains all unique variables which were used in template subject and body.
func AppendVariablesToXML(template Data) Data {
	for i, row := range template.Row {
		variables := FindVariables(row.Subject + row.Body)
		template.Row[i].Variables = strings.Join(variables, ", ")
	}
	return template
}
