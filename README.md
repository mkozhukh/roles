Roles and Rights
=================

Package implements basic authorization system

```
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

auth.RegisterRight(CanEditData)
auth.RegisterRight(CanDeleteData)
auth.RegisterRight(CanManageUsers)

auth.RegisterRole(Editor, CanEditData, CanDeleteData)
auth.RegisterRole(Admin, CanManageUsers)
```

Now you can validate actions like
```
// get check result
var test bool = auth.Check(Editor, CanManageUser)

// or panic if access denied
auth.Guard(Editor, CanManageUser)
```

To make the api more convenient, you can create Role objects like
```
user := auth.NewUser(Admin, Editor)

test := user.Check(CanManageUsers)
user.Guard(CanEditData, CanDeleteData)
```

