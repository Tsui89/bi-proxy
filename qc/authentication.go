package qc

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
	"encoding/json"
)

type QCAuth struct{
	Payload []byte
	Signature []byte
	SecretKey []byte
}

type QCinfo struct{
	Lang string `json:"lang"`
	AccessToken string `json:"access_token"`
	UserId string `json:"user_id"`
	Zone string `json:"zone"`
	Action string `json:"action"`
	TimeStamp string `json:"time_stamp"`
	Expires string `json:"expires"`
}
//[]byte("TpsMi5ejGAkkaaMYDsY9EFZQMnNEI2QxuPNb0qlJ")
func (auth QCAuth)Authentication()bool{
	hash := hmac.New(sha256.New, auth.SecretKey)
	hash.Write(auth.Payload)
	base64Payload := base64.StdEncoding.EncodeToString(hash.Sum(nil))

	index:=len(base64Payload)-1
	for index>0{
		if (base64Payload[index] == '='){
			index --
		}else{
			break
		}
	}
	////去掉尾部 "="
	r1 := string(base64Payload[:index+1])
	//r1:= strings.Split(base64Payload,"=")[0]
	//replace + => -
	r2 := strings.Replace(r1, "+", "-", -1)
	//replace / => _
	r3 := strings.Replace(r2, "/", "_", -1)

	//认证
	if (string(auth.Signature) == r3){
		return true
	}
	return false
}

func (auth QCAuth)ExtractPayload()QCinfo {

	//payload base64 URL 解码
	info, _ := base64.StdEncoding.DecodeString(string(auth.Payload) + strings.Repeat("=", (4 - len(auth.Payload)%4)))
	var qcinfo QCinfo
	json.Unmarshal(info, &qcinfo)

	return qcinfo
}
