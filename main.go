package main

import (
	"fmt"
	"log"

	Dashboards "github.com/golangast/Dashboard/Dashboards"
	Pages "github.com/golangast/Dashboard/Page"
	Servs "github.com/golangast/Dashboard/Server"
	Users "github.com/golangast/Dashboard/User"
)

func main() {

	fmt.Println("....starting")
	Servs.Serv()
	P, err := Pages.CreatePage()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(P)

	U, err := Users.CreateUser()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(U)

	D, err := Dashboards.CreateDashboards()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(D)
}
