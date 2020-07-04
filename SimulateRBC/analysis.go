package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func RunAnalysis() {
	ReadRequest("../../nodes/results/10G/Baseline_norepair_nodes_withBlk_state/")

}

func ReadRequest(path string) {
	names := GetFilesName(path)

	// Open the file
	csvfile, err := os.Open("/home/hank/go/read_old.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	var tmp_timestamp time.Time
	i, j := 0, 0
	// Iterate through the records
	var requests []string
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		layout := "2006-01-02 15:04:05"
		current_timestamp, err := time.Parse(layout, record[1])
		if i == 0 {
			tmp_timestamp = current_timestamp

		} else {
			diff := current_timestamp.Sub(tmp_timestamp)
			// fmt.Println(diff)
			if diff.Minutes() >= 5 {
				fmt.Println(diff)
				fmt.Println(current_timestamp)
				counter := calculateAccessRate(requests, path+names[j]+"_states.json")
				fmt.Println(counter)
				j++

			} else {
				requests = append(requests, record[3])
			}

		}

		
		i++

	}

}

func calculateAccessRate(requests []string, file string) map[string]int {
	counter := make(map[string]int)
	state := make(map[string][]string)
	byteValue, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(byteValue, &state)
	if err != nil {
		fmt.Println("something wrong while parsing json!")
		// return err
	}
	fmt.Println(state)
	for k, v := range state {
		if v == nil || len(v) == 0 {
			delete(state, k)
		} else {
			for i := range requests {
				for j := range v {
					if "blk"+requests[i] == v[j] {
						counter["success"] += 1
					} else if j == len(v)-1 {
						counter["failure"] += 1
					}
				}
			}
		}
	}

	return counter

}
