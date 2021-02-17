package main

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"sort"
	"unicode"
)

func main() {

}

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

// FindVariables is func which searches for variables in given text.
// Variable type looks like {$...}
// Also it is probably very unefficent because it goes throw every char.
func FindVariables(s string) []string {

	var result []string // returning string
	var tempString string
	found := false // indicator of if var is found and it should be recorded

	// charNum is number of character in recording variable string
	// because variable starts with {$, we checking not only {, but also
	// if second character is $, if not then recording is aborted.
	charNum := 0
	dotFound := false

	for _, char := range s {

		if char == '{' { //if { found start recording
			found = true
			tempString = ""
			charNum = 0
		}

		if found == false {
			//skipping rest if start of variable not found
			continue
		}

		// increasing character number at which char we are after finding start ('{')
		charNum++

		//checking if second char is $
		if charNum == 2 {
			if char != '$' {
				//stop recording and delete already recorded, restart counting
				found = false
				continue
			}
		}

		//thought dot is possible, it cant be 3rd character
		if charNum == 3 {
			if char == '.' {
				found = false
				continue
			}
		}
		//characters can be letter, number, dot and }
		if charNum > 3 {
			if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '}' && char != '.' {
				found = false
				continue
			}

			//after 3rd char dot is possible, but only one in a row
			if dotFound == true && (char == '.' || char == '}') {
				found = false
				dotFound = false
				continue
			}

			if char == '.' {
				dotFound = true
			} else {
				dotFound = false
			}

			// if recording and dound } end recording and save result
			if char == '}' {
				found = false
				tempString += string(char)
				result = append(result, tempString)
				continue
			}
		}

		//recording char in temporary string
		tempString += string(char)

	}

	//If no vars were found returning slice with one empty field
	if result == nil {
		return []string{""}
	}
	result = RemoveDuplicates(result)
	sort.Strings(result)

	return result

}

// RemoveDuplicates clears slice dublicates. !!! returns random order
func RemoveDuplicates(s []string) []string {
	encountered := map[string]bool{}
	// Create a map of all unique elements.
	for v := range s {
		encountered[s[v]] = true
	}

	// Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
