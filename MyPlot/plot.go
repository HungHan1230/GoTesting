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
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

func main() {
	// simple()
	// plotsnapshots()
	// RunExample()
	unixTime1 := time.Unix(1592295927, 0) //gives unix time stamp in utc
	unixTime2 := time.Unix(1587121480, 0) //gives unix time stamp in utc
	fmt.Println(unixTime1.Format("2006-01-02 15:04:05"))
	fmt.Println(unixTime2.Format("2006-01-02 15:04:05"))

	//test percernt package
	fmt.Println(percent.PercentFloat(0.3, 200))	
	fmt.Println(percent.PercentOf(5, 200))

	var testf float64
	testf = (10160.0-10136.0)/10160.0
	fmt.Printf("%f %% \n",testf*100)
	testf = 625.0 / 10000.0
	fmt.Printf("%f %% \n",testf*100)
	

}
func plotchurn() {
	p, _ := plot.New()

	p.Title.Text = "The churn rate of bitcoin nodes from 2020/4/17 to 2020/6/16"
	p.X.Label.Text = "timestamp"
	p.Y.Label.Text = "number of nodes"
	p.Add(plotter.NewGrid())
	// var points plotter.XYs
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
	var tmptimestamp, tmpnodes, currenttimestamp, currentnodes int64
	var churn_rate float64

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
			tmptimestamp, err = strconv.ParseInt(record[0], 10, 64)
			tmpnodes, err = strconv.ParseInt(record[1], 10, 64)
		} else {
			currenttimestamp, err = strconv.ParseInt(record[0], 10, 64)
			currentnodes, err = strconv.ParseInt(record[1], 10, 64)

			churn_rate := tmpnodes / currentnodes
			fmt.Printf("%f %f\n", churn_rate, tmptimestamp)
			tmptimestamp = currenttimestamp
			tmpnodes = currentnodes

		}
		fmt.Println(churn_rate)
		counter++

	}

}
func plotsnapshots() {
	p, _ := plot.New()

	p.Title.Text = "The snapshots of bitcoin nodes from 2020/4/17 to 2020/6/16"
	p.X.Label.Text = "timestamp"
	p.Y.Label.Text = "number of nodes"
	p.Add(plotter.NewGrid())

	// points := plotter.XYs{
	// 	{1592295927, 10462},
	// 	{1592295633, 10485},
	// 	{1592295338, 10480},
	// 	{1592294991, 10515},
	// 	{1592294614, 10497},
	// }

	var points plotter.XYs

	points = readcsv()
	// plotutil.AddLinePoints(p, points)

	// Make a scatter plotter and set its style.
	s, err := plotter.NewScatter(points)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}
	s.Shape = draw.PyramidGlyph{}
	p.Add(s)
	p.Legend.Add("scatter", s)

	p.Save(5*vg.Inch, 5*vg.Inch, "nodes.png")

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

	var count int
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
		if y < 7000 {
			fmt.Println("the pair is: ", record[0], record[1])
		}
		count++
	}
	fmt.Println("total: ", count)
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
