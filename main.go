package main

import (
	"net/http"
	"log"
	"github.com/Tsui89/bi-proxy/qc"
	"fmt"
)



func main() {

	p,err:=qc.NewProxy("./config.yaml")
	if err !=nil{
		fmt.Println(err.Error())
		return
	}
	defer p.Close()

	http.HandleFunc("/", p.ProxyServer)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
