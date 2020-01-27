package main

import (
	"log"
	"path/filepath"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

//WaiverYaml representing the output waiver Yaml file
type WaiverYaml struct {
	data map[interface{}]interface{}
	name string
	path string
}

//NewWaiverYaml creates a new object representing customer Yaml file
func NewWaiverYaml(n string, folder string) *WaiverYaml {
	var w WaiverYaml
	w.data = make(map[interface{}]interface{})
	w.name = n
	w.path = filepath.Join(folder, n)
	return &w
}

//Iterates over the input map and constructs a map over
func (wy *WaiverYaml) createWaiverData(cd map[interface{}]interface{}) bool {

	var wd map[interface{}]interface{}
	wd = make(map[interface{}]interface{})

	for key, val := range cd {
		if "controls" == strings.ToLower(key.(string)) {
			switch controlsArray := val.(type) {
			case []interface{}:
				for _, cmap := range controlsArray {
					switch controlsMap := cmap.(type) {
					case map[interface{}]interface{}:
						for cid, cval := range controlsMap {
							// if waiver is present and run is false then it is a waiver
							if "scan" == strings.ToLower(cid.(string)) {
								switch scanMap := cval.(type) {
								case map[interface{}]interface{}:
									if r, ok := scanMap["run"]; ok && !r.(bool) {
										wc := make(map[interface{}]interface{})
										wc["run"] = r
										if j, ok := controlsMap["justification"]; ok {
											wc["justification"] = j
										}
										if i, ok := controlsMap["id"]; ok {
											wd[i] = wc
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	if len(wd) > 0 {
		wy.data = wd
		return true
	}
	return false
}

//Write the waiver Yaml file to disk
func (wy WaiverYaml) writeWaiverYaml() bool {

	d, err := yaml.Marshal(wy.data)
	if err != nil {
		log.Fatalf("error: %v", err)
		return false
	}

	return writeYaml(wy.path, d)
}
