package qc

import (
	"encoding/json"
	"net/http"
)

type Connector interface {
	Connect() error
	Close() error
	IsConnected() bool
	Refresh() error
	SetUser(message json.RawMessage) error
	IsUserExist() bool
	CreateUser() error
	Authorization() error
	Redirect(w http.ResponseWriter, r *http.Request)
}
