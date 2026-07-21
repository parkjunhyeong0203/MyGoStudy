package main

import (
	"net/http"
	"web7/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
