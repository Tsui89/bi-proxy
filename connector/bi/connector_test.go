package bi

import (
	"testing"
	"log"
	"os"
	"net/url"
	"strings"
)

//
//func TestBIConfig_Connect(t *testing.T) {
//	var b BIConfig
//	b.Config.APIURI = "http://192.168.130.44:28080/bi"
//	b.Config.Admin = "admin"
//	b.Config.Password = "g5"
//	b.Logger = log.New(os.Stdout,"test ",log.Lshortfile|log.Ltime)
//
//	err:=b.Connect()
//	if err !=nil{
//		t.Fatal()
//	}
//	err=b.Close()
//	if err !=nil{
//		t.Fatal()
//	}
//}
func TestBIConfig_IsUserExist(t *testing.T) {
	var b BIConfig
	b.Config.APIURI = "http://192.168.130.44:28080/bi"
	b.Config.Adminv = "admin"
	b.Config.AdminPassv = "g5"
	b.Config.DefaultGroup = "qingcloud"
	b.Logger = log.New(os.Stdout, "test ", log.Lshortfile|log.Ltime)
	b.User.UserId = "usr-ibGnrli6"
	b.User.UserName = "ldapuser1"
	err := b.Connect()
	if err != nil {
		t.Fatal()
	}

	if (b.IsUserExist() == false) {
		t.Fatal()
	}
	err = b.Close()
	if err != nil {
		t.Fatal()
	}
}

func TestBIConfig_CreateUser(t *testing.T) {
	var b BIConfig
	b.Config.APIURI = "http://192.168.130.44:28080/bi"
	b.Config.Adminv = "admin"
	b.Config.AdminPassv = "g5"
	b.Config.DefaultPassv = "qingcloud123"
	b.Config.DefaultGroup = "qingcloud"
	b.Logger = log.New(os.Stdout, "test ", log.Lshortfile|log.Ltime)
	b.User.UserId = "usr-ibGnrli6"
	b.User.UserName = "ldapuser1"
	err := b.Connect()
	if err != nil {
		t.Fatal()
	}
	if (b.IsUserExist() == false) {
		err = b.CreateUser()
		if err != nil {
			t.Fatal()
		}
	}
	err = b.Close()
	if err != nil {
		t.Fatal()
	}
}

func TestBIConfig_Redirect(t *testing.T) {
	var b BIConfig
	b.Config.APIURI = "http://192.168.130.44:28080/bi"
	b.Config.Adminv = "admin"
	b.Config.AdminPassv = "g5"
	b.Config.DefaultPassv = "qingcloud123"
	b.Config.RedirectUri = "http://192.168.130.44:28080/bi/Viewer?proc=1&action=viewer&db=%E9%A3%8E%E6%9C%BA%E5%81%A5%E5%BA%B7%E8%AF%84%E4%BC%B0%E5%8A%A9%E6%89%8B/%E9%A3%8E%E6%9C%BA%E5%BA%94%E7%94%A8%E5%8F%AF%E8%A7%86%E5%8C%96%E8%AE%BE%E8%AE%A1V2"
	b.Logger = log.New(os.Stdout, "test ", log.Lshortfile|log.Ltime)
	b.User.UserName = "cwc-demo"

	ru,_ := url.Parse(b.Config.RedirectUri)
	rq,_ := url.ParseQuery(ru.RawQuery)
	t.Log(ru.User)
	t.Log(ru.Host)
	t.Log(ru.Path)
	t.Log(ru.Scheme)
	t.Log(ru.RawQuery)
	t.Log(ru.RawPath)
	t.Log(ru.Query())
	t.Log(ru.RequestURI())
	t.Log(ru.EscapedPath())
	t.Log(ru.String())
	t.Log(rq.Encode())

	rq.Set("actions","login")
    ru.RawQuery = rq.Encode()

	t.Log(rq.Encode())

	t.Log(ru.Query())
	t.Log(ru.String())

}

func TestBIConfig_GetUserInfo(t *testing.T) {
	var b BIConfig
	b.Config.APIURI = "http://192.168.130.44:28080/bi"
	b.Config.Adminv = "admin"
	b.Config.AdminPassv = "g5"
	b.Config.DefaultPassv = "qingcloud123"
	b.Config.DefaultGroup = "qingcloud"
	b.Logger = log.New(os.Stdout, "test ", log.Lshortfile|log.Ltime)
	b.User.UserId = "usr-ibGnrli6"
	b.User.UserName = "ldapuser1"
	err := b.Connect()
	if err != nil {
		t.Fatal()
	}
	_, err = b.GetUserInfo(b.User.UserId, b.Config.DefaultGroup)
	if err != nil {
		t.Fatal()
	}

	err = b.Close()
	if err != nil {
		t.Fatal()
	}
}
