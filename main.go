package main

import (
	"net/http"
	"log"

)



func main() {

	p:=NewProxy("./config.yaml")

	http.HandleFunc("/", p.ProxyServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
