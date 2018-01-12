package main

import (
	"net/http"
	"log"
	"github.com/Tsui89/bi-proxy/qc"
)



func main() {

	p:=qc.NewProxy("./config.yaml")

	http.HandleFunc("/", p.ProxyServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
