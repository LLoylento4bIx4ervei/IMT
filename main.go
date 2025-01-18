package main

import (
	"fmt"
)

func main() {
	quit := ""
	for {
		fmt.Println("Калькулятор индекса массы тела")
		height, kg := inputUser()
		imt := calculateIMT(height, kg)
		outputResult(imt)
		caseIMT(imt)
		answer := outProgram(quit)
		if answer == true {
			return
		}

	}
}

func inputUser() (float64, float64) {
	var height, kg float64
	fmt.Print("Введите ваш рост в сантиметрах: ")
	fmt.Scan(&height)
	fmt.Print("Введите ваш вес в килограммах: ")
	fmt.Scan(&kg)
	return height, kg
}

func calculateIMT(height, kg float64) float64 {
	IMT := kg / (height / 100 * height / 100)
	return IMT
}

func caseIMT(imt float64) {
	switch {
	case imt < 16:
		fmt.Println("Выраженный дефицит массы тела")
	case imt < 18.5:
		fmt.Println("Недостаточная (дефицит) масса тела")
	case imt < 25:
		fmt.Println("Норма")
	case imt < 30:
		fmt.Println("Избыточная масса тела")
	default:
		fmt.Println("Какая-то степень ожирения")
	}
}

func outputResult(imt float64) {
	fmt.Printf("Ваш индекс массы тела: %.0f\n", imt)
}

func outProgram(quit string) bool {
	fmt.Print("Для выхода нажмите q: ")
	fmt.Scan(&quit)
	if quit == "q" {
		return true
	}
	return false
}
