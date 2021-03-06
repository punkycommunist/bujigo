//go:generate goversioninfo -icon=ganja.ico -manifest=bujigo.exe.manifest
/*
STANDARD:
DATA,QUANTITA',QUALITA',METODO,ORE

TODO:
Calculate over how many days i smoked the last x weight
Calculate how much i smoked on the last x day
Update tools
Work on quality
*/

package main

import (
	"log"
	"strconv"

	i "github.com/punkycommunist/bujigo/io"
	m "github.com/punkycommunist/bujigo/menu"
)

func main() {
	i.CheckForUpdates()
	jsp := i.ReadJSONPreferences()
	fileName := i.SearchCsvInCurrentDirectory(jsp)
	lines, err := i.ReadCsv(fileName)
	if err != nil {
		panic(err)
	}
	var c i.CsvFile
	for _, line := range lines {
		data := i.CsvLine{
			Date:     line[0],
			Quantity: line[1],
			Quality:  line[2],
			Method:   line[3],
			Hour:     line[4],
			Remains:  line[5],
		}
		c.Date = append(c.Date, data.Date)
		if data.Quantity != "quantita" { //skippa prima riga
			q, err := strconv.ParseFloat(data.Quantity, 64)
			if err != nil {
				log.Println(err)
			}
			c.Quantity = append(c.Quantity, q)
		} else {
			c.Quantity = append(c.Quantity, 0.0)
		}

		c.Quality = append(c.Quality, data.Quality)
		c.Method = append(c.Method, data.Method)
		if data.Hour != "ore" { //skippa prima riga
			h, err := strconv.ParseInt(data.Hour, 10, 32)
			if err != nil {
				log.Println(err)
			}
			c.Hour = append(c.Hour, int(h))
		} else {
			c.Hour = append(c.Hour, 0)
		}
		if data.Remains != "" {
			r, err := strconv.ParseFloat(data.Remains, 64)
			if err != nil {
				log.Println(err)
			}
			c.Remains = r
		}
	}
	weightSmoked := 0.0
	for i := 1; i < len(c.Quantity); i++ {
		weightSmoked += c.Quantity[i]
	}
	c.Remains -= weightSmoked
	m.PrintMenu(c, i.Version)
}
