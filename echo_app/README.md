# [Guide | echo - High performance, minimalist Go web framework](https://echo.labstack.com/guide/)

## Quick Start

### Installation, Hello World
```
$ mkdir echo_app
$ cd echo_app/
$ go mod init echo_app
go: creating new go.mod: module echo_app
$ ll
Permissions Size User     Date Modified Name
.rw-r--r--    25 bmkazuya  1  8月 14:03 go.mod
$ go get github.com/labstack/echo/v4
go: downloading github.com/labstack/echo v3.3.10+incompatible
go: downloading golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
go: downloading golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f
go: downloading golang.org/x/sys v0.0.0-20211103235746-7861aae1554b
go: added github.com/labstack/echo/v4 v4.7.2
go: added github.com/labstack/gommon v0.3.1
go: added github.com/mattn/go-colorable v0.1.11
go: added github.com/mattn/go-isatty v0.0.14
go: added github.com/valyala/bytebufferpool v1.0.0
go: added github.com/valyala/fasttemplate v1.2.1
go: added golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
go: added golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f
go: added golang.org/x/sys v0.0.0-20211103235746-7861aae1554b
go: added golang.org/x/text v0.3.7
$ vim server.go

$ go run server.go

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.7.2
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:1323
```

then access to [localhost:1323](http://localhost:1323)

### Path Parameters

```go
func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
```
[localhost:1323/users/1](http://localhost:1323/users/1)

### Query Parameters
```go
func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
```

[localhost:1323/show?team=x-mem&member=wolverine](http://localhost:1323/show?team=x-mem&member=wolverine)

### application/x-www-form-urlencoded
```go
func save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
```

### Handling Request
???
```go
e.POST("/users/", func(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
})
e.Logger.Fatal(e.Start(":1323"))
```

### Static Content
SKIP

### Template Rendering
SKIP

### Middleware
???

