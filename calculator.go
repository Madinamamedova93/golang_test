package main

import (
	"fmt"
)

var romanToArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b != 0 {
		return a / b
	}

	fmt.Println("Ошибка деления на ноль!")
	return 0

}

func is_roman_num(s string) bool {
	for _, c := range s {
		if _, ok := romanToArabic[string(c)]; !ok {
			panic(RANGE)
		}
	}
	return true
}

const (
	LOW = "Вывод ошибки, так как строка " +
		"не является математической операцией."
	HIGH = "Вывод ошибки, так как формат математической операции " +
		"не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."
	SCALE = "Вывод ошибки, так как используются " +
		"одновременно разные системы счисления."
	DIV = "Вывод ошибки, так как в римской системе " +
		"нет отрицательных чисел."
	ZERO  = "Вывод ошибки, так как в римской системе нет числа 0."
	RANGE = "Калькулятор умеет работать только с арабскими целыми " +
		"числами или римскими цифрами от 1 до 10 включительно"
)

func get_input() (int, int, string) {
	var a, b int
	var operator string

	for i := 0; i < 2; i++ {

	}

	fmt.Print("Введите первое число: ")
	if a < 0 && a > 10  {
		panic(RANGE)
	}
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	if b < 0 && b > 10 {
		panic(RANGE)
	}
	fmt.Scanln(&b)

	fmt.Print("Введите оператор (+, -, *, /): ")
	if operator != "+" || operator != "-" || operator != "*" || operator != "/" {
		panic(LOW)
	}
	fmt.Scanln(&operator)

	return a, b, operator
}

func main() {
	a, b, operator := get_input()

	var result int

	switch operator {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		result = divide(a, b)
	default:
		fmt.Println("Неизвестный оператор")
		return
	}

	fmt.Println("Результат: ", result)
}
