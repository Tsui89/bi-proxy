package connector

import (
	"net/http"

	"github.com/Tsui89/bi-proxy/user"
)

type Connector interface {
	Connect() error
	Close() error
	IsConnected() bool
	Refresh() error
	//SetUser(message json.RawMessage) error
	IsUserExist(usr user.User) bool
	CreateUser(usr user.User) error
	Authorization() error
	Redirect(w http.ResponseWriter, r *http.Request, usr user.User)
}
