package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	output := flag.String("output", "", "Output data in file. Requires output file path.")
	flag.Parse()
	inputFilePath := flag.Arg(0)

	file, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Println("Error opening file: \n", err)
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

	var replacer = strings.NewReplacer("&#xA;", "\n", "&#34;", string('"'), "&#39;", "'", "&#x9;", "\t")
	dataMarshalled = []byte(replacer.Replace(string(dataMarshalled)))

	if *output != "" {

		err := ioutil.WriteFile(*output, dataMarshalled, 0644)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		fmt.Println(string(dataMarshalled))
	}

}
