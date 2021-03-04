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
	fmt.Println("Input file name (press Enter for default (templates.xml)):")

	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fileName = strings.Trim(fileName, " \n")

	if fileName == "" {
		fileName = "templates.xml"
	}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	templateData, err := ReadUnmarshalXML(file)
	if err != nil {
		fmt.Println(err)
	}

	templateData.AppendVariables()

	dataMarshalled, err := xml.MarshalIndent(templateData, "", " ")
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println("Chose output: \n 1 for Console \n 2 for new file \n 3 for Console and file")
	choice, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	choice = strings.Trim(choice, " \n")
	fmt.Println(choice)
	switch {
	case choice == "1":
		fmt.Println(string(dataMarshalled))
	case choice == "2":
		_ = ioutil.WriteFile("new_"+fileName, dataMarshalled, 0644)
	case choice == "3":
		fmt.Println(string(dataMarshalled))
		_ = ioutil.WriteFile("new_"+fileName, dataMarshalled, 0644)
	}

}
