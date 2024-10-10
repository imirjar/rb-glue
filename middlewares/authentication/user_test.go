package authentication

import (
	"testing"
)

func TestUser_hasGroup(t *testing.T) {

	type args struct {
		group string
	}
	tests := []struct {
		name string
		user User
		args args
		want bool
	}{
		{
			name: "ok",
			user: User{
				ID:     "1",
				Groups: []string{"admin"},
			},
			args: args{
				group: "admin",
			},
			want: true,
		},
		{
			name: "no groups",
			user: User{
				ID:     "1",
				Groups: []string{},
			},
			args: args{
				group: "admin",
			},
			want: false,
		},
		{
			name: "other groups",
			user: User{
				ID:     "1",
				Groups: []string{"manager"},
			},
			args: args{
				group: "admin",
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:     tt.user.ID,
				Groups: tt.user.Groups,
				Roles:  tt.user.Roles,
			}
			if got := u.hasGroup(tt.args.group); got != tt.want {
				t.Errorf("User.hasGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_hasRole(t *testing.T) {

	type args struct {
		role string
	}
	tests := []struct {
		name string
		user User
		args args
		want bool
	}{
		{
			name: "ok",
			user: User{
				ID:    "1",
				Roles: []string{"admin"},
			},
			args: args{
				role: "admin",
			},
			want: true,
		},
		{
			name: "no roles",
			user: User{
				ID:    "1",
				Roles: []string{},
			},
			args: args{
				role: "admin",
			},
			want: false,
		},
		{
			name: "other roles",
			user: User{
				ID:    "1",
				Roles: []string{"manager"},
			},
			args: args{
				role: "admin",
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:     tt.user.ID,
				Groups: tt.user.Groups,
				Roles:  tt.user.Roles,
			}
			if got := u.hasRole(tt.args.role); got != tt.want {
				t.Errorf("User.hasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
