package main

import (
	"golesson/restful"
	"net/http"
)

func main() {
	http.HandleFunc("/get", restful.HandleGet)
	http.HandleFunc("/post", restful.HandlePost)
	http.ListenAndServe(":8080", nil)

}
