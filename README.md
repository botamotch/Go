# Go Practice

Go言語の練習

# 公式ドキュメント等

- [The Go Programming Language](https://go.dev/)
- [A Tour of Go](https://go.dev/tour/welcome/1)
- [Documentation - The Go Programming Language](https://go.dev/doc/)

# コマンドまとめ

```
$ go mod init xxx  # プロジェクトの初期化
$ go get xxx       # プロジェクト内のモジュールのインストール（go.modを編集）
$ go mod tidy      # ソースコードと比較してモジュールを整理する
$ go mod download  # ???
$ go run xxx.go    # プログラムの実行

$ go install xxx   # ツールをグローバルでインストール
```

# Docker

- [golang - Official Image | Docker Hub](https://hub.docker.com/_/golang)

```Dockerfile
FROM golang:1.19

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./...

CMD ["app"]
```

```
$ docker build -t my-golang-app .
$ docker run -it --rm --name my-running-app my-golang-app
```
