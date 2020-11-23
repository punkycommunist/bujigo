package io

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

//StartBujiSequence initializes the process to add another buji in the database
func StartBujiSequence() {
	filename := SearchCsvInCurrentDirectory()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var values [5]string
	var isToday string
	prompt("Il buji e' stato fumato oggi? [y] per si, [n] per no e inserire la data.")
	fmt.Scan(&isToday)
	if err != nil {
		log.Fatal(err)
	}
	if isToday == "y" { //if [enter] gets pressed
		values[0] = time.Now().Format("02/01/2006")
	} else if isToday == "n" {
		prompt("Che giorno era?")
		fmt.Scan(&values[0])
	} else { //if invalid character
		log.Fatal(isToday + " invalid.")
	}

	var thisHour string
	prompt("Il buji e' stato fumato a quest'ora? [y] per si, [n] per no e inserire la data.")
	fmt.Scan(&thisHour)
	if err != nil {
		log.Fatal(err)
	}
	if thisHour == "y" { //if [y] gets pressed
		values[4] = time.Now().Format("15")
		minutes := time.Now().Format("4")
		//from string to int to apply logic
		i, err := strconv.Atoi(minutes)
		if err != nil {
			log.Fatal(err)
		}
		if i >= 30 {
			t, err := strconv.Atoi(values[4])
			if err != nil {
				log.Fatal(err)
			}
			t++
			values[4] = strconv.Itoa(t)
		}
	} else if thisHour == "n" {
		prompt("Che ore erano?")
		fmt.Scan(&values[4])
	} else { //if invalid character
		log.Fatal(thisHour + " invalid.")
	}
	prompt("Quantita': ")
	fmt.Scan(&values[1])
	prompt("Qualita': ")
	fmt.Scan(&values[2])
	prompt("Utilizzo: ")
	fmt.Scan(&values[3])

	s := "\n" + values[0] + "," + values[1] + "," + values[2] + "," + values[3] + "," + values[4] + ","
	f.WriteString(s)
}

//SearchCsvInCurrentDirectory searches .csv files by getting an array of the elements present in that directory
func SearchCsvInCurrentDirectory() string {
	out, err := exec.Command("pwd").Output()
	if err != nil {
		log.Fatal(err)
	}
	out = out[:len(out)-1]
	files, err := filepath.Glob(string(out) + "/*")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(files); i++ {
		if files[i][len(files[i])-4:] == ".csv" {
			return files[i]
		}
	}
	log.Println("No .csv file found. Do you want to create the default [buji.csv]?")
	material := 0.0
	prompt("How much material do you have to smoke? (1g = 1.0): ")
	fmt.Scanf("%f", &material)
	f, err := os.OpenFile("buji.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("giorno,quantita,qualita,tipo,ore," + fmt.Sprintf("%.2f", material))
	prompt("You have also to enter a buji to start. Press Enter.")
	fmt.Scanln()
	StartBujiSequence()
	return "buji.csv"
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
func prompt(s string) {
	fmt.Println(s)
	fmt.Printf("$ ")
}
