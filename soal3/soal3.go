package soal3

import (
	"fmt"
	"io"
	"net/http"
)

type Data struct {
	id     int    `json:"id"`
	tittle string `json:"tittle"`
	body   string `json:"body"`
}

func Soal3() {
	url := "https://jsonplaceholder.typicode.com/posts"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(string(body))

}
