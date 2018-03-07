package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tsui89/bi-proxy/user"
)

/*
	Connect() error
	Close() error
	IsConnected() bool
	Refresh() error
	//SetUser(message json.RawMessage) error
	IsUserExist(usr user.User) bool
	CreateUser(usr user.User) error
	Authorization() error
	Redirect(w http.ResponseWriter, r *http.Request,usr user.User)
*/

type App struct {
	Name        string
	RedirectUri string `json:"redirect_uri"`
}

func NewApp(config json.RawMessage, logger *log.Logger) *App {
	var app App
	return &app
}

func (app *App) Connect() error {
	return nil
}

func (app *App) Close() error {
	return nil
}

func (app *App) Refresh() error {
	return nil
}

func (app *App) IsConnected() bool {
	return true
}

func (app *App) IsUserExist(u user.User) bool {
	return true
}

func (app *App) CreateUser(u user.User) error {
	return nil
}

func (app *App) Authorization() error {
	return nil
}

func (app *App) Redirect(w http.ResponseWriter, r *http.Request, usr user.User) {
	urlStr := app.RedirectUri
	http.Redirect(w, r, urlStr, http.StatusFound)
	return
}
