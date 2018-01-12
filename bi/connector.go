package bi

import (
	"encoding/json"
	"github.com/Tsui89/bi-proxy/qc"
	"net/http"
)

func NewBi(config json.RawMessage) qc.Connector {
	var b BIConfig
	json.Unmarshal(config,&b)

	return &b
}

//Connect()


func (b *BIConfig)Connect(){}

func (b *BIConfig)Close(){}

func (b *BIConfig) IsUserExist(qc.User)bool{
	return true
}

func (b *BIConfig)CreateUser(qc.User)error{
	return nil
}

func (b *BIConfig)Authorization(qc.User)error{
	return nil
}

func (b *BIConfig)Redirect(qc.User)http.Handler{
	return http.RedirectHandler("",304)
}