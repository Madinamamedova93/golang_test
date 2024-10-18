package main

import (
	"fmt"
	 "strings"
	 "strconv"
)

var is_roman bool

type Calc struct {
	operand1 int
	operand2 int
	operator string
}

var romanToArabic = map[string]int{
	"I": 1,
	"II": 2,
	"III": 3,
	"IV": 4,
	"V": 5,
	"VI": 6,
	"VII": 7,
	"VIII": 8,
	"IX": 9,
	"X": 10,
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

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
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
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
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

func operator_func(s []string, str string) string {
	for _, v := range s {
	    if strings.Contains(str, v) == true{
	        return v
	    } 
	}
    panic(HIGH)
}

func get_input() *Calc {
	var input_string string
    operator_list := []string{"+", "-", "*", "/"}
    var calculator *Calc

	fmt.Scanln(&input_string)
	input_string = strings.Replace(input_string, " ", "", -1)
	
	operator := operator_func(operator_list, input_string)
	index := strings.Index(input_string, operator)

	f := input_string[0:index]
	s := input_string[index+1:]

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
	is_roman = true
	return calculator
	}
	
	return calculator
}


func main() {
	input := get_input()

	var result int
	
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
		if is_roman{
		    if result < 0 {
		    panic(DIV)
		    }else if result ==0 {
		        panic(ZERO)
		    }
		    fmt.Println("Результат: ", integerToRoman(result))
		}else{
		    fmt.Println("Результат: ", result)
		}
}
