package soal4

import (
	"fmt"
	"math/rand"
)

func Soal4(n int) []int {
	resul := 0
	var awal []int
	for i := 0; i < n; i++ {
		resul += 1
	}
	if n > 0 {
		awal = Soal4(n - 1)
		awal = append(awal, resul)
		fmt.Println(awal)
	}
	rand.Shuffle(len(awal), func(i, j int) {
		awal[i], awal[j] = awal[j], awal[i]
	})
	return awal
}
