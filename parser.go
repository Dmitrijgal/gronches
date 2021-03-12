package main

import (
	"sort"
	"unicode"
)

// FindVariables is function which searches for variables in given text.
// Variable type looks like {$...}.
func FindVariables(s string) []string {
	// Returning string
	var result []string
	var tempString string
	//Indicator of if variable is found and should be recorded
	found := false
	//Rune position in recording variable string
	charNum := 0
	dotFound := false

	for _, char := range s {
		if char == '{' {
			//If '{' found start recording
			found = true
			tempString = ""
			charNum = 0
		}

		if found == false {
			//Skipping rest if start of variable not found
			continue
		}

		//Increasing character number at which char we are, after finding start ('{')
		charNum++

		//Checking if second char is $
		if charNum == 2 && char != '$' {
			found = false
			continue
		}

		//Character can be dot, but it cant be 3rd character
		if charNum == 3 && char == '.' {
			found = false
			continue
		}

		if charNum > 3 {
			//Characters can be letter, number, dot and '}'
			if !unicode.IsLetter(char) && !unicode.IsNumber(char) && char != '}' && char != '.' {
				found = false
				continue
			}

			//After 3rd char dot is possible, but only one in a row
			if dotFound == true && (char == '.' || char == '}') {
				found = false
				dotFound = false
				continue
			}

			dotFound = char == '.'
			//If recording and found '}', end recording and save result
			if char == '}' {
				found = false
				tempString += string(char)
				result = append(result, tempString)
				continue
			}
		}
		//Recording char in temporary string
		tempString += string(char)
	}

	if result == nil {
		//If no vars were found returning slice with one empty field
		return []string{""}
	}
	//Removing duplicates, sorting and returning result
	result = removeDuplicates(result)
	sort.Strings(result)
	return result
}

// RemoveDuplicates clears slice dublicates. Returns random order!
func removeDuplicates(s []string) []string {
	encountered := map[string]bool{}
	//Create a map of all unique elements.
	for v := range s {
		encountered[s[v]] = true
	}

	//Place all keys from the map into a slice.
	result := []string{}
	for key := range encountered {
		result = append(result, key)
	}
	return result
}
