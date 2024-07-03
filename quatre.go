package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertNumber(i int) string {
	var decimals = []string{"zéro", "un", "deux", "trois", "quatre", "cinq", "six", "sept", "huit", "neuf", "dix", "onze", "douze", "treize", "quatorze", "quinze", "seize", "dix-sept", "dix-huit", "dix-neuf"}
	var tens = []string{"", "dix", "vingt", "trente", "quarante", "cinquante", "soixante", "septante", "huitante", "nonante"}
	var hundreds = []string{"", "cent", "deux-cent", "trois-cent", "quatre-cent", "cinq-cent", "six-cent", "sept-cent", "huit-cent", "neuf-cent"}
	var thousands = []string{"", "mille", "deux-mille", "trois-mille", "quatre-mille", "cinq-mille", "six-mille", "sept-mille", "huit-mille", "neuf-mille"}

	if len(strconv.Itoa(i)) == 1 {
		return decimals[i]
	} else if len(strconv.Itoa(i)) == 2 {
		if i%10 == 0 {
			return tens[i/10]
		} else if i < 20 {
			return decimals[i]
		} else if i%10 == 1 {
			return tens[i/10] + "-et-" + decimals[i%10]
		} else {
			return tens[i/10] + "-" + decimals[i%10]
		}
	} else if len(strconv.Itoa(i)) == 3 {
		if i%100 != 0 {
			return hundreds[i/100] + "-" + convertNumber(i%100)
		} else {
			return hundreds[i/100]
		}
	} else if len(strconv.Itoa(i)) == 4 {
		if i%1000 != 0 {
			return thousands[i/1000] + "-" + convertNumber(i%1000)
		} else {
			return thousands[i/1000]
		}
	}

	return ""

}

func main() {
	var i int
	var correctInput bool

	scanner := bufio.NewScanner(os.Stdin)

	for !correctInput {
		fmt.Print("Enter a number between 0 and 9999: ")

		if scanner.Scan() {
			input := scanner.Text()
			if num, err := strconv.Atoi(strings.TrimSpace(input)); err != nil {
				fmt.Println("Please enter a valid number")
			} else if num < 0 || num > 9999 {
				fmt.Println("You entered a number out of range")
			} else {
				i = num
				correctInput = true
			}
		} else {
			fmt.Println("Error reading input:", scanner.Err())
			break
		}
	}

	for i != 4 {
		var convertedNumber = convertNumber(i)
		var length = len(convertedNumber)
		println(convertedNumber, "est", convertNumber(length))
		i = length
	}

	fmt.Println("quatre est magique")
}
