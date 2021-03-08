package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	filePath := flag.String("path", "templates.xml", "A path to file.")

	flag.Parse()

	file, err := os.Open(*filePath)
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

	var replacer = strings.NewReplacer("&#xA;", "\n", "&#34;", string('"'), "&#39;", "'", "&#x9;", "\t")
	dataMarshalled = []byte(replacer.Replace(string(dataMarshalled)))

	fmt.Println(string(dataMarshalled))

}
