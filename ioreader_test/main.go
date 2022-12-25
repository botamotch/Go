package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type PostWorkRequestBody struct {
	Name       string `json:"name"`
	IsReleased bool   `json:"isReleased"`
}

func main() {
	{
		m := 14
		f, _ := os.Open("memo.txt")
		p := make([]byte, m)
		n, err := f.Read(p)
		if m < n { // 3. 返り値処理
			log.Fatalf("%dバイト読もうとしましたが、%dバイトしか読めませんでした\n", n, m)
		}
		if err != nil { // 4. エラー処理
			log.Fatalf("読込中にエラーが発生しました：%v\n", err)
		}

		fmt.Println(string(p))
	}

	{
		url := "http://localhost:8080/api/v1/admin/works"
		body := `{ "name": "テスト用作品タイトル10001", "isReleased": true }`

		req, _ := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
		req.SetBasicAuth("admin", "password")
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{Timeout: time.Duration(30) * time.Second}

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("> Request do failed : ", err.Error())
			return
		}

		got, err := ioutil.ReadAll(res.Body)
		res.Body.Close()

		fmt.Println("Status Code   :", res.StatusCode)
		fmt.Println("Response body :", string(got))
	}
}
