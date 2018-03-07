package app

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

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
type Config struct {
	RedirectUri string `json:"redirect_uri"`
}

type App struct {
	Name   string
	Config Config
	*log.Logger
}

func NewApp(config json.RawMessage, logger *log.Logger) *App {
	var app App
	app.Logger = logger
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

func (app *App) Redirect(w http.ResponseWriter, r *http.Request, u user.User) {
	app.Logger.Println("redirect to app")
	ru, err := url.Parse(app.Config.RedirectUri)
	if err != nil {
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			return
		}
	}
	rq := ru.Query()
	//rq.Set("au_act",ActionLogin)	ru.RawQuery = rq.Encode()
	//ru.RawQuery = strings.Join([]string{ru.RawQuery,paramStr},"&")

	switch u.UserId {
	case "xxx":
		rq.Set("user", "user-xxx")
		rq.Set("password", "password")
	case "yyy":
		rq.Set("user", "user-yyy")
		rq.Set("password", "password")
	default:
		w.WriteHeader(500)
		w.Write([]byte("no such user"))
		return
	}
	urlStr := ru.String()
	app.Logger.Panicln("redirect Url to: ", urlStr)
	http.Redirect(w, r, urlStr, http.StatusFound)
	return
}
