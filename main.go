package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Ejms-template receives templates list file
//and returns it with new variable row.
//Program should be launched with arguments.
//Input file path is mandatory argument, and should be
//right after flagged ("-") arguments.
//Flagged arguments should be used first.
//Flagged arguments:
//	-output {output file path}

func main() {
	output := flag.String("output", "", "Output data in file. Requires output file path.")
	flag.Parse()
	inputFilePath := flag.Arg(0)

	// Opening file with given path.
	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error opening file: \n", err)
		os.Exit(1)
	}
	defer file.Close()

	//Reading (Unmarshalling) file content into structure.
	templateData, err := ReadUnmarshalXML(file)
	if err != nil {
		fmt.Println("Error unmarshalling file: \n", err)
		os.Exit(1)
	}

	//Appending variable row to structure.
	templateData.AppendVariables()

	//Marshalling data back.
	dataMarshalled, err := xml.MarshalIndent(templateData, "", "	")
	if err != nil {
		fmt.Println("Error marshalling data: \n", err)
		os.Exit(1)

	}

	//Replacing marshal replaced runes back.
	var replacer = strings.NewReplacer("&#xA;", "\n", "&#34;", string('"'), "&#39;", "'", "&#x9;", "\t")
	dataMarshalled = []byte(replacer.Replace(string(dataMarshalled)))

	if *output != "" {
		//Writing data in file if argument were written, else display in console.
		err := ioutil.WriteFile(*output, dataMarshalled, 0644)
		if err != nil {
			fmt.Println("Error writing file: \n", err)
			os.Exit(1)
		}
	} else {
		fmt.Println(string(dataMarshalled))
	}

}
