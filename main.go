package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var word = input()
	fmt.Printf("isPalindrome = %t\n", isPalindrome(word))
}

func input() string {
	fmt.Printf("Введите строку для проверки: ")
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

func isPalindrome(word string) bool {
	i := 0
	j := len(word) - 1
	for i != j {
		if word[i] != word[j] {
			return false
		}
		i++
		j--
	}
	return true
}