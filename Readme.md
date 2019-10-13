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
    router := Koa.Router.New()
    store := Store.New()

    router.Get("/api/users/:userId", func(ctx Koa.Ctx) Store.User {
        userId := ctx.requestParam("userId")
        user := store.users.findById(userId)
        return user
    }).Post("/api/user", func(ctx Koa.Ctx) Store.User {
        userInfo := ctx.requestBody().(Store.UserInfo)
        user := Store.User.New(userInfo)
        store.users.push(user)
        return user
    }).register(
        []Koa.Method{Koa.OPTIONS, Koa.GET},
        "/api/user-groups/:userGroupId/users/:userId/friends",
        func(ctx Koa.Ctx) []Store.User {
            userGroupId, userId := ctx.requestParam("userGroupId", "userId")
            user := store.users.findByUserGroupIdAndUserId(userGroupId, userId)
            friends := store.users.findByUserIds(user.friendUserIds)
            return friends
        },
    )

    koa := Koa.New()
    koa.Use(router.Routes())
    koa.Listen(":8080")
}
```

