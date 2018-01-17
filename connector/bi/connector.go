package bi

import (
	"encoding/json"
	"net/http"
	"strings"
	"fmt"
	"log"
	"io/ioutil"
	"encoding/xml"
	"github.com/pkg/errors"
)

func NewBi(config json.RawMessage, logger *log.Logger) *BIConfig {
	var b BIConfig
	b.Logger = logger
	json.Unmarshal(config, &(b.Config))

	return &b
}

func (b *BIConfig) SetUser(user json.RawMessage) error {
	json.Unmarshal(user, &(b.User))
	return nil
}

func generateParamStr(p map[string]interface{}) string {
	var res []string
	for k, v := range p {
		res = append(res, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(res, "&")
}

//func generateXmlBody(p map[string]interface{})string{
//
//}

func (b *BIConfig) requestPost(urlStr string, f map[string]string) (body []byte, err error) {
	var r http.Request
	//r.Method = http.MethodPost
	r.ParseForm()
	for k, v := range f {
		r.Form.Add(k, v)
	}
	b.Logger.Println(urlStr)

	bodystr := strings.TrimSpace(r.Form.Encode())
	b.Logger.Println(bodystr)
	request, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(bodystr))
	if err != nil {
		b.Logger.Println(err.Error())
		return []byte{}, err //TODO:
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Connection", "Keep-Alive")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	//resp,err := http.Post(urlStr,"Content-Type: application/x-www-form-urlencoded",strings.NewReader(r.Form.Encode()))
	if err != nil {
		b.Logger.Println(err.Error())
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		b.Logger.Println(err.Error())
		return []byte{}, err
	}
	return body, nil
}
func (b *BIConfig) getToken(name, password string) (string, error) {
	params := map[string]interface{}{
		"action": ActionLogin,
		"adminv": name,
		"passv":  password,
	}
	paramStr := generateParamStr(params)

	urlStr := strings.Join(
		[]string{strings.Join([]string{b.Config.URI, UriAPI}, "/"), paramStr}, "?")
	b.Logger.Println(urlStr)

	resp, err := http.Get(urlStr)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	b.Logger.Println(string(body))
	var t TokenResp
	xml.Unmarshal(body, &t)
	b.Logger.Println(t)
	if len(t.TokenResult) > 0 {
		if t.TokenResult[0].Level != 6 {
			return t.TokenResult[0].Message, nil
		}
	}
	return "", errors.New("get token error.")

}
func (b *BIConfig) Connect() error {
	token, err := b.getToken(b.Config.Adminv, b.Config.AdminPassv)
	if err != nil {
		return err
	}
	b.Token = token
	if b.IsGroupExist() == false {
		err := b.CreateDefaultGroup()
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BIConfig) Close() error {
	params := map[string]interface{}{
		"action": ActionLogout,
		"token":  b.Token,
	}
	paramStr := generateParamStr(params)

	urlStr := strings.Join(
		[]string{strings.Join([]string{b.Config.URI, UriAPI}, "/"), paramStr}, "?")
	b.Logger.Println(urlStr)

	resp, err := http.Get(urlStr)
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	b.Logger.Println(string(body))
	var r Results
	xml.Unmarshal(body, &r)
	if r.Result[0].Level != 6 {
		return nil
	}
	return errors.New(r.Result[0].Message)
}

//func (b *BIConfig) GetUser()
func (b *BIConfig) Refresh() error {
	if b.IsConnected() == true {
		b.Close()
	}
	b.Connect()
	return nil
}

func (b *BIConfig) IsConnected() bool {
	if b.Token == "" {
		return false
	}

	if b.isUserExist(b.Config.Adminv, "") {
		return true
	}
	return false
}

func (b *BIConfig) IsUserExist() bool {
	return b.isUserExist(b.User.UserId, b.Config.DefaultGroup)
}

func (b *BIConfig) IsGroupExist() bool {
	params := map[string]interface{}{
		"action": ActionGetNode,
		"token":  b.Token,
	}
	paramStr := generateParamStr(params)
	var reqData Req
	reqData.Name = b.Config.DefaultGroup
	reqData.Type = "group"
	body, err := b.requestGet(paramStr, reqData)
	if err != nil {
		return false
	}
	b.Logger.Println(string(body))
	var t GroupResp
	xml.Unmarshal(body, &t)
	b.Logger.Println(t)
	if len(t.GroupResult) > 0 {
		b.Logger.Println(t.GroupResult[0].Asset == GroupAsset{})
		//是否为空
		if (t.GroupResult[0].Asset == GroupAsset{}) {
			return false
		}
		return true
	}
	//没有result
	return false
}

func (b *BIConfig) requestGet(paramStr string, reqData interface{}) ([]byte, error) {

	urlStr := strings.Join(
		[]string{strings.Join([]string{b.Config.URI, UriAPI}, "/"), paramStr}, "?")

	data, _ := xml.Marshal(reqData)
	body, err := b.requestPost(urlStr, map[string]string{
		"xmlData": string(data),
	})
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (b *BIConfig) GetUserInfo(name, parent string) (UserInfo, error) {
	params := map[string]interface{}{
		"action":    ActionGetNode,
		"token":     b.Token,
		"returnPwd": "yes",
	}
	paramStr := generateParamStr(params)
	var reqData Req
	reqData.Name = name
	reqData.Type = "user"
	body, err := b.requestGet(paramStr, reqData)
	if err != nil {
		return UserInfo{}, err
	}
	b.Logger.Println(string(body))
	var t UserResp
	xml.Unmarshal(body, &t)
	b.Logger.Println(t)
	if len(t.UserResult) > 0 {
		b.Logger.Println(t.UserResult[0].Asset == UserAsset{})
		//是否为空
		for _, result := range t.UserResult {
			if (result.Asset == UserAsset{}) {
				return UserInfo{}, errors.New("has no such user.")
			}
			if parent != "" {
				parents := strings.Split(result.Asset.UserInfo.Parent, ",")
				for _, p := range parents {
					if p == parent {
						return result.Asset.UserInfo, nil
					}
				}

			}

		}
	}
	//没有result
	return UserInfo{}, errors.New("has no such user.")

}

func (b *BIConfig) isUserExist(name, parent string) bool {
	_, err := b.GetUserInfo(name, parent)
	if err != nil {
		return false

	}
	//没有result
	return true
}

func (b *BIConfig) CreateDefaultGroup() error {
	params := map[string]interface{}{
		"action": ActionSaveNode,
		"token":  b.Token,
		"type":   "group",
	}
	paramStr := generateParamStr(params)

	urlStr := strings.Join(
		[]string{strings.Join([]string{b.Config.URI, UriAPI}, "/"), paramStr}, "?")
	var g GroupSaveReq
	g.GroupInfo.Name = b.Config.DefaultGroup
	data, _ := xml.Marshal(g)
	body, err := b.requestPost(urlStr, map[string]string{
		"xmlData": string(data),
	})
	b.Logger.Println(string(body))

	if err != nil {
		return err
	}
	var results Results
	xml.Unmarshal(body, &results)
	for _, r := range results.Result {
		if r.Level != 6 {
			return nil
		}
	}
	return errors.New("create default group error.")

}
func (b *BIConfig) CreateUser() error {

	params := map[string]interface{}{
		"action": ActionSaveNode,
		"token":  b.Token,
		"type":   "user",
	}
	paramStr := generateParamStr(params)

	urlStr := strings.Join(
		[]string{strings.Join([]string{b.Config.URI, UriAPI}, "/"), paramStr}, "?")
	var u UserSaveReq
	u.UserInfo.IsOverWrite = "true"
	u.UserInfo.Name = b.User.UserId
	u.UserInfo.Alias = b.User.UserName
	u.UserInfo.Email = b.User.Email
	u.UserInfo.Password = b.Config.DefaultPassv
	u.UserInfo.Parent = b.Config.DefaultGroup
	data, _ := xml.Marshal(u)
	body, err := b.requestPost(urlStr, map[string]string{
		"xmlData": string(data),
	})
	b.Logger.Println(string(body))
	if err != nil {
		return err
	}

	var results Results
	xml.Unmarshal(body, &results)
	for _, r := range results.Result {
		if r.Level != 6 {
			return nil
		}
	}
	return errors.New("create user error.")

}

func (b *BIConfig) Authorization() error {
	return nil
}

func (b *BIConfig) Redirect(w http.ResponseWriter, req *http.Request) {
	//
	//userToken,err:= b.getToken(b.User.UserId,b.Config.DefaultPassv)
	//if err !=nil{
	//	w.WriteHeader(500)
	//	w.Write([]byte(err.Error()))
	//	return
	//}
	//h := md5.New()
	//h.Write([]byte(b.Config.DefaultPassv)) // 需要加密的字符串为 123456
	//cipherStr := h.Sum(nil)
	user, err := b.GetUserInfo(b.User.UserId, b.Config.DefaultGroup)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}

	params := map[string]interface{}{
		"au_act": ActionLogin,
		"adminv": b.User.UserId,
		//"passv":string(cipherStr),
		"passv": user.Password,
		//"token":userToken,
	}
	paramStr := generateParamStr(params)

	urlStr := strings.Join(
		[]string{strings.Join([]string{b.Config.URI, UriViewer}, "/"), paramStr}, "?")
	b.Logger.Println(urlStr)

	http.Redirect(w, req, urlStr, http.StatusFound)
	return
	//return http.RedirectHandler(urlStr,304)
}
