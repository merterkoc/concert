package enum

type Role string

func (r Role) String() string {
	return string(r)
}

const (
	User  Role = "user"
	Admin Role = "admin"
)
