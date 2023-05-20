package users

import "context"

type key int

const (
	userContextKey key = 1
)

func AddUserToContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

func GetUserFromContext(ctx context.Context) User {
	return ctx.Value(userContextKey).(User)
}
