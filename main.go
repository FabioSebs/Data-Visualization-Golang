package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int               { return len(p) }
func (p PairList) Swap(i, j int)          { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool     { return p[i].Value < p[j].Value }
func (p PairList) ReturnKey(i int) string { return p[i].Key }
func (p PairList) ReturnVal(i int) int    { return p[i].Value }

func generateBarItems(data PairList) []opts.BarData {
	barData := []int{}
	items := make([]opts.BarData, 0)

	for i := 0; i <= 6; i++ {
		barData = append(barData, data[i].Value)
		fmt.Println(barData)
	}

	for _, v := range barData {
		items = append(items, opts.BarData{Value: v})
	}
	return items
}

func main() {
	//Variables
	f, err := os.Open("brooklyn.csv")
	data := []string{}
	freq := map[string]int{}
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Common Foods Wasted in Brooklyn",
		Subtitle: "Edible and wasted food found in retailer trash piles around Brooklyn",
	}))

	i := 0

	// Reading the CSV and Extracting the Data
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(f)
	reader.LazyQuotes = true
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, row[5])
		// fmt.Println(reflect.TypeOf(row[5]))
	}
	// fmt.Println(data)

	//Count of Elements in the Data
	for _, item := range data {
		_, exist := freq[item]

		if exist {
			freq[item] += 1
		} else {
			freq[item] = 1
		}
	}

	p := make(PairList, len(freq))
	// for k, v := range freq {
	// 	fmt.Printf("Item : %s \nCount: %d\n", k, v)
	// }

	//Sorting Map in Descending Order
	for k, v := range freq {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))

	for _, k := range p {
		fmt.Printf("%v\t%v\n", k.Key, k.Value)
	}

	//Setting Instance of Bar
	bar.SetXAxis([]string{
		p[0].Key[0:4],
		p[1].Key[0:4],
		p[2].Key[0:4],
		p[3].Key[0:4],
		p[4].Key[0:4],
		p[5].Key[0:4],
		p[6].Key[0:4],
	}).AddSeries("Values", generateBarItems(p))

	e, _ := os.Create("brooklyn-data.html")
	bar.Render(e)
}
