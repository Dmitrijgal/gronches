package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input file name:")

	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
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

	newFile, err := xml.MarshalIndent(templateData, "", " ")
	if err != nil {
		fmt.Println(err)
	}

	_ = ioutil.WriteFile("result.xml", newFile, 0644)

}
