package menu

import (
	"fmt"
	"log"

	color "github.com/fatih/color"
	i "github.com/punkycommunist/bujigo/io"
	s "github.com/punkycommunist/bujigo/structures"
)

//PrintMenu is a general stats printout
func PrintMenu(c i.CsvFile) {
	jsp := i.ReadJSONPreferences()
	timeInterval := c.Date[1] + " - " + c.Date[len(c.Date)-1]
	color.Set(color.FgYellow)
	fmt.Printf("Intervallo " + timeInterval + "\n")
	color.Unset()
	fmt.Println("Buji fumati: " + fmt.Sprint(s.BujiNumber(c.Date)))
	fmt.Println("Media quantita' materiale: " + fmt.Sprintf("%.2f", s.RoundedAvgQuantity(c.Quantity, c.Date, c.Hour)))
	//QAvgBujiSmokedADay
	t := jsp.QAvgBujiSmokedADay
	switch v := s.BujiNumber(c.Date) / s.TotalDaysElapsed(c.Date, c.Hour); {
	case v >= t.Best:
		color.Set(color.FgRed)
		fmt.Println("Media buji al giorno: " + fmt.Sprintf("%.2f", s.BujiNumber(c.Date)/s.TotalDaysElapsed(c.Date, c.Hour)))
		color.Unset()
		break
	case v >= t.Worst && v < t.Best:
		color.Set(color.FgYellow)
		fmt.Println("Media buji al giorno: " + fmt.Sprintf("%.2f", s.BujiNumber(c.Date)/s.TotalDaysElapsed(c.Date, c.Hour)))
		color.Unset()
		break
	case v <= t.Worst:
		color.Set(color.FgGreen)
		fmt.Println("Media buji al giorno: " + fmt.Sprintf("%.2f", s.BujiNumber(c.Date)/s.TotalDaysElapsed(c.Date, c.Hour)))
		color.Unset()
		break
	}
	//BestHour
	fmt.Print("Ora piu' frequente: ")
	fmt.Println(s.BestHour(c.Hour))
	//DailyAvgQty
	t = jsp.QDayAverage
	switch v := s.DailyAvgQty(c.Date, c.Quantity, c.Hour); {
	case v >= t.Best:
		color.Set(color.FgRed)
		fmt.Println("Quantita' media al giorno: " + fmt.Sprintf("%.2f", s.DailyAvgQty(c.Date, c.Quantity, c.Hour)))
		color.Unset()
		break
	case v >= t.Worst && v < t.Best:
		color.Set(color.FgYellow)
		fmt.Println("Quantita' media al giorno: " + fmt.Sprintf("%.2f", s.DailyAvgQty(c.Date, c.Quantity, c.Hour)))
		color.Unset()
		break
	case v <= t.Worst:
		color.Set(color.FgGreen)
		fmt.Println("Quantita' media al giorno: " + fmt.Sprintf("%.2f", s.DailyAvgQty(c.Date, c.Quantity, c.Hour)))
		color.Unset()
		break
	}
	//RemainingDaysAtRate
	t = jsp.QRemaininingDays
	switch v := s.RemainingDaysAtRate(c.Date, c.Quantity, c.Hour, c.Remains); {
	case v >= t.Best:
		color.Set(color.FgGreen)
		fmt.Println("Giorni rimasti a questo regime: " + fmt.Sprintf("%.2f", s.RemainingDaysAtRate(c.Date, c.Quantity, c.Hour, c.Remains)))
		color.Unset()
		break
	case v >= t.Worst && v < t.Best:
		color.Set(color.FgYellow)
		fmt.Println("Giorni rimasti a questo regime: " + fmt.Sprintf("%.2f", s.RemainingDaysAtRate(c.Date, c.Quantity, c.Hour, c.Remains)))
		color.Unset()
		break
	case v <= t.Worst:
		color.Set(color.FgRed)
		fmt.Println("Giorni rimasti a questo regime: " + fmt.Sprintf("%.2f", s.RemainingDaysAtRate(c.Date, c.Quantity, c.Hour, c.Remains)))
		color.Unset()
		break
	}
	//Remains
	t = jsp.QRemains
	switch v := c.Remains; {
	case v >= t.Best:
		color.Set(color.FgGreen)
		fmt.Println("Quantita' rimasta da fumare: " + fmt.Sprintf("%.2f", c.Remains))
		color.Unset()
		break
	case v >= t.Worst && v < t.Best:
		color.Set(color.FgYellow)
		fmt.Println("Quantita' rimasta da fumare: " + fmt.Sprintf("%.2f", c.Remains))
		color.Unset()
		break
	case v <= t.Worst:
		color.Set(color.FgRed)
		fmt.Println("Quantita' rimasta da fumare: " + fmt.Sprintf("%.2f", c.Remains))
		color.Unset()
		break
	}
	//fumato oggi printout
	fmt.Printf("Oggi: [")
	color.Set(color.FgYellow)
	fmt.Printf("%.2f", s.SmokedToday(c.Date, c.Quantity, c.Hour))
	color.Unset()
	fmt.Printf("/")
	color.Set(color.FgRed)
	fmt.Printf("%.2f", s.DailyAvgQty(c.Date, c.Quantity, c.Hour))
	color.Unset()
	fmt.Printf("] Rimanenti: [")
	color.Set(color.FgGreen)
	fmt.Printf("%.2f", s.DailyAvgQty(c.Date, c.Quantity, c.Hour)-s.SmokedToday(c.Date, c.Quantity, c.Hour))
	color.Unset()
	fmt.Printf("]")
	SpecialFunctions(jsp, c)
}

