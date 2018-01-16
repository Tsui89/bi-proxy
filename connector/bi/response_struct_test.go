package bi

import (
	"testing"
	"encoding/xml"
	"fmt"
)


func TestTokenResp(t *testing.T) {
	var r TokenResp
	xmlStr := "<?xml version=\"1.0\" encoding=\"UTF-8\"?> " +
		"<results><result><level>1</level><message>DB4F2DC2243BA0B6F8D1AFD8DD09BC0B</message></result></results>"
	xml.Unmarshal([]byte(xmlStr),&r)
	if (r.TokenResult[0].Message != "DB4F2DC2243BA0B6F8D1AFD8DD09BC0B"){
		t.Fatal()
	}
}

func TestUserResp(t *testing.T){
	xmlStr := "<?xml version=\"1.0\" encoding=\"UTF-8\"?> <results>"+
	"<result> <asset>"+
	"<user isOverWrite=\"true\"> <name>hunk1</name> <alias></alias>"+
	"<email>test@qq.com</email> <parent>test1</parent>"+
	"<roles></roles> </user>"+
	"</asset> </result></results>"
	var r UserResp
	xml.Unmarshal([]byte(xmlStr),&r)
	fmt.Println(r)
	//if (r.UserResult[0].User.Name!="hunk1"){
	//	t.Fatal()
	//}
}