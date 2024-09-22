package rometointeger

import "fmt"

func romeToInt(s string) int {
	rome := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0
	for i, v := range s {
		sum += rome[string(v)]
		if i != 0 {
			if rome[string(s[i-1])] < rome[string(v)] {
				sum -= 2 * rome[string(s[i-1])]
			}
		}
	}
	return sum
}

func RomeToIntRunner(){
	fmt.Println(romeToInt("MCMXCIV"))
}
