package main

import (
	"errors"
	"fmt"
	"slices"
)

func main () {
	var num int

	fmt.Printf("Inserisci un numero: ")
	fmt.Scanf("%d", &num)

	romanValue, err := decimalToRoman(num)

	if err != nil {
		fmt.Printf("ERRORE: %s\n", err)

		return
	}

	fmt.Printf("Il valore %d in numero romano è %s\n", num, romanValue)


	/* **************************************** */

	for i := range 100 {
		if i % 3 == 0 {
			fmt.Printf("%d. Deficiente\n", i)

			continue
		}

		fmt.Printf("%d. Sei scemo\n", i)
	}
}

func decimalToRoman (decimalValue int) (romanValue string, err error) {
	if (decimalValue < 1) || (decimalValue > 3999) {
		return "", errors.New("out of range (1 - 3999)")
	}

	romanMap := map[int]string{
		1: "I",
		4: "IV",
		5: "V",
		9: "IX",
		10: "X",
		40: "XL",
		50: "L",
		90: "XC",
		100: "C",
		400: "CD",
		500: "D",
		900: "CM",
		1000: "M",
	}

	romanKeys := make([]int, 0)

	for key := range romanMap {
		romanKeys = append(romanKeys, key)
	}

	slices.Sort(romanKeys)

	for decimalValue > 0 {
		for i := range len(romanKeys) {
			if (i == len(romanKeys) - 1) || ((decimalValue >= romanKeys[i]) && (decimalValue < romanKeys[i + 1])) {
				decimalValue -= romanKeys[i]
				romanValue += romanMap[romanKeys[i]]

				break
			}
		}
	}

	return
}