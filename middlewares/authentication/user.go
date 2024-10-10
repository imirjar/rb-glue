package authentication

type UserParams struct {
	ID     string   `json:"id,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Roles  []string `json:"roles,omitempty"`
}

func (u *UserParams) hasGroup(group string) bool {
	for _, g := range u.Groups {
		if g == group {
			return true
		}
	}
	return false
}

func (u *UserParams) hasRole(role string) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}
