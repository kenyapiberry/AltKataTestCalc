package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Калькулятор умеет выполнять операции сложения строк, \n" +
		"вычитания строки из строки, умножения строки на число \n" +
		"и деления строки на число: " +
		"\n \"a\" + \"b\", \"a\" - \"b\", \"a\" * b, \"a\" / b. " +
		"\n Значения строк, передаваемых в выражении, выделяются двойными кавычками." +
		"\n Данные передаются в одну строку: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	data := strings.Fields(input)

	_, err := strconv.Atoi(data[0])
	if err == nil {
		panic("Первое слогаемое должно быть строкой")
	}

	//Проверяем, что строка начинается и заканчивается на одинарные кавычки
	if !strings.HasPrefix(input, "\"") && !strings.HasSuffix(input, "\"") {
		fmt.Println("Выделите двойными кавычками строковые значения")
		return
	}

	num, err := strconv.Atoi(data[2])
	if err == nil && num < 1 || num > 10 {
		fmt.Println("Пожалуйста, введите число от 1 до 10")
		return
	}

	operator := data[1]
	a := data[0]
	b := data[2]

	if strings.Contains(a, "") && strings.Contains(b, "") {
		a = strings.ReplaceAll(a, "\"", "")
		b = strings.ReplaceAll(b, "\"", "")

	} else {
		panic("Одно или оба значения не содержат двойные кавычки")
	}

	var result string

	switch operator {
	case "+":
		result = sumStrings(a, b)
	case "-":
		result = subtractStrings(a, b)
	case "*":
		num, _ := strconv.Atoi(b)
		result = multiplyStringByNumber(a, num)
	case "/":
		num, _ := strconv.Atoi(b)
		result = divideStringByNumber(a, num)
	default:
		fmt.Println("Некорректный ввод математической операции")
		return
	}

	const maxLength = 40
	if len(result) > maxLength {
		cutResult := result[:maxLength] + "..."
		fmt.Println(cutResult)
	} else {
		fmt.Println("Результат:", result)
	}

}

func sumStrings(a, b string) string {
	return a + b
}

func subtractStrings(a, b string) string {
	return strings.Replace(a, b, "", 1)
}

func multiplyStringByNumber(a string, b int) string {
	result := ""
	for i := 0; i < b; i++ {
		result += a
	}
	return result
}

func divideStringByNumber(a string, b int) string {
	return a[:len(a)/b]
}
