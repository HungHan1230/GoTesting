package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Data struct {
	timestamp_d   int64
	total_nodes_d int
}

var has_next bool
var next interface{}

func appendToCSV(data Data) {
	_, err := os.Open("nodes.csv")
	if err != nil {
		// fmt.Println(os.IsNotExist(err)) //true  證明檔案已經存在
		// fmt.Println(err)                //open widuu.go: no such file or directory
		os.Create("nodes.csv")
	}

	var path = "nodes.csv"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	var tocsv [][]string
	tocsv = append(tocsv, []string{strconv.FormatInt(data.timestamp_d, 10), strconv.Itoa(data.total_nodes_d)})
	// tocsv += strconv.FormatInt(data.timestamp_d, 10) + "," + strconv.Itoa(data.total_nodes_d)
	// fmt.Println(tocsv)
	w.WriteAll(tocsv)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

}

func GetApi(nexturl string) {
	res, err := http.Get(nexturl)
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

	next_url := result["next"]
	// fmt.Println("next: ", next_url)
	next = next_url

	if next_url != nil {
		results := result["results"].([]interface{})
		// fmt.Println("results: ", results)
		for _, obj := range results {
			object := obj.(map[string]interface{})
			timestamp := int64(object["timestamp"].(float64))
			total_nodes := int(object["total_nodes"].(float64))
			// var data [][]string
			// data = append(data, []string{timestamp.(string), total_nodes.(string)})

			data := Data{timestamp_d: timestamp, total_nodes_d: total_nodes}
			// data := []string{string(timestamp), string(total_nodes)}
			// fmt.Println(p)

			appendToCSV(data)

			// fmt.Printf("timstamp: %d , total nodes: %d \n", timestamp, total_nodes)
			//change unix timestamp to utc time format
			// unixTimeUTC := time.Unix(timestamp, 0)
			// fmt.Println("unix time stamp in UTC :--->", unixTimeUTC)
		}

	} else {
		has_next = false
	}

}
func main() {
	// other()
	next = "https://bitnodes.io/api/v1/snapshots/"
	has_next = true
	i := 1
	for has_next {
		fmt.Println("Page", i)
		GetApi(next.(string))
		i++
	}

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

//pretty json
// var prettyJSON bytes.Buffer
// error := json.Indent(&prettyJSON, sitemap, "", "\t")
// if error != nil {
// 	fmt.Println("JSON parse error: ", error)
// 	return
// }
// fmt.Println("Pretty Json:", string(prettyJSON.Bytes()))

// func other() {
// 	_, err := os.Open("sample.csv")
// 	if err != nil {
// 		// fmt.Println(os.IsNotExist(err)) //true  證明檔案已經存在
// 		// fmt.Println(err)                //open widuu.go: no such file or directory
// 		os.Create("sample.csv")
// 	}

// 	var path = "sample.csv"
// 	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	var data [][]string
// 	data = append(data, []string{"bruce", "wayne"})
// 	data = append(data, []string{"clark", "kent"})
// 	data = append(data, []string{"hal", "jordan"})

// 	// var data [][]string
// 	// data = append(data, []string{"bruce", "wayne"})
// 	// data = append(data, []string{"clark", "kent"})
// 	// data = append(data, []string{"hal", "jordan"})

// 	w := csv.NewWriter(f)
// 	w.WriteAll(data)

// 	if err := w.Error(); err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Appending succed")
// }
