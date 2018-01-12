package main

import (
	"fmt"
	"sort"
	"net/url"
	"strings"
	"crypto/sha256"
	"encoding/base64"
	"crypto/hmac"
)


type Parameter struct {
	Key   string
	Value interface{}
}

type Parameters []*Parameter

func (s Parameters) Len() int      { return len(s) }
func (s Parameters) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// ByName implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.
type ByKey struct{ Parameters }

func (s ByKey) Less(i, j int) bool { return s.Parameters[i].Key < s.Parameters[j].Key }

// ByWeight implements sort.Interface by providing Less and using the Len and
// Swap methods of the embedded Organs value.


func NewParameters()*Parameters  {
	var p Parameters
	return &p
}

func (p *Parameters)Add(m *Parameter){
	*p = append(*p,m)
}

func (p *Parameters)sort()string{
	sort.Sort(ByKey{*p})

	var paramsList =[]string{}
	for _,k :=range*p{
		paramsList = append(paramsList,fmt.Sprintf("%s=%s",url.QueryEscape(k.Key),url.QueryEscape(fmt.Sprint(k.Value))))
	}
	return strings.Join(paramsList,"&")

}
func (p *Parameters)generateReqParamsStr(method, service,secretAccessKey string)string{

	paramsStr := p.sort()
	urlStr := strings.Join([]string{method,"/"+service+"/",paramsStr},"\n")
	//fmt.Println(urlStr=="GET\n/iaas/\naccess_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1")
	hash := hmac.New(sha256.New, []byte(secretAccessKey))
	hash.Write([]byte(urlStr))
	base64Payload := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	index:=strings.LastIndex(base64Payload,"=")
	r := strings.Replace(base64Payload[:index+1], " ", "+", -1)

	return strings.Join([]string{paramsStr,"signature="+url.QueryEscape(r)},"&")
}

func (p *Parameters)GenerateReqStr(method, endPoint, uri,secretAccessKey string)string{
	ps := p.generateReqParamsStr(method,uri,secretAccessKey)
	return strings.Join([]string{endPoint,uri,"?"+ps},"/")
}