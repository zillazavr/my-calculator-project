package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var operators = []string{"+", "-", "*", "/"}
var m = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
	"XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19, "XX": 20,
}

var m2 = map[int]string{
	0: "", 1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
	20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC", 100: "C",
}

func ArabToRoma(x int) string {

	if x <= 10 {
		return m2[x]
	} else {
		return m2[x-(x%10)] + m2[x%10]
	}

}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		var opers []string
		var v string
		var calc int
		romaOperOne, romaOperTwo := false, false

		for _, v = range operators {
			opers = strings.Split(text, v)
			if len(opers) == 2 {
				opers[0] = strings.TrimSpace(opers[0])
				opers[1] = strings.TrimSpace(opers[1])
				break
			}
		}

		if len(opers) == 2 {
			arabOne, err := strconv.Atoi(opers[0])
			if err != nil {
				_, ok := m2[arabOne]
				if ok {
					romaOperOne = true
				} else {
					fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
					return
				}
			}

			arabTwo, err2 := strconv.Atoi(opers[1])
			if err2 != nil {
				_, ok := m2[arabTwo]
				if ok {
					romaOperTwo = true
				} else {
					fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
					return
				}
			}

			if romaOperOne != romaOperTwo {
				fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
				return
			}

			if romaOperOne && romaOperTwo {
				opOne, okOne := m[opers[0]]
				opTwo, okTwo := m[opers[1]]
				if !okOne || !okTwo {
					fmt.Println("Вывод ошибки, так как один или оба операнда не являютя цифрами или больше 10")
					return
				}

				if opOne > 10 || opTwo > 10 {
					fmt.Println("Вывод ошибки, так как один или оба операнда больше 10")
					return
				}

				switch v {
				case "+":
					calc = opOne + opTwo
				case "-":
					calc = opOne - opTwo
				case "*":
					calc = opOne * opTwo
				case "/":
					calc = opOne / opTwo
				}

				if calc < 1 {
					fmt.Println("Вывод ошибки, так как римские цисла не могут быть отрицательными")
					return
				}

				fmt.Println(ArabToRoma(calc))
				continue

			} else {

				if arabOne > 10 || arabTwo > 10 {
					fmt.Println("Вывод ошибки, так как один или оба операнда больше 10")
					return
				}

				switch v {
				case "+":
					calc = arabOne + arabTwo
				case "-":
					calc = arabOne - arabTwo
				case "*":
					calc = arabOne * arabTwo
				case "/":
					calc = arabOne / arabTwo
				}

				fmt.Println(calc)
				continue

			}
		} else {
			fmt.Printf("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			return
		}
	}

}
