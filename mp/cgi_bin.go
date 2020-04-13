package mp

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/tidwall/gjson"

	"github.com/arden/wechat/utils"
)

// AccessToken wxmp access_token
type AccessToken struct {
	Token     string `json:"access_token"`
	ExpiresIn int64  `json:"expires_in"`
}

// CgiBin cgi-bin
type CgiBin struct {
	mp      *WXMP
	options []utils.HTTPRequestOption
}

// GetAccessToken returns access_token
func (p *CgiBin) GetAccessToken() (*AccessToken, error) {
	resp, err := p.mp.Client.Get(fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", AccessTokenURL, p.mp.AppID, p.mp.AppSecret), p.options...)

	if err != nil {
		return nil, err
	}

	r := gjson.ParseBytes(resp)

	if r.Get("errcode").Int() != 0 {
		return nil, errors.New(r.Get("errmsg").String())
	}

	reply := new(AccessToken)

	if err := json.Unmarshal(resp, reply); err != nil {
		return nil, err
	}

	return reply, nil
}
