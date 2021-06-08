package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var userInput = ""
	fmt.Println("Введите сумму, проценты и срок в месяцах:")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	userInput = sc.Text()
	//fmt.Scanf("%s\n", &userInput)
	var result = calculate(userInput)
	if result < 0 {
		fmt.Println("Неверный ввод!")
	}
	fmt.Println("Ok ", result)
}

func calculate(userInput string) float64 {
	var numbers = strings.Fields(userInput)
	if len(numbers) != 3 {
		return -1
	}
	money, err := strconv.ParseFloat(numbers[0], 64)
	persent, err := strconv.ParseFloat(numbers[1], 64)
	persent = persent * 0.01 / 12
	months, err := strconv.ParseFloat(numbers[2], 64)

	if err != nil {
		return -1
	}
	return money * math.Pow(1+persent, months)
}
