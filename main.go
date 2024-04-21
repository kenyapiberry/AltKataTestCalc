package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func main() {
	fmt.Print("Калькулятор умеет выполнять операции сложения строк, \n" +
		"вычитания строки из строки, умножения строки на число \n" +
		"и деления строки на число: " +
		"\n \"a\" + \"b\", \"a\" - \"b\", \"a\" * b, \"a\" / b. " +
		"\n Значения строк, передаваемых в выражении, выделяются двойными кавычками." +
		"\n Данные передаются в одну строку: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	input = strings.TrimSuffix(input, "\n")

	//Проверяем, что строка начинается и заканчивается на одинарные кавычки
	if !strings.HasPrefix(input, "") && !strings.HasSuffix(input, "") {
		fmt.Println("Выделите двойными кавычками строковые значения")
		return
	}

	data := strings.Fields(input)

	//for i, word := range data {
	//	data[i] = strings.Trim(word, "")
	//}

	if len(data) != 3 {
		fmt.Println("Некорректный формат ввода")
		return
	}

	// Удаляем кавычки из введенной строки
	input = input[1 : len(input)-1]

	// Преобразование чисел и оператора
	//num, err := strconv.Atoi(data[2])
	//if err != nil || num < 1 || num > 10 {
	//	fmt.Println("Пожалуйста, введите число от 1 до 10")
	//	return
	//}

	//num1 = num1[1 : len(num1)-1]
	//num2 = num2[1:len(num2)]

	operator := strings.TrimPrefix(data[1], "")
	a := data[0]
	b := strings.TrimSuffix(data[2], "")

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
		fmt.Println("Некорректный ввод")
		return
	}

	fmt.Println("Результат:", result)
}
