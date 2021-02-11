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
// Also it is probably very unefficent because it goes throw every char.
func FindVarsCharBrute(s string) string {

	var result string // returning string
	found := false    // indicator of if var is found and it should be recorded

	// charNum is number of character in recording variable string
	// because variable starts with {$, we checking not only {, but also
	// if second character is $, if not then recording is aborted.
	charNum := 0

	for _, item := range s {

		if item == '{' { //if { found start recording
			found = true
		}
		if found == true {
			charNum++
			if charNum == 2 { //checking is second char is $
				if item != '$' {
					//stop recording and delete already recorded, restart counting
					found = false
					charNum = 0
					result = result[:len(result)-1]
				}
			}
		}
		if item == '}' && found == true { // closing var if } found
			found = false
			charNum = 0
			result += string(item)
			result += ", "
		}

		if found == true { //recording
			result += string(item)
		}
	}

	//Because we adding ", " after every full found variable it will appear even after last one
	// which we dont need. Trims last ", " if any variable {$...} was found
	if result != "" {
		result = result[:len(result)-2]
	}
	return result
}
