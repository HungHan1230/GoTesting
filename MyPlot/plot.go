package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"image/color"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/negah/percent"
	"github.com/rcrowley/go-metrics"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	
	// plotsnapshots()
	// plotsnapshots_V2()
	
	// calculateChurn()	
	// plotblk()

	// calculateAverageMaximumChurn()

	testplot()
}

func calculateAverageMaximumChurn() {
	csvfile, err := os.Open("./nodes_churn.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)

	// m := make(map[string]int)

	var count int
	var accumulate, max float64
	max = 0

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1]) //record[3]
		var tmp float64
		tmp, err = strconv.ParseFloat(record[1], 64)

		if max < tmp {
			max = tmp
		}
		accumulate += tmp
		count++
	}
	var ans = accumulate / float64(count)
	fmt.Printf("average churn rate: %f %% \n", ans*100) //average churn rate: 10.850174 %
	fmt.Printf("maximum churn rate: %f %% \n", max*100) //maximum churn rate: 3522.046200 %
}
func readblkcsv() (ps plotter.XYs, count int) {
	csvfile, err := os.Open("/home/hank/go/read.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)

	m := make(map[string]int)

	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("Question: %s Answer %s\n", record[0], record[1])//record[3]

		v, ok := m[record[3]]
		if !ok {
			// log.Fatal("It should be true")
			m[record[3]] = 1
		} else {
			m[record[3]] = v + 1
		}
		// if !(v == "Beego") {
		// 	log.Fatal("Wrong value")
		// }
	}
	// fmt.Println(m)
	// blk0000 too much, delete
	delete(m, "0")
	var points plotter.XYs
	for k, v := range m {

		// fmt.Println(fmt.Sprintf("%s: %s", k, v))
		kvalue, err := strconv.ParseFloat(k, 64)

		if err != nil {
			log.Fatal("something wrong")
		}
		points = append(points, struct{ X, Y float64 }{kvalue, float64(v)})
	}
	return points, len(m)

}

func plotblk() {
	p, _ := plot.New()

	p.Title.Text = "The access pattern of blk*.dat files"
	p.X.Label.Text = "blk*.dat"
	p.Y.Label.Text = "count"
	p.Add(plotter.NewGrid())

	var points plotter.XYs

	points, count := readblkcsv()
	// plotutil.AddLinePoints(p, points)

	// Make a scatter plotter and set its style.
	// s, err := plotter.NewScatter(points)
	s, err := plotter.NewHistogram(points, count)
	if err != nil {
		panic(err)
	}
	// s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	// s.Shape = draw.PyramidGlyph{}
	// s.GlyphStyle.Radius = vg.Points(5)
	p.Add(s)
	// p.Legend.Add("scatter", s)

	p.Save(8*vg.Inch, 4*vg.Inch, "blknodes.png")

}

func plotchurn(points plotter.XYs) {
	p, _ := plot.New()

	p.Title.Text = "The churn rate of bitcoin nodes from 2020/4/17 to 2020/6/16"
	p.X.Label.Text = "timestamp"
	p.Y.Label.Text = "churn rate"
	p.Add(plotter.NewGrid())
	// var points plotter.XYs
	s, po, err := plotter.NewLinePoints(points)
	if err != nil {
		panic(err)
	}
	s.Color = color.RGBA{R: 255, A: 255}
	// po.Shape = draw.PyramidGlyph{}
	po.Color = color.RGBA{R: 255, A: 255}
	// s.LineStyle.Width = vg.Points(1)
	// s.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	// s.LineStyle.Color = color.RGBA{B: 255, A: 255}
	// s.Shape = draw.PyramidGlyph{}
	p.Add(s, po)
	p.Legend.Add("linepoint", s)

	p.Save(15*vg.Inch, 5*vg.Inch, "nodes_churn.png")
}
func plotchurn_V2(points plotter.XYs) {
	p, _ := plot.New()

	p.Title.Text = "The churn rate of bitcoin nodes from 2020/4/17 to 2020/6/16"
	p.X.Label.Text = "timestamp"
	p.Y.Label.Text = "churn rate"
	p.Add(plotter.NewGrid())
	// var points plotter.XYs
	s, err := plotter.NewHistogram(points, points.Len())
	if err != nil {
		panic(err)
	}
	s.Color = color.RGBA{R: 255, A: 255}

	p.Add(s)
	// p.Legend.Add("linepoint", s)

	p.Save(8*vg.Inch, 5*vg.Inch, "nodes_churn.png")
}

