package io

import (
	"bufio"
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

	in := bufio.NewReader(os.Stdin)
	var values [5]string
	values[0] = time.Now().Format("02/01/2006")
	values[1] = time.Now().Format("15")
	minutes := time.Now().Format("4")
	i, err := strconv.Atoi(minutes)
	if err != nil {
		log.Fatal(err)
	}
	if i >= 30 {
		t, err := strconv.Atoi(values[1])
		if err != nil {
			log.Fatal(err)
		}
		t++
		values[1] = strconv.Itoa(t)
	}
	fmt.Println("Quantita': ")
	values[1], err = in.ReadString('\n')
	fmt.Println("Qualita': ")
	values[2], err = in.ReadString('\n')
	fmt.Println("Utilizzo: ")
	values[3], err = in.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	for i := 1; i < len(values); i++ {
		values[i] = values[i][0 : len(values[i])-1]
	}
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
	fmt.Println("How much material do you have to smoke? (1g = 1.0): ")
	fmt.Scanf("%f", &material)
	f, err := os.OpenFile("buji.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString("giorno,quantita,qualita,tipo,ore," + fmt.Sprintf("%.2f", material))
	fmt.Println("You have also to enter a buji to start. Press Enter.")
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
