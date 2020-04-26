package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

//define flags
var inputServiceJSON = flag.String("j", "./service.json", "Input the path to service.json")
var inputRep = flag.Int("rep", -1, "Input replication factor")
var inputSecret = flag.String("s", "2123321adsfa", "Input secret value")

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

	cluster := result["cluster"].(map[string]interface{})
	cluster["secret"] = *inputSecret
	cluster["replication_factor_min"] = *inputRep
	cluster["replication_factor_max"] = *inputRep

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
	HandleJson(*inputServiceJSON, *inputServiceJSON)

	//check arguments
	// fmt.Println("service.json=", *inputServiceJSON)
	// fmt.Println("Replication factor=", *inputRep)
	// fmt.Println("secret=", *inputSecret)

	// the ways to print out command arguments
	// for idx, args := range os.Args {
	// 	fmt.Println("argument "+strconv.Itoa(idx)+":", args)
	// }
	// fmt.Println(strings.Join(os.Args[1:], "\n"))
	// fmt.Println(os.Args[1:])

}
