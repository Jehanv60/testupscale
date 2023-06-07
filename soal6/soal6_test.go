package soal6

import (
	"fmt"
	"strings"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Fizz Buzz", i)
		} else if i%3 == 0 {
			fmt.Println("Fizz", i)
		} else if i%5 == 0 {
			fmt.Println("Buzz", i)
		} else {
			fmt.Println("angka", i)
		}

	}
}

func TestPalindromString(t *testing.T) {
	kata := "Kasur Ini Rusak"
	var pembalikKata string

	for i := len(kata) - 1; i >= 0; i-- {
		pembalikKata += string(kata[i])
	}
	// strings.ToLower(kata) == strings.ToLower(pembalikKata) terkena statement pada unit test

	if strings.EqualFold(kata, pembalikKata) {
		fmt.Println("Nama Ini Adalah Palindrom")
	} else {
		fmt.Println("Nama Ini Bukan Palindrom")
	}
}

func faktorial(n int) int {
	if n > 0 {
		result := n * faktorial(n-1)
		return result
	}
	return 1
}
func TestFaktorial(t *testing.T) {
	fmt.Println(faktorial(5))
}
func findMinAndMax(a []int) (min int, max int) {
	min = a[0]
	max = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}

		if value > max {
			max = value

		}

	}
	return min, max
}
func TestMinAndMax(t *testing.T) {
	var a = []int{11, -4, 7, 8, -5}
	min, max := findMinAndMax(a)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}

func TestHurufDanAngka(t *testing.T) {
	kata := "Jehan 123 hehe"
	fmt.Println(len(kata))
}
