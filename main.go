package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Это простой калькулятор")
	fmt.Println(`Введи выражение в формате "n1 s n2",
где "n1" и "n2" цифры от 1 до 10, а "s" знак вычисления (+-/*).
В том числе, могут использоваться римские цифры.`)
	n1, s, n2 := expression()

	if isDigit(n1) && isDigit(n2) && isSymbolNorm(s) {
		result := calcArabNum(n1, s, n2)
		if result == 404 {
			fmt.Println("Нарушен диапозон чисел(числа).")
		} else {
			fmt.Println(result)
		}
	} else if isRomanNumb(n1) && isRomanNumb(n2) && isSymbolNorm(s) {
		result := calcRomanNum(n1, s, n2)
		fmt.Println(result)
	} else {
		fmt.Println("Выражение записано не верно.")
	}

}

// Разделение строки на 3 части
func expression() (string, string, string) {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	m := sc.Text()
	t := strings.Split(m, " ")
	if len(t) == 3 {
		number1 := t[0]
		number2 := t[2]
		sign := t[1]
		return number1, sign, number2
	}
	return "Вы ввели неправильное количество символов.", "", ""
}

// Проверка строки на арабские цифры
func isDigit(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

// Проверка строки на римские цифры
func isRomanNumb(nR string) bool {
	listSymbols := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := range listSymbols {
		if listSymbols[i] == nR {
			return true
		}
	}
	return false
}

// Проверка арифметического знака
func isSymbolNorm(s string) bool {
	strSymbols := []string{"+", "-", "*", "/"}
	for i := range strSymbols {
		if strSymbols[i] == s {
			return true
		}
	}
	return false
}

// Вычисление арабских чисел
func calcArabNum(n1, s, n2 string) int {
	num1, _ := strconv.Atoi(n1)
	num2, _ := strconv.Atoi(n2)
	if num1 <= 10 && num1 >= 1 && (num2 <= 10 && num2 >= 1) {
		switch s {
		case "+":
			return num1 + num2
		case "-":
			return num1 - num2
		case "*":
			return num1 * num2
		case "/":
			return num1 / num2
		}
	}
	return 404
}

// Вычисление римских чисел
func calcRomanNum(nR1, s, nR2 string) string {
	numA1 := strconv.Itoa(convertRToA(nR1))
	numA2 := strconv.Itoa(convertRToA(nR2))
	resultArab := calcArabNum(numA1, s, numA2)
	if resultArab > 0 {
		resultRoman := convertAToR(resultArab)
		return resultRoman
	} else {
		return "В римской системе нет отрицательных чисел"
	}
}

// Преобразование римской цифры в арабскую
func convertRToA(numberRoman string) int {
	var romanNumb = map[rune]int{'I': 1, 'V': 5, 'X': 10}
	if len(numberRoman) == 0 {
		return 0
	}
	firstNumb := romanNumb[rune(numberRoman[0])]
	if len(numberRoman) == 1 {
		return firstNumb
	}
	nextNumb := romanNumb[rune(numberRoman[1])]
	if nextNumb > firstNumb {
		return (nextNumb - firstNumb) + convertRToA(numberRoman[2:])
	}
	return firstNumb + convertRToA(numberRoman[1:])
}

// Преобразование арабского числа в римское
func convertAToR(numberArab int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for numberArab >= conversion.value {
			roman.WriteString(conversion.digit)
			numberArab -= conversion.value
		}
	}

	return roman.String()
}
