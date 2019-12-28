Go Server
==========

A koa like go server

### Generate Proto
```shell script
$ cd naive-go-koa
$ $env:GOPATH = $(pwd)
$ go get -u github.com/golang/protobuf/protoc-gen-go
$ protoc --proto_path=src  --go_out=src src/model/*.proto
```

### Getting Started

```go
package main

import (
    Koa    "kao"
)

type User struct {
    id int
    userGroupId int
    Name string
    Age int
    Description string
}

func main() {
    router := Koa.NewRouter()

    router.Get("/api/users/:userId", func(ctx Koa.Ctx) {
        userId := ctx.getParam("userId")
        ctx.send(&User{
            id: userId,
        })
    }).Post("/api/user", func(ctx Koa.Ctx) {
        user := ctx.getBody().(User)
        ctx.send(user)
    }).Get(
        "/api/user-groups/:userGroupId/users/:userId/friends",
        func(ctx Koa.Ctx) {
            userGroupId, userId := ctx.getParam("userGroupId", "userId")
            ctx.send(&User{
                id: userId,
                userGroupId: userGroupId,
}           )
        },
    )

    koa := Koa.New()
    koa.Use(router)
    koa.Listen(":8080")
}
```
