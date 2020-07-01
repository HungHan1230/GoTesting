package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"go4.org/sort"
	"gonum.org/v1/plot/plotter"
)

type mydata struct {
	timestamp int64
	nodes     int
	churn_r   float64
	churn_n   int
	add_n     int
	add_r     float64
}

func caluculateCount() {
	data := GetChurn()
	m := make(map[int]int)
	for i := 0; i < len(data); i++ {
		if data[i].churn_n == 0 {
			continue
		}
		if m[data[i].churn_n] != 0 {
			m[data[i].churn_n] = m[data[i].churn_n] + 1
		} else {
			m[data[i].churn_n] = 1
		}

	}
	fmt.Println(m)

	var total int = 0
	m2 := make(map[int]int)
	var keys []int

	for k, v := range m {
		keys = append(keys, k)
		total += v
	}
	sort.Ints(keys)
	fmt.Println(keys)
	// var tmp int = 0
	var points plotter.XYs
	for i := 0; i < len(keys); i++ {
		// fmt.Println(keys[i])
		if keys[i] > 300{
			continue
		}
		m2[keys[i]] = total

		points = append(points, struct{ X, Y float64 }{float64(keys[i]), float64(total)})
		total -= m[keys[i]]
	}
	fmt.Println(m2)

	plottest(points)



}

func calculateDailyChurnCumulative() {
	data := GetChurn()
	day_map := make(map[string][]float64)

	var previous, unixTimeUTC time.Time
	layout := "2006-01-02"

	for i := 0; i < len(data); i++ {
		if data[i].churn_n == 0 {
			// fmt.Println("churn nodes = 0")
			continue
		}
		now := data[i].timestamp
		unixTimeUTC = time.Unix(now, 0)
		judgement := unixTimeUTC.Sub(previous)
		key := unixTimeUTC.Format(layout)
		// fmt.Println(unixTimeUTC.Date())
		// fmt.Println(previous.Date())

		if i == 0 {
			previous = unixTimeUTC
			day_map[key] = []float64{float64(data[i].churn_n)}
		} else if judgement.Hours()/24 >= 1 {
			// } else if unixTimeUTC.Date() != previous.Date() {
			previous = unixTimeUTC
			// key := unixTimeUTC.Format(layout)
			if day_map[key] == nil {
				day_map[key] = []float64{float64(data[i].churn_n)}
			} else {
				day_map[key] = append(day_map[key], float64(data[i].churn_n))
			}

		} else {
			p_key := previous.Format(layout)
			day_map[p_key] = append(day_map[p_key], float64(data[i].churn_n))
		}

	}
	fmt.Println(day_map)

}

func calculateChurnCumulative() {

	data := GetChurn()
	// fmt.Println(data)
	percentage_m := make(map[string]int)
	var judge float64 = 0.0

	for i := 0; i < 100; i++ {
		for j := 0; j < len(data); j++ {
			if data[j].churn_r > judge {
				key := strconv.FormatFloat(judge, 'g', 5, 64)
				percentage_m[key] = percentage_m[key] + 1
			}
		}
		// fmt.Println(judge)
		judge += 0.01
	}
	fmt.Println(percentage_m)
}

func calculateChurn() {
	// Open the file
	csvfile, err := os.Open("nodes_snapshots_reverse_forchurn.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	// mycode
	var counter int
	counter = 1
	var tmptimestamp, tmpnodes, previoustimestamp, previousnodes, churn_rate, churn_nodes, add_nodes, add_rate float64
	var dataArr []mydata
	var points_churn, points_churn_nodes plotter.XYs
	var points_add plotter.XYs
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
		if counter == 1 {
			// tmptimestamp, err = strconv.ParseInt(record[0], 10, 64)
			tmptimestamp, err = strconv.ParseFloat(record[0], 64)
			tmpnodes, err = strconv.ParseFloat(record[1], 64)
			//initialization
			previoustimestamp = tmptimestamp
			previousnodes = tmpnodes
			churn_rate = 0.0
			churn_nodes = 0
			add_nodes = 0
			add_rate = 0.0
			// first record
			var tmpdata mydata
			tmpdata.timestamp = int64(previoustimestamp)
			tmpdata.churn_r = churn_rate
			tmpdata.churn_n = 0
			tmpdata.add_n = int(add_nodes)
			tmpdata.nodes = int(previousnodes)
			tmpdata.add_r = add_rate
			dataArr = append(dataArr, tmpdata)
		} else {
			tmptimestamp, err = strconv.ParseFloat(record[0], 64)
			tmpnodes, err = strconv.ParseFloat(record[1], 64)

			var judgement float64
			judgement = previousnodes - tmpnodes

			//If preivousnodes greater than tmpnodes, it means that there were some nodes disconnected (churn).
			if judgement > 0 {
				churn_rate = judgement / previousnodes
				churn_nodes = judgement
				add_nodes = 0
				add_rate = 0
			} else {
				churn_rate = 0
				churn_nodes = 0
				add_nodes = tmpnodes - previousnodes
				add_rate = add_nodes / previousnodes
			}

			previoustimestamp = tmptimestamp
			previousnodes = tmpnodes

			var tmpdata mydata
			tmpdata.timestamp = int64(previoustimestamp)
			tmpdata.churn_r = churn_rate
			tmpdata.churn_n = int(churn_nodes)
			tmpdata.add_n = int(add_nodes)
			tmpdata.nodes = int(previousnodes)
			tmpdata.add_r = add_rate
			dataArr = append(dataArr, tmpdata)
			points_churn = append(points_churn, struct{ X, Y float64 }{previoustimestamp, churn_rate})
			points_churn_nodes = append(points_churn_nodes, struct{ X, Y float64 }{previoustimestamp, churn_nodes})
			points_add = append(points_add, struct{ X, Y float64 }{previoustimestamp, add_rate})

		}
		counter++
	}
	// fmt.Println(dataArr)
	writeChurnToCSV(dataArr)

	plotchurn(points_churn)
	plotchurn(points_churn_nodes)
	plotadd_r(points_add)
	// plotchurn_V2(points)

}

func writeChurnToCSV(data []mydata) {
	// check if nodes.csv exists
	_, err := os.Open("nodes_churn.csv")
	if err != nil {
		// fmt.Println(os.IsNotExist(err)) //true  證明檔案已經存在
		// fmt.Println(err)                //open widuu.go: no such file or directory
		os.Create("nodes_churn.csv")
	}

	var path = "nodes_churn.csv"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	var tocsv [][]string
	for i := 0; i < len(data); i++ {
		originalTimestamp := strconv.FormatInt(data[i].timestamp, 10)
		// originalNodes := string(data[i].nodes) // so interesting
		originalNodes := strconv.Itoa(data[i].nodes)
		var tmptime int64
		tmptime = data[i].timestamp
		unixTimeUTC := time.Unix(tmptime, 0) //gives unix time stamp in utc
		// tocsv = append(tocsv, []string{strconv.FormatInt(data[i].timestamp, 10), fmt.Sprintf("%f", data[i].churn_r), strconv.Itoa(data[i].add_n)})
		tocsv = append(tocsv, []string{originalTimestamp, originalNodes, unixTimeUTC.Format("2006-01-02 15:04:05"), fmt.Sprintf("%f", data[i].churn_r), strconv.Itoa(data[i].churn_n), strconv.Itoa(data[i].add_n), fmt.Sprintf("%f", data[i].add_r)})
	}

	w.WriteAll(tocsv)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
