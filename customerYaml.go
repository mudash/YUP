package main

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

//CustomerYaml representing the input Yaml file
type CustomerYaml struct {
	customerData map[interface{}]interface{}
}

//NewCustomerYaml creates a new object representing customer Yaml file
func NewCustomerYaml() *CustomerYaml {
	var c CustomerYaml
	c.customerData = make(map[interface{}]interface{})
	return &c
}

//Unmarshall the data read from disk into a map
func (cy CustomerYaml) getCustomerData(data []byte) map[interface{}]interface{} {

	err := yaml.Unmarshal(data, cy.customerData)
	if err != nil {
		log.Fatalf("Error unmarshalling data in customer file: %v", err)
	}

	//cy.customerData = cleanDataMap(cy.customerData)
	return cy.customerData
}

func (cy CustomerYaml) printCustomerData() {

	parseDataMap(cy.customerData)
}
