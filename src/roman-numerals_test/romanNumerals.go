package romannumeralstest

import (
	"errors"
	"slices"
)

func DecimalToRoman(decimalValue int) (romanValue string, err error) {
	if (decimalValue < 1) || (decimalValue > 3999) {
		return "", errors.New("out of range (1 - 3999)")
	}

	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	romanKeys := make([]int, 0)

	for key := range romanMap {
		romanKeys = append(romanKeys, key)
	}

	slices.Sort(romanKeys)

	for decimalValue > 0 {
		for i := range len(romanKeys) {
			if (i == len(romanKeys)-1) || ((decimalValue >= romanKeys[i]) && (decimalValue < romanKeys[i+1])) {
				decimalValue -= romanKeys[i]
				romanValue += romanMap[romanKeys[i]]

				break
			}
		}
	}

	return
}