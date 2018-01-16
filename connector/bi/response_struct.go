package bi

import "encoding/xml"

//response get token

type Result struct {
	Level int `xml:"level"`
	Message string `xml:"message"`
}

type TokenResp struct {
	XMLName xml.Name `xml:"results"`
	TokenResult []Result `xml:"result"`
}

//type TokenResult struct{
//	Level int `xml:"level"`
//	Message string `xml:"message"`
//}

//response get user
type UserResp struct{
	XMLName xml.Name `xml:"results"`
	UserResult []UserResult `xml:"result"`
}
type UserResult struct {
	Asset UserAsset `xml:"asset"`
	//User UserInfo `xml:"user"`
}
type UserAsset struct {
	UserInfo UserInfo
}

type UserInfo struct {
	XMLName xml.Name `xml:"user"`
	IsOverWrite string `xml:"isOverWrite,attr"`
	Name string `xml:"name"`
	Alias string `xml:"alias"`
	Email string `xml:"email"`
	Parent string `xml:"parent"`
	Roles string	`xml:"roles"`
	Pass string `xml:"pass"`
}

//respone  action default
type Results struct{
	XMLName xml.Name `xml:"results"`
	Result []Result `xml:"result"`
}

//response get group
type GroupResp struct{
	XMLName xml.Name `xml:"results"`
	GroupResult []GroupResult `xml:"result"`
}
type GroupResult struct {
	Asset GroupAsset `xml:"asset"`
	//User UserInfo `xml:"user"`
}
type GroupAsset struct {
	GroupInfo GroupInfo
}

type GroupInfo struct {
	XMLName xml.Name `xml:"group"`
	Name string `xml:"name"`
	Parent string `xml:"parent"`
	Roles string	`xml:"roles"`
}
