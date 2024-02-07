package main

import (
	"net/http"

	"gopad/routes"
)

func main() {
	indexRoute := routes.GetIndexRoutes()
	http.ListenAndServe("0.0.0.0:8000", &indexRoute)
}
