package src

import (
	"encoding/xml"
	"io"
	"io/ioutil"
)

// Data is structure for xml file XML data. It countains Row structures.
type Data struct {
	XMLName xml.Name `xml:"DATA"`
	Rows    []Row    `xml:"ROW"`
}

// Row is single row structure.
type Row struct {
	XMLName   xml.Name `xml:"ROW"`
	EmailID   string   `xml:"email_id"`
	JournalID string   `xml:"journal_id"`
	EmailKey  string   `xml:"email_key"`
	Subject   string   `xml:"subject"`
	Body      string   `xml:"body"`
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

// FindVarsCharBrute is func which searches for variables in given text.
// Variable type looks like {$...}
//It kinda works, but definately atm with crash at element like {...}
// > without $ (searching for way to fix).
// Also it is probably very unefficent because it goes throw every char.
func FindVarsCharBrute(s string) string {

	var result string // returning string
	found := false    // indicator of if var is found and it should be recorded

	for _, item := range s {

		if item == '{' {
			found = true
		}

		if item == '}' {
			found = false
			result += string(item)
			result += ", "
		}

		if found == true {
			result += string(item)
		}
	}

	//Because we adding ", " after every full found variable it will appear even after last one
	// which we dont need. Next if trims last ", " if any variable was found
	if result != "" {
		result = result[:len(result)-2]
	}
	return result
}
