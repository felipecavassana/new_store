package main

import (
	"net/http"

	"github.com/felipecavassana/new_store/routes"
)

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
