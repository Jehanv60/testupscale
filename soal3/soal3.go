package soal3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type data struct {
	Id     int    `json:"id"`
	Tittle string `json:"title"`
	Body   string `json:"body"`
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

	var data1 []data
	json.Unmarshal(body, &data1)
	for _, v := range data1 {
		fmt.Println(v.Id)
		fmt.Println(v.Tittle)
		fmt.Println(v.Body)

	}

}
