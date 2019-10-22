package Users

import "fmt"

//User for Dashboard.
type User struct {
	Name  string
	Email string
}

//Creates a User.
func CreateUser() (User, error) {
	var err error
	fmt.Println("user created")

	u := User{Name: "jill", Email: "jill@yahoo.com"}

	return u, err
}
