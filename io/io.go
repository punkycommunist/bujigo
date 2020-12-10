package io

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	c "github.com/fatih/color"
	"github.com/tcnksm/go-latest"
)

//Version is the version of the compiled source
const Version string = "2.0.2"

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
	},
	"QAvgBujiSmokedADay": {
		"worst": 3.0,
		"best": 5.0
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
	QRemainingDays struct {
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
	// files, err := filepath.Glob("*")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for i := 0; i < len(files); i++ {
	// 	if files[i] == "settings.json" {
	// 		jsonFile, err := os.Open("settings.json")
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		// read our opened jsonFile as a byte array.
	// 		byteValue, _ := ioutil.ReadAll(jsonFile)
	// 		// we unmarshal our byteArray which contains our
	// 		// jsonFile's content into 'users' which we defined above
	// 		json.Unmarshal(byteValue, &JSONPreferences)
	// 		jsonFile.Close()
	// 		return JSONPreferences
	// 	}
	// }
	jsonFile, err := os.Open("settings.json")
	if err != nil {
		var progress string
		prompt("settings.json non trovato. creare il default ora? [y] per si")
		n, err1 := fmt.Scanf("%s\n", &progress)
		if err1 != nil || n != 1 {
			log.Fatal(err1)
		}
		f, err := os.OpenFile("settings.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.WriteString(defaultSettings)
		if err != nil {
			log.Fatal(err)
		}
		jj := ReadJSONPreferences()
		return jj
	}
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &JSONPreferences)
	jsonFile.Close()
	return JSONPreferences

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
	sc := bufio.NewScanner(os.Stdin)
	var values [5]string
	var isToday string
	prompt("Il buji e' stato fumato oggi? [y] per si, [n] per no e inserire la data.")
	sc.Scan()
	isToday = sc.Text()
	if err != nil {
		log.Fatal(err)
	}
	if isToday == "y" { //if [y] gets pressed
		values[0] = time.Now().Format("02/01/2006")
	} else if isToday == "n" {
		prompt("Che giorno era?")
		sc.Scan()
		values[0] = sc.Text()
		if err != nil {
			log.Fatal(err)
		}
	} else { //if invalid character
		log.Fatal(isToday + " invalid.")
	}

	prompt("Il buji e' stato fumato a quest'ora? [y] per si, [n] per no e inserire l'ora.")
	sc.Scan()
	isToday = sc.Text()
	if err != nil {
		log.Fatal(err)
	}
	if isToday == "y" { //if [y] gets pressed
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
	} else if isToday == "n" {
		prompt("Che ore erano?")
		sc.Scan()
		values[4] = sc.Text()
	} else { //if invalid character
		log.Fatal(isToday + " invalid.")
	}
	prompt("Quantita': ")
	sc.Scan()
	values[1] = sc.Text()
	prompt("Qualita': ")
	sc.Scan()
	values[2] = sc.Text()
	prompt("Utilizzo: ")
	sc.Scan()
	values[3] = sc.Text()
	//removing the newline char from consoleReader.ReadString
	for i := 1; i <= 3; i++ {
		values[i] = strings.TrimSuffix(values[i], "\n")
	}
	s := "\n" + values[0] + "," + values[1] + "," + values[2] + "," + values[3] + "," + values[4] + ","
	f.WriteString(s)
}

//SearchCsvInCurrentDirectory searches .csv files by getting an array of the elements present in that directory
func SearchCsvInCurrentDirectory() string {
	files, fileErr := filepath.Glob("*")
	if fileErr != nil {
		log.Fatal(fileErr)
	}
	for i := 0; i < len(files); i++ {
		if files[i][len(files[i])-4:] == ".csv" {
			return files[i]
		}
	}
	sc := bufio.NewScanner(os.Stdin)
	log.Println("No .csv file found. Do you want to create the default [buji.csv]?")
	prompt("How much material do you have to smoke? (1g = 1.0): ")
	sc.Scan()
	material, err := strconv.ParseFloat(sc.Text(), 64)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile("buji.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.WriteString("giorno,quantita,qualita,tipo,ore," + fmt.Sprintf("%.2f", material))
	if err != nil {
		log.Fatal(err)
	}
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

//IsOnline checks for internet connectivity requesting an http connection to an arbitrary url
func IsOnline() bool {
	//Make a request to icanhazip.com
	//We need the error only, nothing else :)
	_, err := http.Get("https://icanhazip.com/")
	//err = nil means online
	if err == nil {
		return true
	}
	//if the "return statement" in the if didn't executed,
	//this one will execute surely
	return false
}

//CheckForUpdates checks using the library github.com/tcnksm/go-latest and prints if something is updated
func CheckForUpdates() {
	if IsOnline() {
		githubTag := &latest.GithubTag{
			Owner:             "punkycommunist",
			Repository:        "bujigo",
			FixVersionStrFunc: latest.DeleteFrontV(),
		}
		res, err := latest.Check(githubTag, Version)
		if err != nil {
			log.Fatal(err)
		}
		if res.Outdated {
			c.Set(c.FgYellow, c.BgRed)
			fmt.Printf("! Aggiornamento disponibile ! https://github.com/punkycommunist/bujigo/releases/tag/v%s\n", res.Current)

			c.Unset()
		} else {
			c.Set(c.FgHiBlue)
			fmt.Println("[v] " + Version)
			c.Unset()
		}
	} else {
		c.Set(c.FgHiBlue)
		fmt.Println("Nessuna connessione internet! Impossibile controllare aggiornamenti.")
		c.Unset()
	}
}

func prompt(s string) {
	fmt.Println(s)
	fmt.Printf("$ ")
}
