package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// Reading the CSV and Extracting the Data
	f, err := os.Open("brooklyn.csv")
	data := []string{}
	freq := make(map[string]int)
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

	//Counts of Elements in the Data
	for _, item := range data {
		_, exist := freq[item]

		if exist {
			freq[item] += 1
		} else {
			freq[item] = 1
		}
	}

	for k, v := range freq {
		fmt.Printf("Item : %s \nCount: %d\n", k, v)
	}
}
