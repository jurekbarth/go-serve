package main

import "net/http"

func main() {
	http.ListenAndServe(":1234", nil)
}
