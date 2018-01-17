package bi

import "encoding/xml"

//<?xml version="1.0" encoding="UTF-8"?>
//
//<ref>
//	<type>user</type>
//	<name>test1</name>
//</ref>

//request get user
type Req struct {
	XMLName xml.Name `xml:"ref"`
	Type    string   `xml:"type"`
	Name    string   `xml:"name"`
}

//request save user

type UserSaveReq struct {
	XMLName  xml.Name `xml:"info"`
	UserInfo UserInfo
}

//request save group

type GroupSaveReq struct {
	XMLName   xml.Name `xml:"info"`
	GroupInfo GroupInfo
}
