package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//Create the folder if it doesn't exist. Error if the folder exists as a file
func createOrVerifyFolder(folderName string) bool {
	src, err := os.Stat(folderName)

	if os.IsNotExist(err) {
		errFldr := os.MkdirAll(folderName, 0755)
		if errFldr != nil {
			panic(errFldr)
		}
		log.Println("Created new folder")
		return true
	}

	if src.Mode().IsRegular() {
		log.Println(folderName, "already exist as a file!")
		return false
	}

	return true
}

//Reads Yaml file from a given file path
func readYaml(filePath string) []byte {
	//Read the input file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading Yaml file %v, err   #%v ", filePath, err)
	}

	return data
}

//Writes Yaml file to a given file path
func writeYaml(filePath string, data []byte) bool {

	f, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating Yaml file %v, err   #%v ", filePath, err)
		return false
	}
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		log.Printf("Error writing to Yaml file %v, err   #%v ", filePath, err)
	}

	f.Close()
	return true
}

func cleanDataMap(dataMap map[interface{}]interface{}) map[interface{}]interface{} {

	updatedMap := make(map[interface{}]interface{})

	for key, value := range dataMap {
		key = strings.ToLower(strings.TrimSpace(key.(string)))
		switch value.(type) {
		//Nested Map
		case map[interface{}]interface{}:
			updatedMap[key] = cleanDataMap(value.(map[interface{}]interface{}))
		//Nest Array
		case []interface{}:
			updatedMap[key] = cleanDataArray(value.([]interface{}))
		case string:
			value = strings.TrimSpace(value.(string))
			//if key != "id" && key != "provider" && key != "benchmark" {
			//	value = strings.ToLower(value.(string))
			//}
			updatedMap[key] = value
		default:
			updatedMap[key] = value
		}
	}
	return updatedMap
}

func cleanDataArray(anArray []interface{}) []interface{} {

	updatedArray := make([]interface{}, 0, len(anArray))
	for _, value := range anArray {
		switch value.(type) {
		case map[interface{}]interface{}:
			updatedArray = append(updatedArray, cleanDataMap(value.(map[interface{}]interface{})))
		case []interface{}:
			updatedArray = append(updatedArray, cleanDataArray(value.([]interface{})))
		default:
		}
	}
	return updatedArray
}

//ToDo: Remove if not used in final imp
func parseDataMap(dataMap map[interface{}]interface{}) {

	for key, value := range dataMap {
		switch concreteVal := value.(type) {
		//Nested Map
		case map[interface{}]interface{}:
			fmt.Println(key)
			parseDataMap(value.(map[interface{}]interface{}))
		//Nest Array
		case []interface{}:
			fmt.Println(key)
			parseDataArray(value.([]interface{}))
		default:
			fmt.Println(" ", key, ":", concreteVal)
		}
	}
}

//ToDo: Remove if not used in final imp
func parseDataArray(anArray []interface{}) {
	for _, value := range anArray {
		switch concreteVal := value.(type) {
		case map[interface{}]interface{}:
			parseDataMap(value.(map[interface{}]interface{}))
		case []interface{}:
			parseDataArray(value.([]interface{}))
		default:
			fmt.Println("  ", concreteVal)
		}
	}
}
