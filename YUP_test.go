package main

import (
	"log"
	"reflect"
	"testing"

	yaml "gopkg.in/yaml.v2"
)

const customerTestFile = "./testdata/customer_test.yaml"
const waiverTestFile = "./testdata/waiver_test.yaml"
const remTestFile = "./testdata/remediate_test.yaml"

func TestWaiverFile(t *testing.T) {

	customerGoldDM := cleanDataMap(NewCustomerYaml().getCustomerData(readYaml(customerTestFile)))
	waiverGoldDM := make(map[interface{}]interface{})

	wYaml := NewWaiverYaml(waiverFileName, "")
	if !wYaml.createWaiverData(customerGoldDM) {
		t.Fatalf("failed creating waiver data")
	}

	err := yaml.Unmarshal(readYaml(waiverTestFile), waiverGoldDM)
	if err != nil {
		t.Fatalf("Error unmarshalling data in golden waiver file: %v", err)
	}

	eq := reflect.DeepEqual(wYaml.data, waiverGoldDM)
	if !eq {
		t.Errorf("The generated waiver output file doesn't match the expected waiver file")
	}
}

func TestRemediateFile(t *testing.T) {

	customerGoldDM := cleanDataMap(NewCustomerYaml().getCustomerData(readYaml(customerTestFile)))
	remGoldDM := make(map[interface{}]interface{})
	remGenDM := make(map[interface{}]interface{})

	remYaml := NewRemYaml(remediateFileName, "")
	if !remYaml.createRemData(customerGoldDM) {
		t.Fatalf("failed creating remediate data")
	}

	//convert generated data from yaml.mapslice to map[interface{}]interface{}
	d, err := yaml.Marshal(remYaml.data)
	if err != nil {
		t.Fatalf("Error marshalling generated remediate data: %v", err)
	}
	err = yaml.Unmarshal(d, remGenDM)
	if err != nil {
		log.Fatalf("Error unmarshalling generated remediate data: %v", err)
	}

	err = yaml.Unmarshal(readYaml(remTestFile), remGoldDM)
	if err != nil {
		t.Fatalf("Error unmarshalling data in golden remediate file: %v", err)
	}

	eq := reflect.DeepEqual(remGoldDM, remGenDM)

	if !eq {
		t.Errorf("The generated remediate output file doesn't match the expected remediate file. Expected: %v, \n Generated: %v", remGoldDM, remGenDM)
	}
}
