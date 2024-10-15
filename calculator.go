package main

import (
	"fmt"
	 "strings"
	 //"unicode"
	 "strconv"
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
			return false
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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func get_input() *Calc {
	var input_string string
    operator_list := []string{"+", "-", "*", "/"}
    var calculator *Calc

	fmt.Scanln(&input_string)
	input_string = strings.Replace(input_string, " ", "", -1)
	
	f := input_string[0:1]
	s := input_string[len(input_string)-1:]
	operator := input_string[1:2]
	
	if len(input_string) != 3 {
		panic(HIGH)
	}
	if !contains(operator_list, operator) {
			panic(LOW)
		}
	if (is_roman_num(f) && !is_roman_num(s)) || (is_roman_num(s) && !is_roman_num(f)) {
	    panic(SCALE)
	} else if !is_roman_num(f) && !is_roman_num(s)  {
	    f, err := strconv.Atoi(f)
	    s, err := strconv.Atoi(s)    
	    if err != nil {
        panic(RANGE)
    }
		if (f < 0 && f > 10) || (s < 0 && s > 10) {
			panic(RANGE)
		} else if f%1 == '0' && s%1 == '0' {
			panic(SIMPLE)
		}
		
		calculator := &Calc{
		operand1: f,
		operand2: s,
		operator: string(operator),
	}
	return calculator
	}else{
	    calculator := &Calc{
		operand1: romanToArabic[f],
		operand2: romanToArabic[s],
		operator: string(operator),
	}
	return calculator
	}
	
	return calculator
}


func main() {
	input := get_input()

	var result int
	
	if is_roman_num(string(input.operand1)) && is_roman_num(string(input.operand2)) {
		switch input.operator {
		case "+":
			result = input.Add()
		case "-":
		    result = input.Sub()
		case "*":
			result = input.Multiply()
		case "/":
			result = input.Divide()
		default:
			fmt.Println("Неизвестный оператор")
			return
		}
	}else {
		switch input.operator {
		case "+":
			result = input.Add()
		case "-":
		    result = input.Sub()
		case "*":
			result = input.Multiply()
		case "/":
			result = input.Divide()
		default:
			fmt.Println("Неизвестный оператор")
			return
		}
	}

	fmt.Println("Результат: ", result)
}
