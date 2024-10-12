package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Calc struct {
	operand1 int
	operand2 int
	operator string
}

var romanToArabic = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func (c *Calc) Add() int {
	return c.operand1 + c.operand2
}

func (c *Calc) Sub() int {
	return c.operand1 - c.operand2
}

func (c *Calc) Multiply() int {
	return c.operand1 * c.operand2
}

func (c *Calc) Divide() int {
	if c.operand2 != 0 {
		return c.operand1 / c.operand2
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
	SIMPLE = "Вывод ошибки, так как число не простое."
)

func get_input() *Calc {
	var input_string string

	fmt.Scanln(&input_string)
	input_string = strings.Replace(input_string, " ", "", -1)

	if len(input_string) != 3 {
		panic(HIGH)
	}
	if input_string[1] != '=' || input_string[1] != '+' || input_string[1] != '-' || input_string[1] != '*' || input_string[1] != '/' {
		panic(LOW)
	}
	if unicode.IsDigit(rune(input_string[0])) && unicode.IsDigit(rune(input_string[2])) {
		if (input_string[0] < 0 && input_string[0] > 10) || (input_string[2] < 0 && input_string[2] > 10) {
			panic(RANGE)
		} else if input_string[0]%1 == '0' && input_string[2]%1 == '0' {
			panic(SIMPLE)
		}
	} else if !is_roman_num(string(input_string[0])) && !is_roman_num(string(input_string[0])) {
		panic(RANGE)
	} else if unicode.IsDigit(rune(input_string[0])) && !unicode.IsDigit(rune(input_string[2])) || !unicode.IsDigit(rune(input_string[0])) && unicode.IsDigit(rune(input_string[2])) {
		panic(SCALE)
	}

	calculator := &Calc{
		operand1: int(input_string[0]),
		operand2: int(input_string[2]),
		operator: string(input_string[1]),
	}

	return calculator
}

func main() {
	input := get_input()

	var result int

	if unicode.IsDigit(rune(input.operand1)) && unicode.IsDigit(rune(input.operand2)) {
		switch input.operator {
		case "+":
			result = input.Add()
		case "-":
			result = input.Add()
		case "*":
			result = input.Add()
		case "/":
			result = input.Add()
		default:
			fmt.Println("Неизвестный оператор")
			return
		}
	} else if is_roman_num(string(input.operand1)) && is_roman_num(string(input.operand2)) {
		switch input.operator {
		case "+":
			result = input.Add()
		case "-":
			result = input.Add()
		case "*":
			result = input.Add()
		case "/":
			result = input.Add()
		default:
			fmt.Println("Неизвестный оператор")
			return
		}
	}

	fmt.Println("Результат: ", result)
}
