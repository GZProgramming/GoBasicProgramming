package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var userInput = ""
	fmt.Println("Введите сумму в рублях:")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	userInput = sc.Text()
	number, err := strconv.ParseInt(userInput, 0, 32)
	if err != nil {
		fmt.Println("Неверный ввод!")
	}
	fmt.Println(number, pluralize(number))

}

func pluralize(number int64) string {
	var numberDiv = number % 100
	if (numberDiv > 4) && (numberDiv < 21) {
		return "рублей"
	}
	if (number%10 > 1) && (number%10 < 5) {
		return "рубля"
	}
	if (number%10 > 4) && (number%10 <= 9) || (number%10 == 0) {
		return "рублей"
	}
	return "рубль"
}
