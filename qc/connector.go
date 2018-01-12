package qc

import "net/http"

type Connector interface{
	Connect()
	Close()
	IsUserExist(User)bool
	CreateUser(User)error
	Authorization(User)error
	Redirect(User)http.Handler
}

