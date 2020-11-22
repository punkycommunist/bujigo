package menu

import (
	"fmt"

	i "github.com/punkycommunist/bujigo/io"
	s "github.com/punkycommunist/bujigo/structures"
)

//PrintMenu x
func PrintMenu(date []string, quantity []float64, quality []string, method []string, hour []int32, remains float64) {
	timeInterval := s.GetStringTimeInterval(date)
	rounded := s.GetRoundedAvgQuantity(quantity)
	sRounded := fmt.Sprintf("%.2f", rounded)
	smokedBuji := s.GetBujiNumber(date)
	fmt.Println("Intervallo " + timeInterval)
	fmt.Println("Buji fumati: " + fmt.Sprint(smokedBuji))
	fmt.Println("Media quantita' materiale: " + sRounded)
	fmt.Print("Media buji al giorno: ")
	fmt.Println(fmt.Sprintf("%.2f", float64(smokedBuji)/s.NewGetDaysElapsed(date)))
	fmt.Print("Ora piu' frequente: ")
	fmt.Println(s.GetBestHour(hour))
	fmt.Println("Quantita' media al giorno: " + fmt.Sprintf("%.2f", s.GetDailyAvgQty(date, quantity)))
	fmt.Println("Giorni rimasti a questo regime: " + fmt.Sprintf("%.2f", s.GetRemainingDays(date, quantity, remains)))
	fmt.Println("Quantita' rimasta da fumare: " + fmt.Sprintf("%.2f", remains))
	fmt.Println("Oggi: " + s.HowMuchLeftString(date, quantity, hour))
	SpecialFunctions(date, quantity, quality, method, hour, remains)
}

//SpecialFunctions x
func SpecialFunctions(date []string, quantity []float64, quality []string, method []string, hour []int32, remains float64) {
	var selection string
	fmt.Println("\n\nFunzioni speciali!")
	for selection != "q" {
		fmt.Println("[a] Aggiungi un buji!")
		fmt.Println("[s] Mostra gli ultimi 5 buji!")
		fmt.Println("[c] Per calcolare i giorni rimanenti fumando una certa quantita' al giorno.")
		fmt.Println("[h] Quanto devo fumare per farmi durare il materiale per un numero personalizzato di giorni?")
		fmt.Scanf("%s", &selection)
		switch selection {
		case "a":
			i.StartBujiSequence()
			break

		case "s":
			var n int
			fmt.Println("Quanti ultimi buji?")
			fmt.Scan(&n)
			s.ShowLastFiveBujis(date, quantity, quality, method, hour, remains, n)
			break

		case "c":
			quantitaAlGiorno := 0.0
			fmt.Println("Quale sarebbe la quantita' al giorno?")
			fmt.Scan(&quantitaAlGiorno)
			s.HowManyDaysWithCustom(quantity, remains, quantitaAlGiorno)
			break

		case "h":
			giorni := 0.0
			fmt.Println("Di quanti giorni stiamo parlando?")
			fmt.Scan(&giorni)
			s.HowMuchQuantityWithCustomDays(quantity, remains, giorni)
			break

		default:
		case "q":
			break
		}
	}
}
