package src

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

// XMLData is structure for xml file XML data. It countains XMLRow structures.
type XMLData struct {
	XMLName xml.Name `xml:"DATA"`
	Rows    []XMLRow `xml:"Row"`
}

// XMLRow is single row structure.
type XMLRow struct {
	XMLName   xml.Name `xml:"ROW"`
	EmailID   string   `xml:"email_id"`
	JournalID string   `xml:"journal_id"`
	EmailKey  string   `xml:"email_key"`
	Subject   string   `xml:"subject"`
	Body      string   `xml:"body"`
}

// ReadXML is function to read xml file.
func ReadXML(r io.Reader) (template XMLData, err error) {
	xmlFileData, err := ioutil.ReadAll(r)
	if err != nil {
		return template, err
	}
	xml.Unmarshal(xmlFileData, &template)
	return template, nil
}
