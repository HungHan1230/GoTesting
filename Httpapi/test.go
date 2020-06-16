package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	res, err := http.Get("https://bitnodes.io/api/v1/snapshots/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//request body, which is byte[]
	// fmt.Printf("%s\n", sitemap)

	// parse json
	var result map[string]interface{}
	err = json.Unmarshal(sitemap, &result)
	if err != nil {
		return
	}

	next_uri := result["next"]
	fmt.Println("next: ", next_uri)
	results := result["results"].([]interface{})
	// fmt.Println("results: ", results)
	for _, obj := range results {
		object := obj.(map[string]interface{})
		timestamp := int64(object["timestamp"].(float64))
		total_nodes := int(object["total_nodes"].(float64))
		fmt.Printf("timstamp: %d , total nodes: %d \n", timestamp, total_nodes)
		unixTimeUTC := time.Unix(timestamp, 0)
		fmt.Println("unix time stamp in UTC :--->", unixTimeUTC)
	}

	//pretty json
	// var prettyJSON bytes.Buffer
	// error := json.Indent(&prettyJSON, sitemap, "", "\t")
	// if error != nil {
	// 	fmt.Println("JSON parse error: ", error)
	// 	return
	// }
	// fmt.Println("Pretty Json:", string(prettyJSON.Bytes()))

}

//My appendix

// //test
// data := []byte(`{"cId" : "A" , "cType" : "English" , "saddr" : { "hsnId" : "C" , "addr" : "中正路12號" } , "persons" : [{"id" : 1 , "name" : "Daniel"},{"id" : 2 , "name" : "Allen"},{"id" : 3 , "name" : "Sam"}]}`)
// var jsonObj map[string]interface{}
// json.Unmarshal([]byte(data), &jsonObj)
// classID := jsonObj["cId"].(string)
// classType := jsonObj["cType"].(string)

// fmt.Println(classID)
// fmt.Println(classType)

// studentsAddr := jsonObj["saddr"].(map[string]interface{})
// hsnID := studentsAddr["hsnId"].(string)
// addr := studentsAddr["addr"].(string)

// fmt.Println(hsnID)
// fmt.Println(addr)

// persons := jsonObj["persons"].([]interface{})
// for _, p := range persons {
// 	person := p.(map[string]interface{})
// 	id := int(person["id"].(float64))
// 	name := person["name"].(string)
// 	fmt.Printf("%d , %v \n", id, name)
// }

// for _, obj := range results {
// 	object := obj.(map[string]interface{})
// 	id := int(object["id"].(float64))
// 	name := object["name"].(string)
// 	fmt.Printf("%d , %v \n", id, name)
// }

// var results []string
// _ = json.Unmarshal([]byte(result["results"]), &results)
// log.Printf("Unmarshaled: %v", results)
