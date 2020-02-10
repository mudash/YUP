// YUP - Yaml Unbound Parser
// Usage:
//	go run YUP --input=./sample/customer.yaml --output=./sample/run1
// input flag is mandatory, provide a path to the input Yaml file
// if no output folder is specified then output files are written in the current folder
//

package main

import (
	"flag"
	"log"
)

const waiverFileName = "inspec.yaml"
const remediateFileName = "remediate.yaml"

func main() {

	var input, output string
	flag.StringVar(&input, "input", "", "-input=./sample/customer.yaml ")
	flag.StringVar(&output, "output", "", "-output=./sample/output")

	flag.Parse()

	if input == "" {
		log.Fatalf("Please provide input yaml file. Usage: ./YUP --input=./sample/customer.yaml --output=./sample/output")
	}

	if output == "" {
		log.Println("Output folder is not specified, assuming current directory as output folder")
		output = "."
	}

	if !createOrVerifyFolder(output) {
		log.Fatalf("Cannot create files in specified output folder")
	}

	customerDM := cleanDataMap(NewCustomerYaml().getCustomerData(readYaml(input)))

	//log.Println("Customer Data: ")
	//parseDataMap(customerDM)

	wYaml := NewWaiverYaml(waiverFileName, output)
	rYaml := NewRemYaml(remediateFileName, output)

	if wYaml.createWaiverData(customerDM) {
		wYaml.writeWaiverYaml()
	}

	if rYaml.createRemData(customerDM) {
		rYaml.writeRemYaml()
	}

	log.Println("Done!!")
}
