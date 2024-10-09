package authentication

import (
	"testing"
)

func TestUser_hasGroup(t *testing.T) {
	type fields struct {
		ID     string
		Groups []string
		Roles  []string
	}
	type args struct {
		group string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "ok",
			fields: fields{
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
			fields: fields{
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
			fields: fields{
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
				ID:     tt.fields.ID,
				Groups: tt.fields.Groups,
				Roles:  tt.fields.Roles,
			}
			if got := u.hasGroup(tt.args.group); got != tt.want {
				t.Errorf("User.hasGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_hasRole(t *testing.T) {
	type fields struct {
		ID     string
		Groups []string
		Roles  []string
	}
	type args struct {
		role string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "ok",
			fields: fields{
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
			fields: fields{
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
			fields: fields{
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
				ID:     tt.fields.ID,
				Groups: tt.fields.Groups,
				Roles:  tt.fields.Roles,
			}
			if got := u.hasRole(tt.args.role); got != tt.want {
				t.Errorf("User.hasRole() = %v, want %v", got, tt.want)
			}
		})
	}
}
