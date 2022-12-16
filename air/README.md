### Setup project
```
$ mkdir air
$ cd air
$ go mod init air
go: creating new go.mod: module air
go: to add module requirements and sums:
        go mod tidy
$ go mod tidy
$ go get github.com/labstack/echo/v4
...
```

### Global install - air
```
$ go install github.com/cosmtrek/air@latest
...
```

### Run App
```
$ air
  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ , built with Go

watching .
!exclude tmp
building...
running...

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.9.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```
