package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 10; i++ {
	// 	x := rand.Intn(100)
	// 	fmt.Println(x)
	// }

	//replication factor, storage limit per node (10G -> 80blks, 15G->120blks) blk02116
	// number of nodes, churn rate (just calculate difference)

	// var nodes []Data
	// nodes = readCSV()

	// churn_n := readChurnCSV()

	// for i := len(nodes) - 1; i > 0; i-- {
	// 	if i == len(nodes)-1 {
	// 		// initiate first environment

	// 	}

	// 	// fmt.Println(nodes[i])
	// }

}

type LiveNode struct {
	Name     string
	numofblk int
	blk      []string
	alive    bool
	isfull   bool
}
type ChurnData struct {
	timestamp int64
	churn_r   float64
	add_n     int
}
type Data struct {
	timestamp_d   int64
	total_nodes_d int
}

func readChurnCSV() (churn_n []ChurnData) {
	csvfile, err := os.Open("nodes.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)

	var churn_nodes []ChurnData

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1], record[2])
		var data ChurnData
		data.timestamp, err = strconv.ParseInt(record[0], 10, 64)
		data.churn_r, err = strconv.ParseFloat(record[1], 64)
		data.add_n, err = strconv.Atoi(record[1])
		churn_nodes = append(churn_nodes, data)
	}
	return churn_nodes
}
func readCSV() (nodes []Data) {
	// Open the file
	csvfile, err := os.Open("nodes.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	var data []Data

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
		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1])
		var content Data
		content.timestamp_d, err = strconv.ParseInt(record[0], 10, 64)
		content.total_nodes_d, err = strconv.Atoi(record[1])
		data = append(data, content)
	}
	return data

}
func appendToCSV(data Data) {
	// check if nodes.csv exists
	_, err := os.Open("ipfs_nodes.csv")
	if err != nil {
		// fmt.Println(os.IsNotExist(err)) //true  證明檔案已經存在
		// fmt.Println(err)                //open widuu.go: no such file or directory
		os.Create("ipfs_nodes.csv")
	}

	var path = "ipfs_nodes.csv"
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

// GenerateRandnum 生成最大範圍內隨機數
func GenerateRandnum() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(100)
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}

// GenerateRangeNum 生成一個區間範圍的隨機數
func GenerateRangeNum(min, max int) int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max - min)
	randNum = randNum + min
	fmt.Printf("rand is %v\n", randNum)
	return randNum
}
