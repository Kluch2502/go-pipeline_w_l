package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    
	fmt.Print("Введите выражение: ")

    reader := bufio.NewReader(os.Stdin)
    expression, err := reader.ReadString('\n')
    if err != nil && len(expression) == 0 {
        fmt.Println("Ошибка чтения:", err)
        return
    }

    expression = strings.TrimSpace(expression)

    if err := isMistake(expression); err != nil {
        fmt.Println("Некорректное выражение:", err)
        return
    }

    result := Calculation(expression)
    fmt.Println("Результат:", result)
	
}

func Calc(expression string) (float64, error) {
	err := isMistake(expression)

	if err != nil {
		return 0, err
	} else {

		return Calculation(expression), nil

	}
}

func Calculation(expression string) float64 {
	var opstack []string
	var outstack []string
	sliseofTokens := stringtoTokens(expression)
	for _, k := range sliseofTokens {
		if isNum(k) {
			outstack = append(outstack, k)
		}
		if k == "(" {
			opstack = append(opstack, k)
		}

		if k == ")" {
			for len(opstack) > 0 {
				if opstack[len(opstack)-1] == "(" {
					opstack = opstack[:len(opstack)-1]
					break
				}
				outstack = append(outstack, opstack[len(opstack)-1])
				opstack = opstack[:len(opstack)-1]

			}
		}

		if isOperator(k) {
			for len(opstack) > 0 && getPrioritet(opstack[len(opstack)-1]) >= getPrioritet(k) {
				outstack = append(outstack, opstack[len(opstack)-1])
				opstack = opstack[:len(opstack)-1]
			}
			opstack = append(opstack, k)
		}
	}

	for len(opstack) > 0 {
		outstack = append(outstack, opstack[len(opstack)-1])
		opstack = opstack[:len(opstack)-1]
	}

	var operandStack []float64

	for i, k := range outstack {
		if isNum(outstack[i]) {
			num, _ := strconv.ParseFloat(k, 64)
			operandStack = append(operandStack, num)
		}
		if isOperator(outstack[i]) {
			num1 := operandStack[len(operandStack)-1]
			num2 := operandStack[len(operandStack)-2]
			operandStack = operandStack[:len(operandStack)-2]
			switch outstack[i] {
			case "-":
				operandStack = append(operandStack, num2-num1)
			case "*":
				operandStack = append(operandStack, num2*num1)
			case "+":
				operandStack = append(operandStack, num2+num1)
			default:
				operandStack = append(operandStack, num2/num1)
			}

		}
	}

	return operandStack[0]
}

func stringtoTokens(expression string) []string {
	var tokens []string
	var fullNum string

	for i := 0; i < len(expression); i++ {

		char := expression[i]

		if isNum(string(char)) {
			fullNum += string(char)
		} else {
			if len(fullNum) > 0 {
				tokens = append(tokens, fullNum)
				fullNum = ""
			}

			if isOperator(string(char)) || char == '(' || char == ')' {
				tokens = append(tokens, string(char))
			}
		}
	}

	if len(fullNum) > 0 {
		tokens = append(tokens, fullNum)
	}

	return tokens

}

func getPrioritet(op string) int {
	switch op {
	case "+", "-":
		return 1
	case "*", "/":
		return 2
	}
	return 0
}

func isOperator(op string) bool {
	return op == "+" || op == "-" || op == "*" || op == "/"
}

func isNum(s string) bool {

	_, err := strconv.Atoi(s)

	return err == nil

}

func toSliceOfNumbers(s string) []string {

	input := strings.ReplaceAll(s, " ", "")
	input = strings.ReplaceAll(s, "+", " ")
	input = strings.ReplaceAll(input, "-", " ")
	input = strings.ReplaceAll(input, "/", " ")
	input = strings.ReplaceAll(input, "*", " ")
	input = strings.ReplaceAll(input, "(", " ")
	input = strings.ReplaceAll(input, ")", " ")

	return strings.Fields(input)

}

func isMistake(s string) error {

	var countopen int
	s = strings.ReplaceAll(s, " ", "")

	for _, k := range s {
		if countopen < 0 {
			return fmt.Errorf("Неверный формат выражения 1")
		}
		switch k {
		case '(':
			countopen++
		case ')':
			countopen--
		}
	}
	if countopen != 0 {
		return fmt.Errorf("Неверный формат выражения 2")
	}

	if isOperator(string(s[0])) || isOperator(string(s[len(s)-1])) {
		return fmt.Errorf("Неверный формат выражения 3")
	}

	for i := 1; i < len(s); i++ {
		if isOperator(string(s[i])) && isOperator(string(s[i-1])) {

			return fmt.Errorf("Неверный формат выражения 4")
		}
	}

	for i := 0; i < len(s); i++ {

		if (s[i] == ')' && i > 0 && isOperator(string(s[i-1]))) ||
			(s[i] == '(' && i < len(s)-1 && isOperator(string(s[i+1])) && s[i+1] != '-') {
			return fmt.Errorf("Неверный формат выражения 5")
		}
	}

	numbers := toSliceOfNumbers(s)

	for _, k := range numbers {
		if !isNum(string(k)) {
			return fmt.Errorf("Неверный формат выражения 6 %s", string(k))
		}

	}

	return nil

}
