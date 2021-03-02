package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input file name:")

	fileName, _ := reader.ReadString('\n')
	fileName = strings.Trim(fileName, " \n")

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	templateData, err := ReadUnmarshalXML(file)
	if err != nil {
		fmt.Println(err)
	}

	templateData.AppendVariables()

	newFile, _ := xml.MarshalIndent(templateData, "", " ")

	_ = ioutil.WriteFile("result.xml", newFile, 0644)

}

// ReadUnmarshalXML is function to read xml file.
func ReadUnmarshalXML(r io.Reader) (template TemplatesList, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return template, err
	}
	xml.Unmarshal(data, &template)
	return template, nil
}
