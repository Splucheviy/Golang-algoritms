package palindrome

import "fmt"

func isPalindromeFirstCase(x int) bool {
	// Решение прыгает от уменьшения разрядности числа
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x > 0 {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return original == reversed
}

func IsPalindromeFirstCaseRunner()  {
	fmt.Println(isPalindromeFirstCase(121))
}
