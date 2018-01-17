package bi

import (
	"testing"
	"log"
	"os"
)

//
//func TestBIConfig_Connect(t *testing.T) {
//	var b BIConfig
//	b.Config.URI = "http://192.168.130.44:28080/bi"
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
	b.Config.URI = "http://192.168.130.44:28080/bi"
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
	b.Config.URI = "http://192.168.130.44:28080/bi"
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
	b.Config.URI = "http://192.168.130.44:28080/bi"
	b.Config.Adminv = "admin"
	b.Config.AdminPassv = "g5"
	b.Config.DefaultPassv = "qingcloud123"
	b.Logger = log.New(os.Stdout, "test ", log.Lshortfile|log.Ltime)
	b.User.UserName = "cwc-demo"

	//b.Redirect()

}

func TestBIConfig_GetUserInfo(t *testing.T) {
	var b BIConfig
	b.Config.URI = "http://192.168.130.44:28080/bi"
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
