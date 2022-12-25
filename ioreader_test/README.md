# `io.Reader` について

**バイト列を読み出すためのインターフェース**のこと。

標準ライブラリの中で下記の通り定義されており、`Read([]byte) (int, error)`というシグネチャを持つあらゆる型が該当する。

```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

関数を`io.Reader`を受け取るようにすればあらゆる型の変数を適用することができて非常に強力。

らしいがいまいちピンとこない。下記のサイトで強くおすすめされてるのでまた今度読み返してみる。

- https://qiita.com/ktnyt/items/8ede94469ba8b1399b12

---

例）ファイルの中身を`os.File`から読み出してバイト列に格納する
```go
import "os"

m := 14
f, _ := os.Open("memo.txt")
p := make([]byte, m)
n, err := f.Read(p)

fmt.Println(string(p))
```

出力
```
IO.Reader test
```

---

例）文字列を`strings.Reader`に変換して`http.NewRequest`のbodyに入れてPOSTリクエストを送信する

```go
import "net/http"

req, _ := http.NewRequest(
	http.MethodPost,
	"http://localhost:8080/api/v1/admin/works",
	strings.NewReader(`{ "name": "テスト用作品タイトル10001", "isReleased": true }`),
)

client := &http.Client{Timeout: time.Duration(30) * time.Second}
res, err := client.Do(req)
```
