package main

import (
	"fmt"
	"log"
	"path/filepath"
	"regexp"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

//RemYaml represents the Yaml file object for remediation file
type RemYaml struct {
	data yaml.MapSlice
	name string
	path string
}

//NewRemYaml creates a new object representing remediate Yaml file
func NewRemYaml(n string, folder string) *RemYaml {
	var r RemYaml
	r.name = n
	r.path = filepath.Join(folder, n)
	return &r
}

//Iterates over the input map and constructs a map over
func (ry *RemYaml) createRemData(cd map[interface{}]interface{}) bool {
	var rd yaml.MapSlice
	var controlsMapSlice []yaml.MapSlice
	var sb, sp string
	rd = make(yaml.MapSlice, 0, len(cd))
	if pv, ok := cd["provider_version"]; ok {
		rd = append([]yaml.MapItem{{Key: "provider_version", Value: pv}}, rd...)
	}
	if b, ok := cd["benchmark"]; ok {
		rd = append([]yaml.MapItem{{Key: "benchmark", Value: b}}, rd...)
		sb = b.(string)
	}
	if p, ok := cd["provider"]; ok {
		rd = append([]yaml.MapItem{{Key: "provider", Value: p}}, rd...)
		sp = p.(string)
	}
	for key, val := range cd { // Loop over complete customer yaml data
		if "controls" == strings.ToLower(key.(string)) {
			switch controlsArray := val.(type) {
			case []interface{}:
				controlsMapSlice = make([]yaml.MapSlice, 0, len(controlsArray))
				for _, cmap := range controlsArray {
					switch controlsMap := cmap.(type) {
					case map[interface{}]interface{}:
						for cid, cval := range controlsMap {
							if "remediate" == strings.ToLower(cid.(string)) {
								switch remMap := cval.(type) {
								case map[interface{}]interface{}:
									controlsMapSubSlice := make(yaml.MapSlice, 0, len(remMap))
									if i, ok := controlsMap["id"]; ok {
										controlsMapSubSlice = append(controlsMapSubSlice,
											yaml.MapItem{Key: "id", Value: transformID(sp, sb, i.(string))})
									}
									if r, ok := remMap["run"]; ok {
										controlsMapSubSlice = append(controlsMapSubSlice, yaml.MapItem{Key: "enabled", Value: r})
										if r.(bool) {
											if w, ok := remMap["waiver"]; ok {
												switch remWaMap := w.(type) {
												case map[interface{}]interface{}:
													wvSubMap := make(yaml.MapSlice, 0, len(remWaMap))
													for x, y := range remWaMap {
														wvSubMap = append(wvSubMap, yaml.MapItem{Key: x, Value: y})
													}
													if j, ok := controlsMap["justification"]; ok {
														wvSubMap = append(wvSubMap, yaml.MapItem{Key: "justification", Value: j})
													}
													controlsMapSubSlice = append(controlsMapSubSlice, yaml.MapItem{Key: "waiver", Value: wvSubMap})
												}
											}
											if o, ok := remMap["overlay_command"]; ok {
												switch remOvMap := o.(type) {
												case []interface{}:
													omSubMapSlice := make(yaml.MapSlice, 0, len(remOvMap))
													omSubArraySlice := make([]yaml.MapSlice, 0, len(remOvMap))
													for _, y := range remOvMap {
														switch remOvMapSubArray := y.(type) {
														case map[interface{}]interface{}:
															for a, b := range remOvMapSubArray {
																omSubMapSlice = append(omSubMapSlice, yaml.MapItem{Key: a, Value: b})
															}
														}
													}
													omSubArraySlice = append(omSubArraySlice, omSubMapSlice)
													controlsMapSubSlice = append(controlsMapSubSlice, yaml.MapItem{Key: "overlay_command", Value: omSubArraySlice})
												}
											}
										}
									}
									controlsMapSlice = append(controlsMapSlice, controlsMapSubSlice)
								}
							}
						}
					}
				}
			}
		}
	}
	rd = append(rd, yaml.MapItem{Key: "controls", Value: controlsMapSlice})

	ry.data = rd
	return true

}

//Write the waiver Yaml file to disk
func (ry RemYaml) writeRemYaml() bool {

	d, err := yaml.Marshal(ry.data)
	if err != nil {
		log.Fatalf("error: %v", err)
		return false
	}

	return writeYaml(ry.path, d)
}

func transformID(provider string, benchmark string, ctrlID string) string {

	re1, err := regexp.Compile(`^[\d\.]+_`)

	if err != nil {
		log.Fatalf("Cannot confirm the Control Id: %s against the regular expression. Error: %s", ctrlID, err)
	}

	subCtrlID := strings.TrimSuffix(re1.FindString(ctrlID), "_")
	subCtrlID = strings.ReplaceAll(subCtrlID, ".", "_")

	provider = strings.ReplaceAll(provider, " ", "_")
	benchmark = strings.ReplaceAll(benchmark, " ", "_")

	return fmt.Sprintf("%s_%s_%s", provider, benchmark, subCtrlID)
}
