Go Server
==========

A koa like go server


## Getting Started

```go
package main

import (
    Koa    "koa"
    Store  "store"
)

func main() {
    router := Koa.NewRouter()
    store := Store.New()

    router.Get("/api/users/:userId", func(ctx Koa.Ctx) {
        userId := ctx.requestParam("userId")
        user := store.users.findById(userId)
        ctx.send(user)
    }).Post("/api/user", func(ctx Koa.Ctx) {
        userInfo := ctx.requestBody().(Store.UserInfo)
        user := Store.User.New(userInfo)
        store.users.push(user)
        ctx.send(user)
    }).Register(
        []Koa.Method{Koa.OPTIONS, Koa.GET},
        "/api/user-groups/:userGroupId/users/:userId/friends",
        func(ctx Koa.Ctx) {
            userGroupId, userId := ctx.requestParam("userGroupId", "userId")
            user := store.users.findByUserGroupIdAndUserId(userGroupId, userId)
            friends := store.users.findByUserIds(user.friendUserIds)
            ctx.send(friends)
        },
    )

    koa := Koa.New()
    koa.Use(router.Routes())
    koa.Listen(":8080")
}
```

