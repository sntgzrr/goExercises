package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"sort"
	"strings"
	"unicode"
)

func main() {

}

func invertString(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(string(s[len(s)-i-1]))
	}
}

func charCounter(s string) {
	counter := make(map[rune]int)
	for _, char := range s {
		counter[char]++
	}
	for char, count := range counter {
		fmt.Printf("%c: %d\n", char, count)
	}
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func randomString(s string) {
	arrayS := strings.Split(s, "")
	rand.Shuffle(len(arrayS), func(i, j int) {
		arrayS[i], arrayS[j] = arrayS[j], arrayS[i]
	})
	for _, char := range arrayS {
		fmt.Print(char)
	}
}

func targetSum(list []int, target int) {
	for _, element := range list {
		comp := target - element
		if contains(list, comp) {
			fmt.Printf("Suma encontrada: %d + %d = %d\n", element, comp, target)
		}
	}
}

func contains(list []int, target int) bool {
	for _, element := range list {
		if element == target {
			return true
		}
	}
	return false
}

func fizzBuzz(num int) {
	var array []int
	for i := 1; i <= num; i++ {
		array = append(array, i)
	}

	for i := 0; i < len(array); i++ {
		if array[i]%3 == 0 && array[i]%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if array[i]%3 == 0 {
			fmt.Println("Fizz")
		} else if array[i]%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(array[i])
		}
	}
}

func restArray(list []int) []int {
	var arr []int
	for len(list) > 0 {
		arr = append(arr, len(list))
		sort.Ints(list)
		m := list[0]
		var newArray []int
		for _, element := range list {
			if element != m {
				newArray = append(newArray, element)
			}
		}
		list = newArray
	}
	return arr
}

func sumArrayElement(list []int) int {
	var sum int
	for i := 0; i < len(list); i++ {
		sum += list[i]
	}
	return sum
}

func asciiCode() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%c\n", i)
	}
}

func uniqueElement(list []int) int {
	var unique int
	for _, element := range list {
		for _, element2 := range list {
			if element != element2 {
				unique = element
			}
		}
	}
	return unique
}

func pairNumbers() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

/*-----------------Basic Go Certification-----------------*/

func ModifyString(str string) string {
	var result string
	for _, char := range str {
		if !unicode.IsDigit(char) && !unicode.IsSpace(char) {
			result += string(char)
		}
	}
	runes := []rune(result)
	for i, r := 0, len(runes)-1; i < r; i, r = i+1, r-1 {
		runes[i], runes[r] = runes[r], runes[i]
	}
	return string(runes)
}

type Manager struct {
	FullName       string `json:"full_name,omitempty"`
	Position       string `json:"position,omitempty"`
	Age            int32  `json:"age,omitempty"`
	YearsInCompany int32  `json:"years_in_company,omitempty"`
}

func EncodeManager(manager *Manager) (io.Reader, error) {
	sb, err := json.Marshal(manager)
	if err != nil {
		return nil, err
	}
	b := bytes.NewBuffer(sb)
	return b, err
}
