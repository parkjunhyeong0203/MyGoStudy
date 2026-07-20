package main

import (
	"net/http"

	"web4/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
