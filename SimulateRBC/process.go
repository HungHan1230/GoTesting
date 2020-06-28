package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var globalstate map[string]string

func GetFilesName() {
	files, err := ioutil.ReadDir("../../node_jsons")
	if err != nil {
		log.Fatal(err)
	}
	// a string slice to hold the keys
	// k := make([]string, len(files))
	var k []string
	for _, file := range files {
		k = append(k, file.Name()[:10])
		// fmt.Println(file.Name()[:10])
		// fmt.Println(file.Name()[:10])
	}
	// fmt.Println(k)

	for i := 0; i < len(k); i++ {
		var path string = "../../node_jsons/"+k[i] + ".json"
		byteValue, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("something wrong!")
			// return err
		}
		
		var result map[string]json.RawMessage
		err = json.Unmarshal(byteValue, &result)
		if err != nil {
			fmt.Println("something wrong while parsing json!")
			// return err
		}
		states := make(map[string]string)
		for s, _ := range result {
			states[s] = "on"
		}
		//first time just store the state to disk
		if i == 0 {
			globalstate = states
			outputjson(k[i], states)
		} else {
			RecordState(globalstate, states)
			outputjson(k[i], globalstate)
		}
		// k := make(map[string]string)
		// for s, _ := range result {
		// 	k[s] = "on"
		// }

	}

}

func outputjson(filename string, result map[string]string) {
	// Convert golang object back to byte
	byteValue, err := json.Marshal(result)
	if err != nil {
		fmt.Println("something wrong!")
		// return err
	}

	// Write back to file
	err = ioutil.WriteFile("../../nodes_states/"+filename+"_states.json", byteValue, 0644)
	// fmt.Println(file.Name()[:10])
}

func mytest2() {

	m1 := make(map[string]string)
	m1["1.172.126.13:8333"] = "on"
	m1["1.32.249.37:8333"] = "on"
	m1["100.24.150.236:8333"] = "on"

	globalstate = m1

	m2 := make(map[string]string)
	m2["1.32.249.37:8333"] = "on"
	m2["100.24.150.236:8333"] = "on"
	RecordState_test(globalstate, m2)
	fmt.Println(globalstate)

	m3 := make(map[string]string)
	m3["1.32.249.37:8333"] = "on"
	m3["100.24.150.236:8333"] = "on"
	m3["Iamfuckingonian:8333"] = "on"
	RecordState_test(globalstate, m3)
	fmt.Println(globalstate)

	m4 := make(map[string]string)
	m4["1.172.126.13:8333"] = "on"
	m4["1.32.249.37:8333"] = "on"
	m4["100.24.150.236:8333"] = "on"
	m4["Iamfuckingonian:8333"] = "on"
	m4["Iamfuckingonian2:8333"] = "on"
	RecordState_test(globalstate, m4)
	fmt.Println(globalstate)

	m5 := make(map[string]string)
	m5["Iamfuckingonian2:8333"] = "on"
	RecordState_test(globalstate, m5)
	fmt.Println(globalstate)

	m6 := make(map[string]string)
	m6["1.172.126.13:8333"] = "on"
	m6["1.32.249.37:8333"] = "on"
	m6["100.24.150.236:8333"] = "on"
	m6["Iamfuckingonian:8333"] = "on"
	m6["Iamfuckingonian2:8333"] = "on"
	RecordState_test(globalstate, m6)
	fmt.Println(globalstate)
}

