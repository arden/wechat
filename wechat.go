package wechat

import (
	"crypto/tls"

	"github.com/arden/wechat/mch"
	"github.com/arden/wechat/mp"
	"github.com/arden/wechat/pub"
	"github.com/arden/wechat/utils"
)

// NewWXMch 微信商户
func NewWXMch(appid, mchid, apikey string, cert tls.Certificate) *mch.WXMch {
	wxmch := &mch.WXMch{
		AppID:  appid,
		MchID:  mchid,
		ApiKey: apikey,
		Client: utils.DefaultHTTPClient,
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

	wxmch.SSLClient = utils.NewHTTPClient(utils.WithHTTPTLSConfig(tlsConfig))

	return wxmch
}

// NewWXPub 微信公众号
func NewWXPub(accountid, appid, appsecret, signToken, encodingAESKey string) *pub.WXPub {
	wxpub := &pub.WXPub{
		AccountID:      accountid,
		AppID:          appid,
		AppSecret:      appsecret,
		SignToken:      signToken,
		EncodingAESKey: encodingAESKey,
		Client:         utils.DefaultHTTPClient,
	}

	return wxpub
}

// NewWXMP 微信小程序
func NewWXMP(appid, appsecret string) *mp.WXMP {
	wxmp := &mp.WXMP{
		AppID:     appid,
		AppSecret: appsecret,
		Client:    utils.DefaultHTTPClient,
	}

	return wxmp
}
