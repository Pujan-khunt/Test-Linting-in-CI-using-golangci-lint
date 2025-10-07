package user

import (
	"testing"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		description string
		name        string
		email       string
		age         int
		wantErr     bool
	}{
		{
			description: "valid user",
			name:        "Alice",
			email:       "alice@example.com",
			age:         25,
			wantErr:     false,
		},
		{
			description: "too young",
			name:        "Bob",
			email:       "bob@example.com",
			age:         15,
			wantErr:     true,
		},
		{
			description: "name too short",
			name:        "Al",
			email:       "al@example.com",
			age:         20,
			wantErr:     true,
		},
		{
			description: "invalid email",
			name:        "Charlie",
			email:       "charlieexample.com",
			age:         30,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.description, func(t *testing.T) {
			user, err := CreateUser(tt.name, tt.email, tt.age)

			if (err != nil) != tt.wantErr {
				t.Errorf("got error = %v, wantErr = %v", err, tt.wantErr)
			}

			if !tt.wantErr {
				if user.name != tt.name {
					t.Errorf("expected name %s, got %s", tt.name, user.name)
				}
				if user.email != tt.email {
					t.Errorf("expected email %s, got %s", tt.email, user.email)
				}
				if user.age != tt.age {
					t.Errorf("expected age %d, got %d", tt.age, user.age)
				}
				if user.id == 0 {
					t.Errorf("expected non-zero id, got %d", user.id)
				}
			}
		})
	}
}
