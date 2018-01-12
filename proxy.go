package main

import (
	"fmt"
	"net/http"
	"net/url"
	"io/ioutil"
	"io"
	"log"
	"encoding/json"
	"github.com/ghodss/yaml"
	"os"
)

type BiConfig struct{
	BiUri string	`json:"bi_uri"`
}


type Config struct {
	AccessKeyID     string `json:"access_key_id"`
	SecretAccessKey string `json:"secret_access_key"`
	Host              string `json:"host"`
	Port              int    `json:"port"`
	Protocol          string `json:"protocol"`
	URI               string `json:"uri"`
}

type Proxy struct{
	QYConfig Config `json:"qy_config"`
	BIConfig BiConfig `json:"bi_config"`
	logger *log.Logger
}



func NewProxy(fp string)*Proxy  {
	var p Proxy

	conf,_ := ioutil.ReadFile(fp)
	fmt.Println(string(conf))

	yaml.Unmarshal(conf,&p)
	p.logger = log.New(os.Stdout,"proxy ", log.Lshortfile|log.Ltime)
	p.logger.Println(p)
	return &p
}

func parseBody(data string)(payload,signature []byte){
	s,_:=url.ParseQuery(data)
	return []byte(s.Get("payload")),[]byte(s.Get("signature"))
}

// hello world, the web server
func (p Proxy)ProxyServer(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost{
		http.Redirect(w,req,req.URL.Path,304)
		return
	}

	log.Println("Method: ",req.Method)
	log.Println("Cookies: ",req.Cookies())
	log.Println("Form: ",req.Form)
	log.Println()
	//body,_ := req.GetBody()
	data,_:=ioutil.ReadAll(req.Body)
	log.Println(string(data))

	var auth QCAuth
	auth.Payload,auth.Signature = parseBody(string(data))
	auth.SecretKey = []byte(p.QYConfig.SecretAccessKey)
	if !auth.Authentication(){
		p.logger.Println("authentication error")
		w.WriteHeader(401)
		w.Write([]byte("Unauthorized"))
		return
	}
	p.logger.Println("authentication success")
	qcinfo:=auth.ExtractPayload()
	p.logger.Println(qcinfo)
	user,err := p.GetUser(auth.Payload,auth.Signature)
	if err != nil{
		w.WriteHeader(500)
		w.Write([]byte("something error"))
		return
	}
	resStr,_:=json.Marshal(user)
	io.WriteString(w, string(resStr))
}


func (p Proxy)GetUser(payload, signature []byte)(User,error){
	var auth QCAuth
	var qcInfo QCinfo
	var rb DescribeUsersResp
	var user User
	user.BiUri = p.BIConfig.BiUri
	auth.Payload,auth.Signature = payload,signature
	auth.SecretKey = []byte(p.QYConfig.SecretAccessKey)

	//if !auth.Authentication() {
	//	return &user, errors.New("authentication error.")
	//}

	qcInfo =auth.ExtractPayload()
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
		"2018-01-13T05:14:17Z",
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
		"	app-ptfa3keh",
	})
	method := "GET"
	endPoint:=p.QYConfig.Protocol+"://"+p.QYConfig.Host
	uri :=p.QYConfig.URI
	secretKey := p.QYConfig.SecretAccessKey
	p.logger.Println(method,endPoint,uri,secretKey)
	urlStr:=ps.GenerateReqStr(method,endPoint,uri,secretKey)
	p.logger.Println(urlStr)
	resp,err := http.Get(urlStr)

	if err!=nil{
		p.logger.Println(err.Error())
		return user,err
	}
	defer resp.Body.Close()
	body,err:= ioutil.ReadAll(resp.Body)
	if err!=nil{
		p.logger.Println(err.Error())
		return user,err
	}
	json.Unmarshal(body,&rb)
	if rb.RetCode == 0 && len(rb.UserSet)>0{
		json.Unmarshal(rb.UserSet[0],&user)
	}
	p.logger.Println(user)
	return user,nil

}
