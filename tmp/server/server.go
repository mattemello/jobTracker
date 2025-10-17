package server

import (
	"fmt"
	"net/http"
)

func HandlerMainPage(w *http.ResponseWriter) {
	fmt.Fprintf(*w, "This is the main page\n")
}
