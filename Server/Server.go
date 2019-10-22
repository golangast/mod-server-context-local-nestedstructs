package Server

import (
	"context"
	"log"
	"net/http"

	Context "github.com/golangast/Dashboard/Context"
	Handlers "github.com/golangast/Dashboard/Handler"
)

//starts the server
func Serv() {

	mux := http.NewServeMux()
	cc := context.Background()
	//handlers

	mux.HandleFunc("/servers", Handlers.Serves)

	//context handlers
	contextedMux := Context.AddContext(cc, mux)
	log.Fatal(http.ListenAndServe(":8080", contextedMux))

}
