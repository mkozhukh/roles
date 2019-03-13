package roles

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
