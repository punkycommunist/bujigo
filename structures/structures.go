package structures

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

// RoundedAvgQuantity returns how much did you smoke on average per buji
func RoundedAvgQuantity(quantity []float64, date []string, hour []int) float64 {
	sumQuantity := 0.0
	for i := 1; i < len(quantity); i++ {
		sumQuantity += quantity[i]
	}
	avgQuantity := sumQuantity / float64(BujiNumber(date))
	roundedAvgQuantity := math.Round(avgQuantity*100.0) / 100.0
	return roundedAvgQuantity
}

// BujiNumber returns how many bujis got smoked in total
func BujiNumber(dates []string) float64 {
	return float64(len(dates)) - 1.0
}

//TotalDaysElapsed returns how many days have passed from the start to the end of records
func TotalDaysElapsed(date []string, hour []int) float64 {
	fDay, err := strconv.Atoi(date[1][0:2])
	fMonth, err := strconv.Atoi(date[1][3:5])
	fYear, err := strconv.Atoi(date[1][6:10])
	// l := len(date) - 1
	// lDay, err := strconv.Atoi(date[l][0:2]) --> days counting until last report <--
	// lMonth, err := strconv.Atoi(date[l][3:5])
	// lYear, err := strconv.Atoi(date[l][6:10])
	// t2 := time.Date(lYear, time.Month(lMonth), lDay, int(hour[l]), 0, 0, 0, time.UTC)
	nowYear, err := strconv.Atoi(time.Now().Format("2006"))
	nowMonth, err := strconv.Atoi(time.Now().Format("01"))
	nowDay, err := strconv.Atoi(time.Now().Format("02"))
	nowHour, err := strconv.Atoi(time.Now().Format("15"))
	nowMinute, err := strconv.Atoi(time.Now().Format("4"))
	nowSecond, err := strconv.Atoi(time.Now().Format("5"))
	t, err := TimeIn(time.Now(), "Europe/Rome")
	if err != nil {
		log.Fatal("Europe/Rome", "<time unknown>")
	}
	t2 := time.Date(nowYear, time.Month(nowMonth), nowDay, nowHour, nowMinute, nowSecond, 0, t.Location())
	if err != nil {
		log.Println(err)
	}
	t1 := time.Date(fYear, time.Month(fMonth), fDay, int(hour[1]), 0, 0, 0, t.Location())
	days := t2.Sub(t1).Hours() / 24.0
	fmt.Println(t2)
	return days
}

//DaysElapsedFromLastBuji returns the number of days elapsed from the last buji on the record
func DaysElapsedFromLastBuji(date []string, hour []int) float64 {
	t1 := time.Now()
	l := len(date) - 1
	lYear, err := strconv.Atoi(date[l][6:10])
	lMonth, err := strconv.Atoi(date[l][3:5])
	lDay, err := strconv.Atoi(date[l][0:2])
	if err != nil {
		log.Fatal(err)
	}
	t2 := time.Date(lYear, time.Month(lMonth), lDay, int(hour[l]), 0, 0, 0, time.UTC)
	days := float64(t2.Sub(t1).Hours()) / 24.0
	return days
}

//BestHour returns the most occurrent hour
func BestHour(hour []int) int {
	var hours [24]int
	for i := 1; i < len(hour); i++ {
		hours[hour[i]] = hours[hour[i]] + 1
	}
	var max int
	var last int
	for i := 0; i < 24; i++ {
		if hours[i] > max {
			last = int(i)
			max = hours[i]
		}
	}
	return last
}

//DailyAvgQty returns how much you smoke a day on average
func DailyAvgQty(date []string, quantity []float64, hour []int) float64 {
	sum := 0.0
	for i := 1; i < len(quantity); i++ {
		sum += quantity[i]
	}
	return sum / TotalDaysElapsed(date, hour)
}

// RemainingDaysAtRate returns how many days you got left before finishing supplies at this rate
func RemainingDaysAtRate(date []string, quantity []float64, hour []int, remains float64) float64 {
	return remains / DailyAvgQty(date, quantity, hour)
}

// HowManyDaysWithCustom returns how many days the supplies are gonna last with a specific amount a day
func HowManyDaysWithCustom(quantity []float64, remains float64, quantityPerDay float64) float64 {
	fmt.Println("Giorni rimanenti a finire per uno " + fmt.Sprintf("%.2f", quantityPerDay) + " al giorno: " + fmt.Sprintf("%.2f", remains/quantityPerDay))
	return remains / quantityPerDay
}

// HowMuchQuantityWithCustomDays returns how much you would have to smoke a day to endure the supplies x days
func HowMuchQuantityWithCustomDays(quantity []float64, remains float64, days float64) float64 {
	fmt.Println("Dovrai fumare " + fmt.Sprintf("%.2f", remains/days))
	return remains / days
}

// ShowLastBujis x
func ShowLastBujis(date []string, quantity []float64, quality []string, method []string, hour []int, remains float64, n int) {
	fmt.Println("Data\t\tOra\tQuantita'\tQualita'\tTipo")
	for i := len(date) - 1; i > len(date)-1-n; i-- {
		fmt.Println(date[i] + "\t" + fmt.Sprintf("%d", hour[i]) + "\t" + fmt.Sprintf("%.2f", quantity[i]) + "\t\t" + quality[i] + "\t" + method[i])
	}
}

// SmokedToday returns [0]: how much you smoked today [1]: the average quantity you smoke a day [2]: how much you have left
func SmokedToday(date []string, quantity []float64, hour []int) float64 {
	csvLastCellIndex := len(date) - 1
	now := time.Now()
	var smokedToday float64
	lastDay := date[csvLastCellIndex]
	today := now.Format("02/01/2006")
	for i := csvLastCellIndex; i > 1; i-- {
		if lastDay == today {
			smokedToday += quantity[i]
			lastDay = date[i-1]
		}
	}
	return smokedToday
}
func prompt(s string) {
	fmt.Println(s)
	fmt.Printf("$ ")
}

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}