func mytest() {
	m1 := make(map[string]string)
	m1["a"] = "on"
	m1["b"] = "on"
	m1["c"] = "on"
	m1["d"] = "on"

	globalstate = m1

	m2 := make(map[string]string)
	m2["a"] = "on"
	m2["b"] = "on"
	m2["c"] = "on"

	RecordState_test(globalstate, m2)
	// fmt.Println(globalstate)
	fmt.Println()

	m3 := make(map[string]string)
	m3["a"] = "on"
	m3["b"] = "on"
	m3["e"] = "on"
	RecordState_test(globalstate, m3)
	fmt.Println()

	m4 := make(map[string]string)
	m4["a"] = "on"
	m4["b"] = "on"
	m4["e"] = "on"
	m4["f"] = "on"
	RecordState_test(globalstate, m4)

	m5 := make(map[string]string)
	m5["a"] = "on"
	m5["b"] = "on"
	m5["c"] = "on"
	m5["e"] = "on"
	m5["f"] = "on"
	RecordState_test(globalstate, m5)

	m6 := make(map[string]string)
	m6["b"] = "on"
	m6["c"] = "on"
	RecordState_test(globalstate, m6)

	m7 := make(map[string]string)
	m7["a"] = "on"
	m7["b"] = "on"
	m7["c"] = "on"
	RecordState_test(globalstate, m7)

	fmt.Println("g: ", globalstate)

	// m := make(map[string]string)
	// m["1.172.126.13:8333"] = "a"
	// m["1.32.249.37:8333"] = "c"
	// m["100.24.150.236:8333"] = "b"
	// fmt.Println(m)
}
func RecordState(resultjson map[string]string, m2 map[string]string) {
	//copy new map
	m1 := make(map[string]string)
	for k, v := range resultjson {
		m1[k] = v
	}
	// fmt.Println(m1)
	// remove the same nodes and record to the reultjson as "on"
	for k2, _ := range m2 {
		for k1, _ := range m1 {
			result := k1 == k2
			if result {
				if m1[k1] == "off" || m1[k1] == "left" {
					delete(m1, k1)
					delete(m2, k2)
					resultjson[k1] = "join"
				} else {
					delete(m1, k1)
					delete(m2, k2)
					resultjson[k1] = "on"
				}

			}

		}
	}

	// fmt.Println("to on: ", resultjson)
	// the rest of nodes in k2 are the new nodes, recorded as "join"
	for k2, _ := range m2 {
		resultjson[k2] = "join"
	}
	// fmt.Println("to join: ", resultjson)

	//the nodes who has already off
	for k, v := range resultjson {
		if v == "left" {
			resultjson[k] = "off"
		}
	}
	// fmt.Println("to off: ", resultjson)

	// the rest of nodes in k1 are the left nodes, recorded as "left"
	// fmt.Println(m1)
	for k1, _ := range m1 {
		if m1[k1] != "left" && m1[k1] != "off" {
			resultjson[k1] = "left"
		}
	}
	// fmt.Println("to left: ", resultjson)
	fmt.Println()
	globalstate = resultjson

}
func RecordState_test(resultjson map[string]string, m2 map[string]string) {
	//copy new map
	m1 := make(map[string]string)
	for k, v := range resultjson {
		m1[k] = v
	}
	// fmt.Println(m1)
	// remove the same nodes and record to the reultjson as "on"
	for k2, _ := range m2 {
		for k1, _ := range m1 {
			result := k1 == k2
			if result {
				if m1[k1] == "off" || m1[k1] == "left" {
					delete(m1, k1)
					delete(m2, k2)
					resultjson[k1] = "join"
				} else {
					delete(m1, k1)
					delete(m2, k2)
					resultjson[k1] = "on"
				}

			}

		}
	}
	// fmt.Println("m1", m1)
	// fmt.Println("m2", m2)

	// fmt.Println("to on: ", resultjson)
	// the rest of nodes in k2 are the new nodes, recorded as "join"
	for k2, _ := range m2 {
		resultjson[k2] = "join"
	}
	// fmt.Println("to join: ", resultjson)

	//the nodes who has already off
	for k, v := range resultjson {
		if v == "left" {
			resultjson[k] = "off"
		}
	}
	// fmt.Println("to off: ", resultjson)

	// the rest of nodes in k1 are the left nodes, recorded as "left"
	// fmt.Println(m1)
	for k1, _ := range m1 {
		if m1[k1] != "left" && m1[k1] != "off" {
			resultjson[k1] = "left"
		}
	}

	// fmt.Println("to left: ", resultjson)
	fmt.Println()
	globalstate = resultjson

	// fmt.Println("m1: ", m1)
	// fmt.Println("m2: ", m2)
}
