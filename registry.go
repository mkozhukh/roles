package roles

import (
	"net/http"
)

// Registry stores links between roles and rights
type Registry struct {
	roles map[Role][]Right
}

// NewRegistry returns new role access registry
func NewRegistry() *Registry {
	t := Registry{}
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
	return &User{Registry: rg, ID: id, rights: rg.GetRights(roles...)}
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
