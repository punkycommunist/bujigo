package io

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const defaultSettings string = `{
	"QDayAverage": {
	  "worst": 0.55,
	  "best": 0.45
	},
	"QRemains": {
	  "worst": 1.0,
	  "best": 5.0
	},
	"QRemainingDays": {
	  "worst": 3.0,
	  "best": 7.0
	}
  }`

//CsvLine is a struct for one horizontal line in csv
type CsvLine struct {
	Date     string
	Quantity string
	Quality  string
	Method   string
	Hour     string
	Remains  string
}

//CsvFile represents the loaded .csv file
type CsvFile struct {
	Date     []string
	Quantity []float64
	Quality  []string
	Method   []string
	Hour     []int
	Remains  float64
}

//JSONPreferences is the struct representing the settings.json file
type JSONPreferences struct {
	QDayAverage struct {
		Worst float64 `json:"worst"`
		Best  float64 `json:"best"`
	} `json:"QDayAverage"`
	QRemains struct {
		Worst float64 `json:"worst"`
		Best  float64 `json:"best"`
	} `json:"QRemains"`
	QRemaininingDays struct {
		Worst float64 `json:"worst"`
		Best  float64 `json:"best"`
	} `json:"QRemainingDays"`
	QAvgBujiSmokedADay struct {
		Worst float64 `json:"worst"`
		Best  float64 `json:"best"`
	} `json:"QAvgBujiSmokedADay"`
}

//ReadJSONPreferences reads from the settings.json file in the directory of the program
func ReadJSONPreferences() JSONPreferences {
	var JSONPreferences JSONPreferences
	files, err := filepath.Glob("*")
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(files); i++ {
		if files[i] == "settings.json" {
			jsonFile, err := os.Open("settings.json")
			if err != nil {
				log.Fatal(err)
			}
			// read our opened jsonFile as a byte array.
			byteValue, _ := ioutil.ReadAll(jsonFile)
			// we unmarshal our byteArray which contains our
			// jsonFile's content into 'users' which we defined above
			json.Unmarshal(byteValue, &JSONPreferences)
			jsonFile.Close()
			return JSONPreferences
		}
	}
	prompt("settings.json non trovato. creare il default ora? [Enter] per si")
	fmt.Scanln()
	f, err := os.OpenFile("settings.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	f.WriteString(defaultSettings)
	return ReadJSONPreferences()
}

//WriteJSONPreferences x
func WriteJSONPreferences(JSONPreferences JSONPreferences) {
	file, _ := json.MarshalIndent(JSONPreferences, "", " ")

	_ = ioutil.WriteFile("settings.json", file, 0644)
}

//StartBujiSequence initializes the process to add another buji in the database
func StartBujiSequence() {
	filename := SearchCsvInCurrentDirectory()
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	//initialize bufio consoleReader
	consoleReader := bufio.NewReader(os.Stdin)
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
	values[1], err = consoleReader.ReadString('\n')
	prompt("Qualita': ")
	values[2], err = consoleReader.ReadString('\n')
	prompt("Utilizzo: ")
	values[3], err = consoleReader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	//removing the newline char from consoleReader.ReadString
	for i := 1; i <= 3; i++ {
		values[i] = strings.TrimSuffix(values[i], "\n")
	}
	s := "\n" + values[0] + "," + values[1] + "," + values[2] + "," + values[3] + "," + values[4] + ","
	f.WriteString(s)
}

//SearchCsvInCurrentDirectory searches .csv files by getting an array of the elements present in that directory
func SearchCsvInCurrentDirectory() string {
	files, err := filepath.Glob("*")
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
