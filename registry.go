package guard

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
)

// Registry stores links between roles and rights
type Registry struct {
	rights map[Right]Right
	roles  map[Role][]Right
}

// NewRegistry returns new role access registry
func NewRegistry() *Registry {
	t := Registry{}

	t.rights = make(map[Right]Right)
	t.Reset()

	return &t
}

// Router is a http mux
type Router interface {
	Get(pattern string, handlerFn http.HandlerFunc)
}

// Right is an enumeration of possible access levels
type Right uint

// Role represents a set of rights
type Role uint

//RegisterRight allows to define a new right
func (rg *Registry) RegisterRight(right Right) {
	rg.rights[right] = right
}

// Reset clears all known roles
func (rg *Registry) Reset() {
	rg.roles = make(map[Role][]Right)
}

// RegisterRole allows to configure list of rights for the role
func (rg *Registry) RegisterRole(role Role, rights ...Right) {
	rg.roles[role] = rights
}

// NewUser returns new role object
func (rg *Registry) NewUser(id uint, roles ...Role) *User {
	return &User{Registry: rg, ID: id, roles: roles}
}

// ParseRightsString converts a string to a list of rights
func (rg *Registry) ParseRightsString(name string) ([]Right, error) {
	chunks := strings.Split(name, ",")
	rights := make([]Right, 0)

	var err error
	for i := range chunks {
		num, _ := strconv.Atoi(chunks[i])
		right, ok := rg.rights[Right(num)]
		if ok {
			rights = append(rights, right)
		} else {
			err = errors.New("unknown access right")
		}
	}

	return rights, err
}

// Serialize converts a list of rights to a string
func (rg *Registry) Serialize(rights ...Right) string {
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

// GetRights returns all right for the role
func (rg *Registry) GetRights(roles ...Role) []Right {
	result := make([]Right, 0)
	for _, roleID := range roles {
		rights, ok := rg.roles[roleID]
		if ok {
			result = append(result, rights...)
		}
	}

	return result
}

// Check confirms that role has all of provided rights
func (rg *Registry) Check(role Role, types ...Right) bool {
	rights := rg.GetRights(role)
	return intersection(rights, types)
}

// Guard stops execution if role has all of provided rights
func (rg *Registry) Guard(role Role, types ...Right) {
	if !rg.Check(role, types...) {
		panic("Access denied")
	}
}

func intersection(a, b []Right) bool {
	if a == nil || b == nil {
		return false
	}

	for j := range a {
		for i := range b {
			if a[j] == b[i] {
				return true
			}
		}
	}
	return false
}

// CheckRequest helper validates if role behind the request has one of the provided rights
func (rg *Registry) CheckRequest(types ...Right) func(*http.Request) bool {
	return func(r *http.Request) bool {
		return UserFromContext(r.Context()).Check(types...)
	}
}

// GuardRequest middleware, checks role access and redirects to "denied" page when access denied
func (rg *Registry) GuardRequest(redirect string, types ...Right) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ok := UserFromContext(r.Context()).Check(types...)
			if !ok {
				http.Redirect(w, r, redirect, http.StatusTemporaryRedirect)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
