package roles

import (
	"context"
)

// User represents a single user
type User struct {
	ID       uint
	Registry *Registry `json:"-"`

	rights []Right
}

// Check returns true if user has all of provided rights
func (u *User) Check(rights ...Right) bool {
	return intersection(u.rights, rights)
}

// Guard stops process if user has not any of provided rights
func (u *User) Guard(right ...Right) {
	if !u.Check(right...) {
		panic("Access denied")
	}
}

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

func intersection(base, search []Right) bool {
	if search == nil {
		return true
	}
	if base == nil {
		return false
	}

nextrule:
	for j := range search {
		for i := range base {
			if search[j] == base[i] {
				continue nextrule
			}
		}
		return false
	}

	return true
}
