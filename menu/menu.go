package menu

import (
	"fmt"

	color "github.com/fatih/color"
	i "github.com/punkycommunist/bujigo/io"
	s "github.com/punkycommunist/bujigo/structures"
)

//PrintMenu is a general stats printout
func PrintMenu(date []string, quantity []float64, quality []string, method []string, hour []int, remains float64) {
	jsp := i.ReadJSONPreferences()
	timeInterval := date[1] + " - " + date[len(date)-1]
	rounded := s.RoundedAvgQuantity(quantity, date, hour)
	sRounded := fmt.Sprintf("%.2f", rounded)
	smokedBuji := s.BujiNumber(date)
	color.Set(color.FgYellow)
	fmt.Printf("Intervallo " + timeInterval + "\n")
	color.Unset()
	fmt.Println("Buji fumati: " + fmt.Sprint(smokedBuji))
	fmt.Println("Media quantita' materiale: " + sRounded)
	fmt.Print("Media buji al giorno: ")
	fmt.Println(fmt.Sprintf("%.2f", s.BujiNumber(date)/s.TotalDaysElapsed(date, hour)))
	fmt.Print("Ora piu' frequente: ")
	fmt.Println(s.BestHour(hour))
	fmt.Println("Quantita' media al giorno: " + fmt.Sprintf("%.2f", s.DailyAvgQty(date, quantity, hour)))
	if s.RemainingDaysAtRate(date, quantity, hour, remains) >= 2.0 {
		color.Set(color.FgGreen)
		fmt.Println("Giorni rimasti a questo regime: " + fmt.Sprintf("%.2f", s.RemainingDaysAtRate(date, quantity, hour, remains)))
		color.Unset()
	}
	fmt.Println("Quantita' rimasta da fumare: " + fmt.Sprintf("%.2f", remains))
	fmt.Printf("Oggi: [")
	color.Set(color.FgYellow)
	fmt.Printf("%.2f", s.SmokedToday(date, quantity, hour))
	color.Unset()
	fmt.Printf("/")
	color.Set(color.FgRed)
	fmt.Printf("%.2f", s.DailyAvgQty(date, quantity, hour))
	color.Unset()
	fmt.Printf("] Rimanenti: [")
	color.Set(color.FgGreen)
	fmt.Printf("%.2f", s.DailyAvgQty(date, quantity, hour)-s.SmokedToday(date, quantity, hour))
	color.Unset()
	fmt.Printf("]")
	SpecialFunctions(jsp, date, quantity, quality, method, hour, remains)
}

//SpecialFunctions is a menu with a for loop that operates "special functions"
func SpecialFunctions(jsp i.JSONPreferences, date []string, quantity []float64, quality []string, method []string, hour []int, remains float64) {
	var selection string
	color.Green("\n\nFunzioni speciali!")
	for selection != "q" {
		fmt.Println("[a] Aggiungi un buji!.")
		fmt.Println("[s] Mostra gli ultimi buji!")
		fmt.Println("[c] Per calcolare i giorni rimanenti fumando una certa quantita' al giorno.")
		fmt.Println("[h] Quanto devo fumare per farmi durare il materiale per un numero personalizzato di giorni?")
		fmt.Println("[help]")
		prompt("[d] Per mostrare le preferenze dei colori!")
		fmt.Scanf("%s", &selection)
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
			fmt.Scan(&n)
			s.ShowLastBujis(date, quantity, quality, method, hour, remains, n)
			break

		case "c":
			quantitaAlGiorno := 0.0
			prompt("Quale sarebbe la quantita' al giorno?")
			fmt.Scan(&quantitaAlGiorno)
			s.HowManyDaysWithCustom(quantity, remains, quantitaAlGiorno)
			break

		case "h":
			giorni := 0.0
			prompt("Di quanti giorni stiamo parlando?")
			fmt.Scan(&giorni)
			s.HowMuchQuantityWithCustomDays(quantity, remains, giorni)
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
