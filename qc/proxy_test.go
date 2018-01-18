package qc

import (
	"testing"
	"log"
	"os"
)

func TestProxy_GetUser(t *testing.T) {
	var p Proxy
	p.QYConfig.AccessKeyID = "WJJWLKWMBVUOYDTOTMPP"
	p.QYConfig.SecretAccessKey = "YdqDyBaqxZ6vt0su8XN7Q2A1CsIp9Bo33ZTuz0LK"
	p.QYConfig.ApiUri = "http://api.kunlcloud.com/iaas/"
	p.logger = log.New(os.Stdout,"test ",log.Lshortfile|log.Ltime)
	var auth QCAuth
	auth.Signature = []byte("yDjiq4B04aSTw7AWx34NedRQogZoy3A5H2Oh7RMvwVA")
	auth.Payload = []byte("eyJsYW5nIjoiemgtY24iLCJhY2Nlc3NfdG9rZW4iOiJ0eWc1OVNGVjhRcGFpWjJBeWZpcVJ2VXFFV05UN2VMRSIsInVzZXJfaWQiOiJ1c3ItS1RWWXMzY2ciLCJ6b25lIjoicG9jMSIsImFjdGlvbiI6InZpZXdfYXBwIiwidGltZV9zdGFtcCI6IjIwMTgtMDEtMTJUMDg6NTY6NTFaIiwiZXhwaXJlcyI6IjIwMTgtMDEtMTJUMTQ6NTY6NTFaIn0")
	auth.SecretKey = []byte("YdqDyBaqxZ6vt0su8XN7Q2A1CsIp9Bo33ZTuz0LK")

	p.GetUser(auth.Payload, auth.Signature)
}
