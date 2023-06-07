package soal3

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApi(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/posts", nil)
	recorder := httptest.NewRecorder()

	respose := recorder.Result()
	body, _ := io.ReadAll(respose.Body)
	fmt.Println(string(body))
	jsonData := json.NewDecoder(request.Body)

	jsonData.Decode(request)
	fmt.Println(request)
}
