package main

import (
	"net/http"

	"github.com/sudhanshu-chauhan/gopad/routes"
)

func main() {
	indexRoute := routes.GetIndexRoutes()
	http.ListenAndServe("0.0.0.0:8000", &indexRoute)
}
