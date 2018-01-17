package qc

import (
	"net/http"
	"net/url"
	"io/ioutil"
	"log"
	"encoding/json"
	"github.com/ghodss/yaml"
	"os"
	"github.com/Tsui89/bi-proxy/connector/bi"
	"fmt"
)

func NewProxy(fp string) (*Proxy, error) {
	var p Proxy
	var err error

	conf, _ := ioutil.ReadFile(fp)
	//fmt.Println(string(conf))

	yaml.Unmarshal(conf, &p)
	//var logFile *os.File
	logFile := os.Stdout
	if (p.LogFile != "") {
		logFile, err = os.OpenFile(p.LogFile, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			return nil, err
		}
	}
	p.logger = log.New(logFile, "proxy ", log.Lshortfile|log.Ltime)

	switch p.PConfig.Type {
	case "bi":
		p.conn = bi.NewBi(p.PConfig.Config, p.logger)
		p.conn.Connect()
	}
	p.logger.Println(p)
	return &p, nil
}

func (p Proxy) Close() {
	p.conn.Close()
}

func parseBody(data string) (payload, signature []byte) {
	s, _ := url.ParseQuery(data)
	return []byte(s.Get("payload")), []byte(s.Get("signature"))
}

// hello world, the web server
func (p Proxy) ProxyServer(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Redirect(w, req, req.URL.Path, 304)
		return
	}

	data, _ := ioutil.ReadAll(req.Body)

	var auth QCAuth
	auth.Payload, auth.Signature = parseBody(string(data))
	auth.SecretKey = []byte(p.QYConfig.SecretAccessKey)
	if !auth.Authentication() {
		p.logger.Println("authentication error")
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	}

	p.logger.Println("authentication success")
	qcinfo := auth.ExtractPayload()
	p.logger.Println(qcinfo)
	if p.conn.IsConnected() == false {
		err := p.conn.Refresh()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("refresh error: " + err.Error()))
			return
		}
	}

	user, err := p.GetUser(auth.Payload, auth.Signature)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("get user error: " + err.Error()))
		return
	}
	userRaw, _ := json.Marshal(user)
	p.conn.SetUser(userRaw)
	if p.conn.IsUserExist() == false {
		err := p.conn.CreateUser()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("create user error: " + err.Error()))
			return
		}
	}
	p.conn.Redirect(w, req)
	return
}

func (p Proxy) GetUser(payload, signature []byte) (User, error) {
	var auth QCAuth
	var qcInfo QCinfo
	var rb DescribeUsersResp
	var user User
	//user.BiUri = p.BIConfig.BiUri
	auth.Payload, auth.Signature = payload, signature
	auth.SecretKey = []byte(p.QYConfig.SecretAccessKey)

	//if !auth.Authentication() {
	//	return &user, errors.New("authentication error.")
	//}

	qcInfo = auth.ExtractPayload()
	user.UserId = qcInfo.UserId

	ps := NewParameters()
	ps.Add(&(Parameter{
		"users.1",
		qcInfo.UserId,
	}))
	ps.Add(&Parameter{
		"action",
		"DescribeUsers",
	})
	ps.Add(&Parameter{
		"time_stamp",
		qcInfo.TimeStamp,
	})

	ps.Add(&Parameter{
		"access_key_id",
		p.QYConfig.AccessKeyID,
	})

	ps.Add(&Parameter{
		"version",
		1,
	})
	ps.Add(&Parameter{
		"signature_method",
		"HmacSHA256",
	})
	ps.Add(&Parameter{
		"signature_version",
		1,
	})
	ps.Add(&Parameter{
		"access_token",
		qcInfo.AccessToken,
	})
	ps.Add(&Parameter{
		"app_id",
		p.QYConfig.AppId,
	})
	method := "GET"
	endPoint := p.QYConfig.Protocol + "://" + p.QYConfig.Host + ":" + fmt.Sprintf("%d", p.QYConfig.Port)
	uri := p.QYConfig.URI
	secretKey := p.QYConfig.SecretAccessKey
	p.logger.Println(method, endPoint, uri, secretKey)
	urlStr := ps.GenerateReqStr(method, endPoint, uri, secretKey)
	p.logger.Println(urlStr)
	resp, err := http.Get(urlStr)

	if err != nil {
		p.logger.Println(err.Error())
		return user, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.logger.Println(err.Error())
		return user, err
	}
	json.Unmarshal(body, &rb)
	if rb.RetCode == 0 && len(rb.UserSet) > 0 {
		json.Unmarshal(rb.UserSet[0], &user)
	}
	p.logger.Println(user)
	return user, nil
}
