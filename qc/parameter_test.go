package qc

import (
	"testing"
	"fmt"
	"strings"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/url"
)

func TestParameters_GenerateReqStr(t *testing.T) {

	//uri := "iaas/"
	ru,_ := url.Parse("https://api.qingcloud.com/iaas/")
	//ru.Path = path.Join(ru.RawPath,uri)
	rq,_ := url.ParseQuery(ru.RawQuery)

	rq.Add(
		"count",
		fmt.Sprint(1),
	)
	rq.Add(
		"vxnets.1",
		"vxnet-0",
	)

	rq.Add(
		"zone",
		"pek1",
	)
	rq.Add(
		"instance_type",
		"small_b",
	)
	rq.Add(
		"signature_version",
		fmt.Sprint(1),
	)
	rq.Add(
		"signature_method",
		"HmacSHA256",
	)
	rq.Add(
		"instance_name",
		"demo",
	)
	rq.Add(
		"image_id",
		"centos64x86a",
	)
	rq.Add(
		"login_mode",
		"passwd",
	)
	rq.Add(
		"login_passwd",
		"QingCloud20130712",
	)
	rq.Add(
		"version",
		fmt.Sprint(1),
	)
	rq.Add(
		"access_key_id",
		"QYACCESSKEYIDEXAMPLE",
	)
	rq.Add(
		"action",
		"RunInstances",
	)
	rq.Add(
		"time_stamp",
		"2013-08-27T14:30:10Z",
	)
	fmt.Println(ru.Path)
	sc := strings.Join([]string{"GET", ru.Path, rq.Encode()}, "\n")
	//fmt.Println(urlStr=="GET\n/iaas/\naccess_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1")
	hash := hmac.New(sha256.New, []byte("SECRETACCESSKEY"))
	hash.Write([]byte(sc))
	base64Payload := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	index := strings.LastIndex(base64Payload, "=")
	r := strings.Replace(base64Payload[:index+1], " ", "+", -1)
	//rq.Add("signature",r)
	ru.RawQuery = rq.Encode()+"&"+"signature="+url.QueryEscape(r)
	fmt.Println(rq.Encode(),ru.RawQuery)
	urlStr := ru.String()
	fmt.Println(urlStr)
	//res := ps.GenerateReqStr("GET", "https://api.qingcloud.com", "iaas", "SECRETACCESSKEY")
	if (urlStr != "https://api.qingcloud.com/iaas/?access_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1&signature=32bseYy39DOlatuewpeuW5vpmW51sD1A%2FJdGynqSpP8%3D") {
		t.Fatal()

	}
}

//http://api.kunlcloud.com/iaas/?access_key_id=WJJWLKWMBVUOYDTOTMPP&action=DescribeUsers&signature_method=HmacSHA256&signature_version=1&time_stamp=2018-01-13T05%3A14%3A17Z&users.1=usr-KTVYs3cg&version=1&signature=sza%2F5bRJQWwaN0v0pXLn6JOWa18mOSWnC1ReWDH50Fs%3D
//http://api.kunlcloud.com/iaas/?access_key_id=WJJWLKWMBVUOYDTOTMPP&action=DescribeUsers&signature_method=HmacSHA256&signature_version=1&time_stamp=2018-01-13T05%3A14%3A17Z&users.1=usr-KTVYs3cg&version=1&signature=Mw0NlMY05i3bNEvYfT05i4ryBqN9FYNloqT62XEttr8%3D
