/*
STANDARD:
DATA,QUANTITA',QUALITA',METODO,ORE

TODO:
Calculate over how many days i smoked the last x weight
Calculate how much i smoked on the last x days
CCC
*/

package main

import (
	"log"
	"strconv"

	i "github.com/punkycommunist/bujigo/io"
	m "github.com/punkycommunist/bujigo/menu"
)

//CsvLine is a struct for one horizontal line in csv
type CsvLine struct {
	date     string
	quantity string
	quality  string
	method   string
	hour     string
	remains  string
}

func main() {
	fileName := i.SearchCsvInCurrentDirectory()
	lines, err := i.ReadCsv(fileName)
	if err != nil {
		panic(err)
	}
	var iDate []string
	var iQuantity []float64
	var iQuality []string
	var iMethod []string
	var iHour []int32
	var iRemains float64
	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			date:     line[0],
			quantity: line[1],
			quality:  line[2],
			method:   line[3],
			hour:     line[4],
			remains:  line[5],
		}
		iDate = append(iDate, data.date)
		if data.quantity != "quantita" { //skippa prima riga, metti 0 per tenere con gli altri index
			q, err := strconv.ParseFloat(data.quantity, 64)
			if err != nil {
				log.Println(err)
			}
			iQuantity = append(iQuantity, q)
		} else {
			iQuantity = append(iQuantity, 0.0)
		}

		iQuality = append(iQuality, data.quality)
		iMethod = append(iMethod, data.method)
		if data.hour != "ore" { //skippa prima riga, metti 0 per tenere con gli altri index
			h, err := strconv.ParseInt(data.hour, 10, 32)
			if err != nil {
				log.Println(err)
			}
			iHour = append(iHour, int32(h))
		} else {
			iHour = append(iHour, 0)
		}
		if data.remains != "" {
			r, err := strconv.ParseFloat(data.remains, 64)
			if err != nil {
				log.Println(err)
			}
			iRemains = r
		}
	}
	weightSmoked := 0.0
	for i := 1; i < len(iQuantity); i++ {
		weightSmoked += iQuantity[i]
	}
	iRemains -= weightSmoked
	m.PrintMenu(iDate, iQuantity, iQuality, iMethod, iHour, iRemains)
}
