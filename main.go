package main

import (
	"fmt"
	"net/http"

	"github.com/okhrimko/simplemiddleware/middleware"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hey, hello man!<br/>")
}

func main() {

	http.Handle("/", middleware.Default.Then(hello))
	http.ListenAndServe(":3000", nil)
}
