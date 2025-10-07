// Package user
package user

import (
	"errors"
	"strings"
	"sync/atomic"
)

type User struct {
	id    int
	name  string
	email string
	age   int
}

var lastUsedID int32 // default value is 0

func CreateUser(name, email string, age int) (*User, error) {
	if age < 18 {
		return nil, errors.New("age cannot be less than 18")
	}
	if len(name) < 3 {
		return nil, errors.New("name should be at least 3 characters")
	}
	if !strings.Contains(email, "@") {
		return nil, errors.New("invalid email address")
	}

	user := &User{
		id:    getLastUsedID(),
		name:  name,
		email: email,
		age:   age,
	}

	return user, nil
}

func getLastUsedID() int {
	return int(atomic.AddInt32(&lastUsedID, 1))
}
