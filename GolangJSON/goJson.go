package main

import (
	"encoding/json"
	"io/ioutil"
)


func HandleJson(jsonFile string, outFile string) error {
	// Read json buffer from jsonFile
	byteValue, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}

	// We have known the outer json object is a map, so we defineresult as map.
	// otherwise, result could be defined as slice if outer is an array
	var result map[string]interface{}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return err
	}

	// handle peers
	nodes := result["nodes"].([]interface{})
	for _, node := range nodes {
		m := node.(map[string]interface{})
		if name, exists := m["name"]; exists {
			if name == "node1" {
				m["location"] = "new-value1"
			} else if name == "node2" {
				m["location"] = "new-value2"
			}
		}
	}

	// Convert golang object back to byte
	byteValue, err = json.Marshal(result)
	if err != nil {
		return err
	}

	// Write back to file
	err = ioutil.WriteFile(outFile, byteValue, 0644)
	return err
}

func main() {
	HandleJson("./node.json","./new.json")

}
