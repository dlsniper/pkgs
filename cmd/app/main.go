package main

import (
	"fmt"

	"github.com/dlsniper/pkgs/db"
	"github.com/dlsniper/pkgs/user"
)

func main() {
	u := &user.User{}
	fmt.Printf("exists: %v\n", db.UserExistsByEmail(u))
}