type mydata struct {
	timestamp int64
	churn_r   float64
	add_n     int
}

func calculateChurn() {
	// Open the file
	csvfile, err := os.Open("nodes.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	// Parse the file
	r := csv.NewReader(csvfile)
	// mycode
	var counter int
	counter = 1
	var tmptimestamp, tmpnodes, previoustimestamp, previousnodes, churn_rate, add_nodes float64
	var dataArr []mydata
	var points plotter.XYs
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
			previoustimestamp = tmptimestamp
			previousnodes = tmpnodes
			// churn_rate = 0.0
		} else {
			previoustimestamp, err = strconv.ParseFloat(record[0], 64)
			previousnodes, err = strconv.ParseFloat(record[1], 64)

			var judgement float64
			judgement = previousnodes - tmpnodes
			if judgement > 0 {
				churn_rate = (judgement / previousnodes) * 100
				add_nodes = 0
			} else {
				churn_rate = 0
				add_nodes = tmpnodes - previousnodes
			}

			var tmpdata mydata
			tmpdata.timestamp = int64(tmptimestamp)
			tmpdata.churn_r = churn_rate
			tmpdata.add_n = int(add_nodes)
			dataArr = append(dataArr, tmpdata)
			points = append(points, struct{ X, Y float64 }{tmptimestamp, churn_rate})

			tmptimestamp = previoustimestamp
			tmpnodes = previousnodes
		}
		counter++
		// if churn_rate != 0 && churn_rate > 1 {
		// 	fmt.Println("pair: ", tmptimestamp, churn_rate)
		// }
	}
	// fmt.Println(dataArr)
	writeChurnToCSV(dataArr)

	// plotchurn(points)
	plotchurn_V2(points)

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
		var tmptime int64
		tmptime = data[i].timestamp
		unixTimeUTC := time.Unix(tmptime, 0) //gives unix time stamp in utc
		// tocsv = append(tocsv, []string{strconv.FormatInt(data[i].timestamp, 10), fmt.Sprintf("%f", data[i].churn_r), strconv.Itoa(data[i].add_n)})
		tocsv = append(tocsv, []string{unixTimeUTC.Format("2006-01-02 15:04:05"), fmt.Sprintf("%f", data[i].churn_r), strconv.Itoa(data[i].add_n)})
	}
	// fmt.Println(tocsv)

	// tocsv += strconv.FormatInt(data.timestamp_d, 10) + "," + strconv.Itoa(data.total_nodes_d)
	// fmt.Println(tocsv)
	w.WriteAll(tocsv)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
func plotsnapshots() {
	p, _ := plot.New()

	p.Title.Text = "The snapshots of bitcoin nodes from 2020/4/17 to 2020/6/16"
	p.X.Label.Text = "timestamp"
	p.Y.Label.Text = "number of nodes"
	p.Add(plotter.NewGrid())
	var points plotter.XYs

	points = readcsv()
	// plotutil.AddLinePoints(p, points)

	// Make a scatter plotter and set its style.
	s, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.Shape = draw.BoxGlyph{}
	// s.Shape = draw.CrossGlyph{}
	// s.Shape = draw.CircleGlyph{}
	p.Add(s)
	p.Legend.Add("scatter", s)

	p.Save(15*vg.Inch, 6*vg.Inch, "nodes.png")

}
func plotsnapshots_V2() {
	p, _ := plot.New()

	p.Title.Text = "The snapshots of bitcoin nodes from 2020/4/17 to 2020/6/16"
	p.X.Label.Text = "timestamp"
	p.Y.Label.Text = "number of nodes"
	p.Add(plotter.NewGrid())

	var points plotter.XYs

	points = readcsv()
	// plotutil.AddLinePoints(p, points)

	// Make a scatter plotter and set its style.
	s, err := plotter.NewHistogram(points, points.Len())
	if err != nil {
		panic(err)
	}
	// s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	// s.Shape = draw.PyramidGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)
	// p.Legend.Add("scatter", s)
	p.Save(8*vg.Inch, 6*vg.Inch, "nodes.png")

}

