package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "help" {
		fmt.Println("Program requires 2 arguments path and output argument. \n Output arguments: \n -c \t console \n -f \t file \n -cf \t console and file")
		os.Exit(1)
	}
	if len(os.Args) < 3 {
		fmt.Println("Not enought arguments. See help.")
		os.Exit(1)
	}

	if len(os.Args) > 3 {
		fmt.Println("Too many arguments. See help.")
		os.Exit(1)
	}

	filePath := os.Args[1]
	outputArguments := os.Args[2]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	templateData, err := ReadUnmarshalXML(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	templateData.AppendVariables()

	dataMarshalled, err := xml.MarshalIndent(templateData, "	", "	")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}

	switch {
	case outputArguments == "-c":
		fmt.Println(string(dataMarshalled))
	case outputArguments == "-f":
		_ = ioutil.WriteFile("output.xml", dataMarshalled, 0644)
	case outputArguments == "-fc" || outputArguments == "-cf":
		fmt.Println(string(dataMarshalled))
		_ = ioutil.WriteFile("output.xml", dataMarshalled, 0644)
	default:
		fmt.Println(outputArguments, "is not a valid comand. See help.")
	}

}
