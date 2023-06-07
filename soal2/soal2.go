package soal2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Soal2() {
	fmt.Print("Kata/Kalimat Yang Diinput : ")
	var kata string
	var pembalikKata string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		kata = scanner.Text()
	}

	for i := len(kata) - 1; i >= 0; i-- {
		pembalikKata += string(kata[i])
	}
	// strings.ToLower(kata) == strings.ToLower(pembalikKata) terkena statement pada unit test
	if strings.EqualFold(kata, pembalikKata) {
		fmt.Println("Merupakan Palindrom")
	} else {
		fmt.Println("Bukan Palindrom")
	}
}