//SpecialFunctions is a menu with a for loop that operates "special functions"
func SpecialFunctions(jsp i.JSONPreferences, c i.CsvFile) {
	var selection string
	color.Green("\n\nFunzioni speciali!")
	for selection != "q" {
		fmt.Println("[a] Aggiungi un buji!")
		fmt.Println("[s] Mostra gli ultimi buji!")
		fmt.Println("[c] Per calcolare i giorni rimanenti fumando una certa quantita' al giorno.")
		fmt.Println("[h] Quanto devo fumare per farmi durare il materiale per un numero personalizzato di giorni?")
		prompt("[q] Per uscire!")
		n, err := fmt.Scanf("%s\n", &selection)
		if err != nil || n != 1 {
			log.Fatal(err)
		}
		switch selection {
		case "d":
			showColorPreferences(jsp)
			break
		case "a":
			i.StartBujiSequence()
			break

		case "s":
			var n int
			prompt("Quanti ultimi buji?")
			n, err = fmt.Scanf("%d\n", &n)
			if err != nil || n != 1 {
				log.Fatal(err)
			}
			s.ShowLastBujis(c.Date, c.Quantity, c.Quality, c.Method, c.Hour, c.Remains, n)
			break

		case "c":
			quantitaAlGiorno := 0.0
			prompt("Quale sarebbe la quantita' al giorno?")
			n, err = fmt.Scanf("%f\n", &quantitaAlGiorno)
			if err != nil || n != 1 {
				log.Fatal(err)
			}
			s.HowManyDaysWithCustom(c.Quantity, c.Remains, quantitaAlGiorno)
			break

		case "h":
			giorni := 0.0
			prompt("Di quanti giorni stiamo parlando?")
			n, err = fmt.Scanf("%f\n", &giorni)
			if err != nil || n != 1 {
				log.Fatal(err)
			}
			s.HowMuchQuantityWithCustomDays(c.Quantity, c.Remains, giorni)
			break

		default:
		case "q":
			break
		}
	}
}

func prompt(s string) {
	fmt.Println(s)
	fmt.Printf("$ ")
}

func showColorPreferences(jsp i.JSONPreferences) {
	fmt.Printf("Quantita' media al giorno:\nvalore peggiore: [%.2f]\nvalore migliore: [%.2f]\n", jsp.QDayAverage.Worst, jsp.QDayAverage.Best)
	fmt.Printf("Rimanenze:\nvalore peggiore: [%.2f]\nvalore migliore: [%.2f]\n", jsp.QRemains.Worst, jsp.QRemains.Best)

}