func readcsv() plotter.XYs {
	// Open the file
	csvfile, err := os.Open("nodes.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r := csv.NewReader(bufio.NewReader(csvfile))

	//my code
	var points plotter.XYs

	var count, min int
	min = 100000
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
		var x, y float64
		x, err = strconv.ParseFloat(record[0], 64)
		y, err = strconv.ParseFloat(record[1], 64)
		points = append(points, struct{ X, Y float64 }{x, y})
		// if y < 7000 {
		// 	fmt.Println("the pair is: ", record[0], record[1])
		// }
		count++
		if min > int(y) {
			min = int(y)
		}

	}
	fmt.Println("min: ", min)     //6053
	fmt.Println("total: ", count) //15250
	return points
}

// type xy struct{ x, y float64 }

func RunExample() {
	xys, err := readData("data.txt")
	if err != nil {
		log.Fatalf("could not read data.txt: %v", err)
	}
	_ = xys

	err = plotData("out.png", xys)
	if err != nil {
		log.Fatalf("could not plot data: %v", err)
	}
}

func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var xys plotter.XYs
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
			continue
		}
		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
}

func plotData(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}

	// create scatter with all data points
	s, err := plotter.NewScatter(xys)
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	s.GlyphStyle.Shape = draw.CrossGlyph{}
	s.Color = color.RGBA{R: 255, A: 255}
	p.Add(s)

	var x, c float64
	x = 1.2
	c = -3

	// create fake linear regression result
	l, err := plotter.NewLine(plotter.XYs{
		{3, 3*x + c}, {20, 20*x + c},
	})
	if err != nil {
		return fmt.Errorf("could not create line: %v", err)
	}
	p.Add(l)

	wt, err := p.WriterTo(256, 256, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}

	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	return nil
}
func simple() {
	p, _ := plot.New()

	p.Title.Text = "Blk.dat access frequency"
	p.X.Label.Text = "Blk.dat"
	p.Y.Label.Text = "Count"

	// var coordination plotter.XY
	// coordination.X = 1.10
	// coordination.Y = 18000

	points := plotter.XYs{
		{2.0, 60000.0},
		{4.0, 40000.0},
		{6.0, 30000.0},
		{8.0, 25000.0},
		{10.0, 23000.0},
	}

	plotutil.AddLinePoints(p, points)

	p.Save(4*vg.Inch, 4*vg.Inch, "price.png")
}
func mytest() {
	unixTime1 := time.Unix(1592295927, 0) //gives unix time stamp in utc
	unixTime2 := time.Unix(1587121480, 0) //gives unix time stamp in utc
	fmt.Println(unixTime1.Format("2006-01-02 15:04:05"))
	fmt.Println(unixTime2.Format("2006-01-02 15:04:05"))

	//test percernt package
	fmt.Println(percent.PercentFloat(0.3, 200))
	fmt.Println(percent.PercentOf(5, 200))

	// 1 10136.0
	// 2 10160.0
	// 3 9901.0
	var a, b, testf float64
	a = 10136.0
	b = 10160.0
	testf = (b - a) / a
	fmt.Printf("%f %% \n", testf*100)
}

func testplot() {
	s := metrics.NewExpDecaySample(1024, 0.015) // or metrics.NewUniformSample(1028)

	h := metrics.NewHistogram(s)

	metrics.Register("baz", h)
	h.Update(1)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		j++
		h.Update(j)
	}
}
