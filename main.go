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

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		for v := range operators {
			opers := strings.Split(text, operators[v])
			if len(opers) == 2 {
				one, err := strconv.Atoi(opers[0])
				if err != nil {
					fmt.Print(err)

				}

				fmt.Printf("Введене цифра %v\n", one)
				return
			}
		}

		fmt.Printf("Введенный пример не содержит операторов + - * /")
		return
	}

}
