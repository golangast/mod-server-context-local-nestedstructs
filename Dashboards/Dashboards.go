package Dashboard

import (
	"fmt"

	Pages "github.com/golangast/Dashboard/Page"
	Users "github.com/golangast/Dashboard/User"
)

//Dashboards controls this.
type Dashboards struct {
	ID   int
	User Users.User
	Page Pages.Page
}

func CreateDashboards() (Dashboards, error) {
	var err error
	fmt.Println("this is dashboard")
	d := Dashboards{ID: 1, User: Users.User{Name: "jane", Email: "jane@yahoo.com"}, Page: Pages.Page{ID: 33}}
	return d, err
}
