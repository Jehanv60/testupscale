package soal1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func Soal1() {
	var result int
	var hasil int
	f, err := os.Open("soal1/angka.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {

		for _, v := range scanner.Text() {
			a := len(string(v))
			hasil += a
		}
		for _, v := range scanner.Text() {
			a := string(v)
			b, _ := strconv.Atoi(a)
			result += b
		}

	}
	fmt.Println("Total Angka Pada File", hasil)
	fmt.Println("Jumlah Semua Angka", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// var i int
	// var n int
	// var result int
	// var err error
	// file, err := os.Open("soal1/angka.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// defer file.Close()
	// for {
	// 	n, err = fmt.Fscanln(file, &i)
	// 	if n == 0 || err != nil {
	// 		break
	// 	}
	// }
	// fmt.Println("Total Angka Pada File", i)
	// for i > n {
	// 	result += i
	// 	i--
	// }
	// fmt.Println("Jumlah Semua Angka", result)

}
