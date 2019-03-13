Roles and Rights
=================

Package implements basic authorization system

```go
const (
        CanEditData role.Right = iota
        CanDeleteData
        CanManageUsers
)

const (
		None role.Role = iota
        Editor 
        Admin
)


auth := role.NewRegistry();

auth.RegisterRole(Editor, CanEditData, CanDeleteData)
auth.RegisterRole(Admin, CanManageUsers)
```

Now you can validate access like next

```go
user := auth.NewUser(Admin, Editor)

test := user.Check(CanManageUsers)
user.Guard(CanEditData, CanDeleteData)
```

Also, you can use roles checking as part of routing

```go
router.Use(auth.GuardRequest("/denied", CanEdit))
```

combined with mkozhukh/remote it can be used to block access to API methods

```go
remote.RegisterWithGuard("api", AdminApi, auth.CheckRequest(CanAdminUsers))
```

Both CheckRequest and GuardRequest expect to find User object in the context
Something like next is expected

```go
router.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
                user := getUserInfo(r); //some custom API to get user info

                ctx := role.UserToContext(r.Context(), role.NewUser(user.ID, user.Roles))
                next.ServeHTTP(w, r.WithContext(ctx))
        })
})
```