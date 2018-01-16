package bi

import (
	"testing"
	"encoding/xml"
)

func TestUserReqStruct(t *testing.T){
	xmlStr :="<?xml version=\"1.0\" encoding=\"UTF-8\"?><ref><type>user</type><name>test1</name></ref>"

	var r Req
	xml.Unmarshal([]byte(xmlStr),&r)
	if r.Name != "test1"{
		t.Fail()
	}
}
