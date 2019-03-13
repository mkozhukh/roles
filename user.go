package guard

import "context"

// User represents a single user
type User struct {
	ID       uint
	Registry *Registry `json:"-"`

	roles []Role
}

// Check returns true if user has one of provided rights
func (u *User) Check(right ...Right) bool {
	if u.ID == 0 {
		return false
	}

	for i := range u.roles {
		if u.Registry.Check(u.roles[i], right...) {
			return true
		}
	}

	return false
}

// Guard stops process if user has not any of provided rights
func (u *User) Guard(right ...Right) {
	if u.ID == 0 {
		panic("Unknown user")
	}

	for i := range u.roles {
		if u.Registry.Check(u.roles[i], right...) {
			return
		}
	}

	panic("Access denied")
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
