package db

import "fmt"

type Emailer interface {
	Email() string
}

// given this is starting with user it would be better fitted in the user package
func UserExistsByEmail(u Emailer) bool {
	// do stuff here
	fmt.Printf("exists by email: %v\n", u.Email())
	return true
}
