package enum

import (
	"encoding/json"
	"fmt"
)

type Role int

const (
	User Role = iota
	Admin
)

// Role -> String dönüşümü
func (r Role) String() string {
	switch r {
	case User:
		return "user"
	case Admin:
		return "admin"
	default:
		return "unknown"
	}
}

// JSON için Marshal ve Unmarshal metodları

// Role -> JSON (String)
func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

// JSON (String) -> Role
func (r *Role) UnmarshalJSON(data []byte) error {
	var roleStr string
	if err := json.Unmarshal(data, &roleStr); err != nil {
		return err
	}

	role, err := ParseRole(roleStr)
	if err != nil {
		return err
	}

	*r = role
	return nil
}

// String'ten Role enum'una çevirme fonksiyonu
func ParseRole(s string) (Role, error) {
	switch s {
	case "user":
		return User, nil
	case "admin":
		return Admin, nil
	default:
		return -1, fmt.Errorf("invalid rol: %s", s)
	}
}
