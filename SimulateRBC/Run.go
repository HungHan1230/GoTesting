package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// GetNodeSnapshots()
	// plotsnapshots()
	// // getfromTimestamp()
	// // test()
	timestamplist := GetTimestamps()

	for i := 0; i < len(timestamplist); i++ {
		fmt.Println("downlaoding timestamp: ", timestamplist[i])
		if !GetSnapshotsWithTimestamps(timestamplist[i]){
			fmt.Println("request got throttled")
			break
		}
	}

	// mytest()

	// mytest2()
	GetFilesName()
	// readcsv_reverse()

}

func readcsv_reverse() {
	csvfile, err := os.Open("nodes_snapshots.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))
	var slice1, slice2 []string

	//my code

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		slice1 = append(slice1, record[0])
		slice2 = append(slice2, record[1])
	}

	for i := len(slice1) - 1; i <= len(slice1); i-- {
		appendToCSV_pure(slice1[i], slice2[i], "nodes_snapshots_reverse.csv")
	}

}

// func appendToJsonl(nodesjson map[string]json.RawMessage, totalNodes string, filename string) {
// 	outputmap := make(map[string]map[string]string)
// 	// output.total_nodes = totalNodes

// 	for k, v := range nodesjson {
// 		objectmap := make(map[string]string)
// 		//The cool thing is that here I set int slice, it only read the int data into this slice
// 		var temp []int
// 		var tmp []string
// 		json.Unmarshal(v, &temp)
// 		json.Unmarshal(v, &tmp)
// 		// fmt.Println(temp)
// 		// fmt.Println(tmp)
// 		objectmap["start_time"] = strconv.Itoa(temp[2]) // result: s = "-18"temp[2]
// 		objectmap["Timezone"] = tmp[10]
// 		objectmap["ASN"] = tmp[11]
// 		objectmap["Organization"] = tmp[12]
// 		// tmpstr = tmpstr[:len(tmpstr)-1]
// 		// objectmap["info"] = tmpstr
// 		outputmap[k] = objectmap
// 		// fmt.Println(objectmap)
// 	}
// 	// fmt.Println(outputmap)
// 	// fmt.Println(output.total_nodes)
// 	file, err := json.Marshal(outputmap)

// 	if err != nil {
// 		fmt.Println("something wrong while writing json!")
// 	}
// 	var path string = "./node_jsons/" + filename + ".json"
// 	// Write to file
// 	_ = ioutil.WriteFile(path, file, 0644)
// }
// func getfromTimestamp() {
// 	url := "https://bitnodes.io/api/v1/snapshots/1593319209/"
// 	res, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer res.Body.Close()
// 	//request body, which is byte[]
// 	sitemap, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// a map container to decode the JSON structure into
// 	result := make(map[string]json.RawMessage)

// 	// unmarschal JSON
// 	err = json.Unmarshal(sitemap, &result)
// 	if err != nil {
// 		return
// 	}
// 	// fmt.Println(result)
// 	total_nodes := result["total_nodes"]
// 	nodes := result["nodes"]

// 	nodesjson := make(map[string]json.RawMessage)
// 	err = json.Unmarshal(nodes, &nodesjson)
// 	appendToJsonl(nodesjson, string(total_nodes), "1593319209")

// 	//pretty json
// 	// var prettyJSON bytes.Buffer
// 	// error := json.Indent(&prettyJSON, nodes, "", "\t")
// 	// if error != nil {
// 	// 	fmt.Println("JSON parse error: ", error)
// 	// 	return
// 	// }
// 	// fmt.Println("Pretty Json:", string(prettyJSON.Bytes()))

// }

// type UpLoadSomething struct {
// 	Type   string
// 	Object interface{}
// }

// type File struct {
// 	FileName string
// }

// type Png struct {
// 	Wide  int
// 	Hight int
// }
// func test() {

// 	input := `
//     {
//         "type": "File",
//         "object": {
//             "filename": "for test"
// 		}
//     }
//     `
// 	var object json.RawMessage
// 	ss := UpLoadSomething{
// 		Object: &object,
// 	}
// 	if err := json.Unmarshal([]byte(input), &ss); err != nil {
// 		panic(err)
// 	}
// 	switch ss.Type {
// 	case "File":
// 		var f File
// 		if err := json.Unmarshal(object, &f); err != nil {
// 			panic(err)
// 		}
// 		println(f.FileName)
// 	case "Png":
// 		var p Png
// 		if err := json.Unmarshal(object, &p); err != nil {
// 			panic(err)
// 		}
// 		println(p.Wide)
// 	}
// }

// func main() {
// 	// fmt.Println(base)
// 	GetNodeSnapshots()

// }
