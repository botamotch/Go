## [Developer Guide | AWS SDK for Go V2](https://aws.github.io/aws-sdk-go-v2/docs/)

- SDKのインストール、アクセスキーの取得、リージョンの設定などを行う

### Installation
```
$ mkdir aws_s3
$ cd aws_s3/
$ go mod init aws_s3

$ go get github.com/aws/aws-sdk-go-v2
$ go get github.com/aws/aws-sdk-go-v2/config
$ go get github.com/aws/aws-sdk-go-v2/service/s3
$ go get github.com/joho/godotenv

$ vim server.go
$ cat << EOF > .env
> AWS_ACCESS_KEY_ID=...
> AWS_SECRET_ACCESS_KEY=...
> AWS_REGION=us-east-1
> EOF
$ go run server.go
```

