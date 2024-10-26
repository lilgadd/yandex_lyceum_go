package main

import (
	"fmt"
	"strconv"
)
func Calc(expression string) (float64, error){
	var numbers []float64
	var operators []byte
	var num string
	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' {
			sk := 1
			startsk := i
			for j := i + 1; j < len(expression); j++ {
				if expression[j] == '(' {
					sk++
				} else if expression[j] == ')' {
					sk--
					if sk == 0 {
						subres, err := Calc(expression[startsk+1:j])
						if err != nil {
							return 0, err
						}
						numbers = append(numbers, subres)
						i = j
						break
					}
				}
			}
			if sk != 0 {
				return 0, fmt.Errorf("несбалансированные скобки")
			}
		} else if expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.'{
			num += string(expression[i])
		} else {
			if num != ""{
				n, err := strconv.ParseFloat(num, 64)
				if err != nil{
					return 0, err
				}
				numbers = append(numbers, n)
				num = ""
			}
			operators = append(operators, expression[i])
		}
	}
	if num != ""{
		n, err := strconv.ParseFloat(num, 64)
		if err != nil{
			return 0, err
		}
		numbers = append(numbers, n)
	}

	for i := 0; i < len(operators); i++{
		if i+1 >= len(numbers) {
            return 0, fmt.Errorf("недостаточно чисел для операции")
        }
		if operators[i] == '*' || operators[i] == '/'{
			if operators[i] == '*'{
				numbers[i] *= numbers[i + 1]
			} else{
				if numbers[i + 1] == 0{
					return 0, fmt.Errorf("деление на ноль")
				}
				numbers[i] /= numbers[i + 1]
			}
			numbers = append(numbers[:i+1], numbers[i+2:]...)
			operators = append(operators[:i], operators[i+1:]...)
			i--
		}
	}
	if len(numbers) == 0{
		return 0, fmt.Errorf("нет чисел для вычисления")
	}

	result := numbers[0]
	for i, op := range operators{
		if op == '+'{
			result += numbers[i + 1]
		} else if op == '-'{
			result -= numbers[i + 1]
		}
	}
	return result, nil
}