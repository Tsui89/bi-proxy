package bi

import "log"

type Config struct {
	APIURI       string `json:"api_uri"`
	Adminv       string `json:"adminv"`
	AdminPassv   string `json:"admin_passv"`
	DefaultPassv string `json:"default_passv"`
	DefaultGroup string `json:"default_group"`
	RedirectUri  string `json:"redirect_uri"`
}

type User struct {
	UserName      string `json:"user_name"`
	Email         string `json:"email"`
	UserId        string `json:"user_id"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	Status        string `json:"status"`
	CompanyPhone  string `json:"company_phone"`
	CompanyName   string `json:"company_name"`
	CompanyCode   string `json:"company_code"`
	PersonalName  string `json:"personal_name"`
	PersonalCode  string `json:"personal_code"`
	GravatarEmail string `json:"gravatar_email"`
	NotifyEmail   string `json:"notify_email"`
	VerifyType    string `json:"verify_type"`
	Currency      string `json:"currency"`
}

type BIConfig struct {
	Config Config `json:"config"`
	User   User   `json:"user"`
	Logger *log.Logger
	Token  string
}
