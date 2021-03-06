package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

//define flags
var inputName = flag.String("name", "XuChao", "Input your name")
var inputAge = flag.Int("age", 25, "Input your age")
var inputGender = flag.String("gender", "boy", "Input your gender")

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
	// nodes := result["nodes"].([]interface{})
	// for _, node := range nodes {
	// 	m := node.(map[string]interface{})
	// 	if name, exists := m["name"]; exists {
	// 		if name == "node1" {
	// 			m["location"] = "new-value1"
	// 		} else if name == "node2" {
	// 			m["location"] = "new-value2"
	// 		}
	// 	}
	// }

	cluster := result["cluster"].(map[string]interface{})
	cluster["secret"] = "123456abc"
	cluster["replication_factor_min"] = 2
	cluster["replication_factor_max"] = 2

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
	flag.Parse() //parse flag
	fmt.Println("name=", *inputName)
	fmt.Println("age=", *inputAge)
	fmt.Println("gender=", *inputGender)

	HandleJson("./service.json", "./service.json")

	// the ways to print out command arguments
	// for idx, args := range os.Args {
	// 	fmt.Println("argument "+strconv.Itoa(idx)+":", args)
	// }
	// fmt.Println(strings.Join(os.Args[1:], "\n"))
	// fmt.Println(os.Args[1:])

}
