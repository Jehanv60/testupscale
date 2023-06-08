package soal5

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func Soal5() {
	var layoutFormat, value string
	var date time.Time
	fmt.Print("Masukkan Waktu Dalam Format 12 Jam (HH:MM:SS AM/PM) : ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		value = scanner.Text()
	}
	layoutFormat = "15:04:05 PM"
	date, _ = time.Parse(layoutFormat, value)
	waktu := date.Format(layoutFormat)
	if date.Format(value) == date.Format(waktu) {
		fmt.Println("Waktu Dalam Format 24 Jam", "\t->", waktu)
	} else {
		fmt.Println(value, layoutFormat)
		fmt.Println("Invalid")
	}

}
