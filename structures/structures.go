package structures

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"time"
)

// GetRoundedAvgQuantity returns how much did you smoke on average per day
func GetRoundedAvgQuantity(quantity []float64, date []string, hour []int32) float64 {
	sumQuantity := 0.0
	for i := 1; i < len(quantity); i++ {
		sumQuantity += quantity[i]
	}
	avgQuantity := sumQuantity / GetDaysElapsed(date, hour)
	roundedAvgQuantity := math.Round(avgQuantity*100.0) / 100.0
	return roundedAvgQuantity
}

// GetStringTimeInterval returns a string that represent the time interval between the first buji on record and the last one
func GetStringTimeInterval(dates []string) string {
	s1 := dates[1]
	s2 := dates[len(dates)-1]
	return s1 + " - " + s2
}

// GetBujiNumber returns how many bujis got smoked in total
func GetBujiNumber(dates []string) int {
	return len(dates) - 1
}

//GetDaysElapsed returns how many days have passed from the start of records
func GetDaysElapsed(date []string, hour []int32) float64 {
	l := len(date) - 1
	fDay, err := strconv.Atoi(date[1][0:2])
	fMonth, err := strconv.Atoi(date[1][3:5])
	fYear, err := strconv.Atoi(date[1][6:10])
	lDay, err := strconv.Atoi(date[l][0:2])
	lMonth, err := strconv.Atoi(date[l][3:5])
	lYear, err := strconv.Atoi(date[l][6:10])
	if err != nil {
		log.Println(err)
	}
	t1 := time.Date(fYear, time.Month(fMonth), fDay, int(hour[1]), 0, 0, 0, time.UTC)
	t2 := time.Date(lYear, time.Month(lMonth), lDay, int(hour[l]), 0, 0, 0, time.UTC)
	days := t2.Sub(t1).Hours() / 24
	return days
}

//DaysElapsedFromLastBuji returns the number of days elapsed from the last buji on the record
func DaysElapsedFromLastBuji(date []string) float64 {
	counter := 1.0
	for i := 2; i < len(date); i++ {
		if date[i-1] != date[i] {
			counter++
		}
	}
	return counter
}

//GetBestHour returns the most occurrent hour
func GetBestHour(hour []int32) int32 {
	var hours [24]int32
	for i := 1; i < len(hour); i++ {
		hours[hour[i]] = hours[hour[i]] + 1
	}
	var max int32
	var last int32
	for i := 0; i < 24; i++ {
		if hours[i] > max {
			last = int32(i)
			max = hours[i]
		}
	}
	return last
}

//GetDailyAvgQty returns how much you smoke a day on average
func GetDailyAvgQty(date []string, quantity []float64) float64 {
	sum := 0.0
	for i := 1; i < len(quantity); i++ {
		sum += quantity[i]
	}
	return sum / DaysElapsedFromLastBuji(date)
}

// GetRemainingDays returns how many days you got left before finishing supplies at this rate
func GetRemainingDays(date []string, quantity []float64, remains float64) float64 {
	return remains / GetDailyAvgQty(date, quantity)
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
func ShowLastBujis(date []string, quantity []float64, quality []string, method []string, hour []int32, remains float64, n int) {
	fmt.Println("Data\t\tOra\tQuantita'\tQualita'\tTipo")
	for i := len(date) - 1; i > len(date)-1-n; i-- {
		fmt.Println(date[i] + "\t" + fmt.Sprintf("%d", hour[i]) + "\t" + fmt.Sprintf("%.2f", quantity[i]) + "\t\t" + quality[i] + "\t" + method[i])
	}
}

// HowMuchLeftString x
func HowMuchLeftString(date []string, quantity []float64, hour []int32) string {
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
	avg := GetDailyAvgQty(date, quantity)

	return "[" + fmt.Sprintf("%.2f", smokedToday) + "/" + fmt.Sprintf("%.2f", avg) + "] [Rimanenti: " + fmt.Sprintf("%.2f", avg-smokedToday) + "]"
}
func prompt(s string) {
	fmt.Println(s)
	fmt.Printf("$ ")
}
