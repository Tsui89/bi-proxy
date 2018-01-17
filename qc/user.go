package qc

import "encoding/json"

type DescribeUsersResp struct {
	Action     string            `json:"action"`
	TotalCount int               `json:"total_count"`
	UserSet    []json.RawMessage `json:"user_set"`
	RetCode    int               `json:"ret_code"`
}

//"company_code": "",
//"notify_email": "ldapuser1@kunlcloud.com",
//"verify_type": "personal",
//"user_type": 0,
//"gravatar_email": "ldapuser1@kunlcloud.com",
//"zones": [
//"poc1"
//],
//"currency": "cny",
//"create_time": "2018-01-11T10:47:37Z",
//"personal_code": "",
//"is_email_confirmed": 1,
//"user_id": "usr-ibGnrli6",
//"personal_name": "",
//"industry_category": "",
//"company_name": "",
//"user_name": "ldapuser1",
//"email": "ldapuser1@kunlcloud.com",
//"status": "active",
//"company_phone": "",
//"is_phone_verified": 0,
//"phone": "",
//"birthday": null,
//"preference": 1,
//"address": "",
//"total_sub_user_count": 0,
//"lang": "zh-cn",
//"gender": "male",
//"nologin": 0,
//"verify_status": "new"
type User struct {
	UserName      string   `json:"user_name"`
	Email         string   `json:"email"`
	UserId        string   `json:"user_id"`
	Phone         string   `json:"phone"`
	Address       string   `json:"address"`
	Status        string   `json:"status"`
	CompanyPhone  string   `json:"company_phone"`
	CompanyName   string   `json:"company_name"`
	CompanyCode   string   `json:"company_code"`
	PersonalName  string   `json:"personal_name"`
	PersonalCode  string   `json:"personal_code"`
	GravatarEmail string   `json:"gravatar_email"`
	NotifyEmail   string   `json:"notify_email"`
	VerifyType    string   `json:"verify_type"`
	Currency      string   `json:"currency"`
	Zones         []string `json:"zones"`
}
