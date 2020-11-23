package menu

import (
	"fmt"

	color "github.com/fatih/color"
	i "github.com/punkycommunist/bujigo/io"
	s "github.com/punkycommunist/bujigo/structures"
)

//PrintMenu is a general stats printout
func PrintMenu(date []string, quantity []float64, quality []string, method []string, hour []int32, remains float64) {
	timeInterval := s.GetStringTimeInterval(date)
	rounded := s.GetRoundedAvgQuantity(quantity)
	sRounded := fmt.Sprintf("%.2f", rounded)
	smokedBuji := s.GetBujiNumber(date)
	color.Set(color.FgYellow)
	fmt.Printf("Intervallo " + timeInterval + "\n")
	color.Unset()
	fmt.Println("Buji fumati: " + fmt.Sprint(smokedBuji))
	fmt.Println("Media quantita' materiale: " + sRounded)
	fmt.Print("Media buji al giorno: ")
	fmt.Println(fmt.Sprintf("%.2f", float64(smokedBuji)/s.DaysElapsedFromLastBuji(date)))
	fmt.Print("Ora piu' frequente: ")
	fmt.Println(s.GetBestHour(hour))
	fmt.Println("Quantita' media al giorno: " + fmt.Sprintf("%.2f", s.GetDailyAvgQty(date, quantity)))
	fmt.Println("Giorni rimasti a questo regime: " + fmt.Sprintf("%.2f", s.GetRemainingDays(date, quantity, remains)))
	fmt.Println("Quantita' rimasta da fumare: " + fmt.Sprintf("%.2f", remains))
	fmt.Println("Oggi: " + s.HowMuchLeftString(date, quantity, hour))
	SpecialFunctions(date, quantity, quality, method, hour, remains)
}

//SpecialFunctions is a menu with a for loop that operates "special functions"
func SpecialFunctions(date []string, quantity []float64, quality []string, method []string, hour []int32, remains float64) {
	var selection string
	color.Green("\n\nFunzioni speciali!")
	for selection != "q" {
		fmt.Println("[a] Aggiungi un buji!.")
		fmt.Println("[s] Mostra gli ultimi buji!")
		fmt.Println("[c] Per calcolare i giorni rimanenti fumando una certa quantita' al giorno.")
		prompt("[h] Quanto devo fumare per farmi durare il materiale per un numero personalizzato di giorni?")
		fmt.Scanf("%s", &selection)
		switch selection {
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
