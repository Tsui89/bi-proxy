package qc

import (
	"testing"
	"fmt"
)

func TestParameters_GenerateReqStr(t *testing.T) {
	ps :=NewParameters()
	var p1 Parameter=Parameter{
		"count",
		1,
	}
	var p2 Parameter=Parameter{
		"vxnets.1",
		"vxnet-0",
	}

	var p3 Parameter=Parameter{
		"zone",
		"pek1",
	}
	var p4 Parameter=Parameter{
		"instance_type",
		"small_b",
	}
	p5:=Parameter{
		"signature_version",
		1,
	}
	p6:=Parameter{
		"signature_method",
		"HmacSHA256",
	}
	p7:=Parameter{
		"instance_name",
		"demo",
	}
	p8:=Parameter{
		"image_id",
		"centos64x86a",
	}
	p9:=Parameter{
		"login_mode",
		"passwd",
	}
	p10:=Parameter{
		"login_passwd",
		"QingCloud20130712",
	}
	p11:=Parameter{
		"version",
		1,
	}
	p12:=Parameter{
		"access_key_id",
		"QYACCESSKEYIDEXAMPLE",
	}
	p13:=Parameter{
		"action",
		"RunInstances",
	}
	p14:=Parameter{
		"time_stamp",
		"2013-08-27T14:30:10Z",
	}
	ps.Add(&p1)
	ps.Add(&p2)
	ps.Add(&p3)
	ps.Add(&p4)
	ps.Add(&p5)
	ps.Add(&p6)
	ps.Add(&p7)
	ps.Add(&p8)
	ps.Add(&p9)
	ps.Add(&p10)
	ps.Add(&p11)
	ps.Add(&p12)
	ps.Add(&p13)
	ps.Add(&p14)
	res :=ps.GenerateReqStr("GET","https://api.qingcloud.com","iaas","SECRETACCESSKEY")
	if (res != "https://api.qingcloud.com/iaas/?access_key_id=QYACCESSKEYIDEXAMPLE&action=RunInstances&count=1&image_id=centos64x86a&instance_name=demo&instance_type=small_b&login_mode=passwd&login_passwd=QingCloud20130712&signature_method=HmacSHA256&signature_version=1&time_stamp=2013-08-27T14%3A30%3A10Z&version=1&vxnets.1=vxnet-0&zone=pek1&signature=32bseYy39DOlatuewpeuW5vpmW51sD1A%2FJdGynqSpP8%3D"){
		t.Fatal()

	}

	p1=Parameter{
		"users.1",
		"usr-ibGnrli6",
	}
	p2=Parameter{
		"action",
		"DescribeUsers",
	}
	p3=Parameter{
		"time_stamp",
		"2018-01-13T05:14:17Z",
	}

	p4=Parameter{
		"access_key_id",
		"WJJWLKWMBVUOYDTOTMPP",
	}

	p5=Parameter{
		"version",
		1,
	}
	p6=Parameter{
		"signature_method",
		"HmacSHA256",
	}
	p7=Parameter{
		"signature_version",
		1,
	}

	ps2 := NewParameters()
	ps2.Add(&p1)
	ps2.Add(&p2)
	ps2.Add(&p3)
	ps2.Add(&p4)
	ps2.Add(&p5)
	ps2.Add(&p6)
	ps2.Add(&p7)

	fmt.Println(ps2.GenerateReqStr("GET","http://api.kunlcloud.com","iaas","TXHU8yAeIOGIlZAmMNpLFQdHnAmZYzaAlYuEGrHl"))
}
//http://api.kunlcloud.com/iaas/?access_key_id=WJJWLKWMBVUOYDTOTMPP&action=DescribeUsers&signature_method=HmacSHA256&signature_version=1&time_stamp=2018-01-13T05%3A14%3A17Z&users.1=usr-KTVYs3cg&version=1&signature=sza%2F5bRJQWwaN0v0pXLn6JOWa18mOSWnC1ReWDH50Fs%3D
//http://api.kunlcloud.com/iaas/?access_key_id=WJJWLKWMBVUOYDTOTMPP&action=DescribeUsers&signature_method=HmacSHA256&signature_version=1&time_stamp=2018-01-13T05%3A14%3A17Z&users.1=usr-KTVYs3cg&version=1&signature=Mw0NlMY05i3bNEvYfT05i4ryBqN9FYNloqT62XEttr8%3D
