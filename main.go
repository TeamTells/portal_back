package main

import (
	"log"
	"net/http"
	"portal_back/authentication/impl/di"
)

func main() {
	di.InitAuthModule()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
