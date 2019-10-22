package Handler

import (
	"io"
	"net/http"
)

//get handler
func Serves(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "This works")

}
