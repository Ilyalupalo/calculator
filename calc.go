package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например: 2 + 3 или II + III):")

	expression, _ := reader.ReadString('\n')
	expression = strings.ToUpper(expression)
	expression = strings.TrimSpace(expression)

	isRoman := hasRomanNumerals(expression)
	isArabic := hasArabicNumerals(expression)

	if isRoman && isArabic || (!isRoman && !isArabic) {
		fmt.Println("Ошибка: Выражение должно содержать либо только римские числа, либо только арабские числа.")
		return
	}

	num1, operator, num2, err := parseExpression(expression, isRoman)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	result, err := calculate(num1, operator, num2)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	if num1 > 10 || num1 < 1 || num2 > 10 || num2 < 1 {
		fmt.Println("Ошибка: Числа должны быть в диапазоне от 1 до 10.")
		return
	}
	if isRoman {
		if result <= 0 {
			fmt.Println("Ошибка: Римские числа не могут быть отрицательными или равны нулю.")
			return
		}
	}

	resultOutput := formatResult(result, isRoman)
	fmt.Println("Результат:", resultOutput)
}

func parseExpression(expression string, isRoman bool) (int, string, int, error) {
	parts := strings.Split(expression, " ")
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("Некорректный ввод. Введите выражение в формате число оператор число")
	}

	num1, err := parseNumber(parts[0], isRoman)
	if err != nil {
		return 0, "", 0, err
	}

	operator := parts[1]

	num2, err := parseNumber(parts[2], isRoman)
	if err != nil {
		return 0, "", 0, err
	}

	return num1, operator, num2, nil
}

func parseNumber(num string, isRoman bool) (int, error) {
	if isRoman {
		return convertRomanToArabic(num)
	}
	return strconv.Atoi(num)
}

func calculate(num1 int, operator string, num2 int) (int, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 != 0 {
			return num1 / num2, nil
		} else {
			return 0, fmt.Errorf("Делить на ноль нельзя")
		}
	default:
		return 0, fmt.Errorf("Некорректный ввод. Введите оператор +, -, /, *")
	}
}

func formatResult(result int, isRoman bool) string {
	if isRoman {
		romanResult, _ := convertArabicToRoman(result)
		return romanResult
	}
	return strconv.Itoa(result)
}

func hasRomanNumerals(expression string) bool {
	romanNumerals := []string{"I", "V", "X", "L", "C", "D", "M"}
	for _, numeral := range romanNumerals {
		if strings.Contains(expression, numeral) {
			return true
		}
	}
	return false
}

func hasArabicNumerals(expression string) bool {
	arabicNumerals := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, numeral := range arabicNumerals {
		if strings.Contains(expression, numeral) {
			return true
		}
	}
	return false
}

func convertRomanToArabic(romanNum string) (int, error) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	result := 0
	prevValue := 0

	for _, r := range romanNum {
		value, ok := romanNumerals[r]
		if !ok {
			return 0, fmt.Errorf("Введено некорректное римское число.")
		}

		if value > prevValue {
			result += value - 2*prevValue
		} else {
			result += value
		}

		prevValue = value
	}

	return result, nil
}

func convertArabicToRoman(arabicNum int) (string, error) {
	if arabicNum < 1 || arabicNum > 3999 {
		return "", fmt.Errorf("Введено некорректное арабское число.")
	}

	romanNumerals := []struct {
		Value  int
		Symbol string
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

	romanNum := ""

	for _, numeral := range romanNumerals {
		for arabicNum >= numeral.Value {
			romanNum += numeral.Symbol
			arabicNum -= numeral.Value
		}
	}

	return romanNum, nil
}
