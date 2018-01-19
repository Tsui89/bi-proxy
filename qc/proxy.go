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
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/Tsui89/bi-proxy/user"
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
	if !strings.HasSuffix(p.QYConfig.ApiUri,"/"){
		p.QYConfig.ApiUri = p.QYConfig.ApiUri+"/"
	}

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

	qcuser, err := p.GetUser(auth.Payload, auth.Signature)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("get user error: " + err.Error()))
		return
	}
	//userRaw, _ := json.Marshal(user)
	if p.conn.IsUserExist(qcuser) == false {
		err := p.conn.CreateUser(qcuser)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("create user error: " + err.Error()))
			return
		}
	}
	p.conn.Redirect(w, req,qcuser)
	return
}

func (p Proxy) GetUser(payload, signature []byte) (user.User, error) {
	var auth QCAuth
	var qcInfo QCinfo
	var rb user.DescribeUsersResp
	var qcuser user.User
	//user.BiUri = p.BIConfig.BiUri
	auth.Payload, auth.Signature = payload, signature
	auth.SecretKey = []byte(p.QYConfig.SecretAccessKey)

	qcInfo = auth.ExtractPayload()
	qcuser.UserId = qcInfo.UserId

	method := "GET"
	//endPoint := p.QYConfig.Host
	//uri := p.QYConfig.URI
	ru,_ := url.Parse(p.QYConfig.ApiUri)

	//ru.Path = path.Join(ru.RawPath +p.QYConfig.URI) +"/"
	rq:= ru.Query()
	rq.Add(
		"users.1",
		fmt.Sprint(qcInfo.UserId),
	)
	rq.Add(
		"action",
		fmt.Sprint("DescribeUsers"),
	)
	rq.Add(
		"time_stamp",
		fmt.Sprint(qcInfo.TimeStamp),
	)

	rq.Add(
		"access_key_id",
		fmt.Sprint(p.QYConfig.AccessKeyID),
	)

	rq.Add(
		"version",
		fmt.Sprint(1),
	)
	rq.Add(
		"signature_method",
		"HmacSHA256",
	)
	rq.Add(
		"signature_version",
		fmt.Sprint(1),
	)
	rq.Add(
		"access_token",
		fmt.Sprint(qcInfo.AccessToken),)
	rq.Add(
		"app_id",
		p.QYConfig.AppId,
	)

	sc := strings.Join([]string{method, ru.Path, rq.Encode()}, "\n")
	//fmt.Println(urlStr=="GET\n/iaas/\naccess_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1")
	hash := hmac.New(sha256.New, []byte(p.QYConfig.SecretAccessKey))
	hash.Write([]byte(sc))
	base64Payload := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	index := strings.LastIndex(base64Payload, "=")
	r := strings.Replace(base64Payload[:index+1], " ", "+", -1)
	ru.RawQuery = rq.Encode()+"&"+"signature="+url.QueryEscape(r)

	urlStr := ru.String()
	p.logger.Println(urlStr)
	resp, err := http.Get(urlStr)

	if err != nil {
		p.logger.Println(err.Error())
		return qcuser, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.logger.Println(err.Error())
		return qcuser, err
	}
	json.Unmarshal(body, &rb)
	if rb.RetCode == 0 && len(rb.UserSet) > 0 {
		json.Unmarshal(rb.UserSet[0], &qcuser)
	}
	p.logger.Println(qcuser)
	return qcuser, nil
}
