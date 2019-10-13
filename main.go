package main

import (
	"koa"
)

func main() {
	koa := Koa.New()
	router := Router.New()
	store := daos.New()

	router.Get("/api/users/:userId", func(ctx Koa.Ctx) daos.User {
		userId := ctx.requestParam("userId")
		user := store.users.findById(userId)
		return user
	}).Post("/api/user", func(ctx Koa.Ctx) daos.User {
		userInfo := ctx.requestBody().(daos.UserInfo)
		user := daos.NewUser(userInfo)
		store.users.push(user)
		return user
	}).register(
		[]daos.Method{daos.OPTIONS, daos.GET},
		"/api/user-groups/:userGroupId/users/:userId/friends",
		func(ctx Koa.Ctx) []daos.User {
			userGroupId, userId := ctx.requestParam("userGroupId", "userId")
			user := store.users.findByUserGroupIdAndUserId(userGroupId, userId)
			friends := store.users.findByUserIds(user.friendUserIds)
			return friends
		},
	)

	koa.use(router.routes())
	koa.listen(":8080")
}
