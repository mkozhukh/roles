package roles

import (
	"context"
	"strconv"
	"strings"
)

type rolekeyType string

// UserFromContext gets current user from context
func UserFromContext(ctx context.Context) *User {
	user, ok := ctx.Value(rolekeyType("user")).(*User)
	if !ok {
		return &User{}
	}

	return user
}

// UserToContext stores user info in a context
func UserToContext(ctx context.Context, user *User) context.Context {
	return context.WithValue(ctx, rolekeyType("user"), user)
}

// ParseRights converts a string to a list of rights
func ParseRights(name string) ([]Right, error) {
	chunks := strings.Split(name, ",")
	rights := make([]Right, 0)

	var err error
	for i := range chunks {
		num, _ := strconv.Atoi(chunks[i])
		rights = append(rights, Right(num))
	}

	return rights, err
}

// SerializeRights converts a list of rights to a string
func SerializeRights(rights ...Right) string {
	if len(rights) == 0 {
		return ""
	}

	strs := make([]string, 0, len(rights))
	for i := range rights {
		code := strconv.Itoa(int(rights[i]))
		strs = append(strs, code)
	}

	return strings.Join(strs, ",")
}
