package app

import "fmt"

type User struct {
	ID   int64
	Name string
}

func (u User) String() string {
	return fmt.Sprintf("ID: %d Name: %s", u.ID, u.Name)
}
