package api_lib

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

type JwtTokenInfo struct {
	AppDisplayname string `json:"app_displayname"` //应用名称
	Name           string `json:"name"`            //使用者名称
	UniqueName     string `json:"unique_name"`     //使用者邮箱
	Iat            int64  `json:"iat"`             //token签发时间
	Exp            int64  `json:"exp"`             //token过期时间
}

const timeLayout = "2006-01-02 15:04:05"

func (t *JwtTokenInfo) ShowTokenInfo() {
	fmt.Println("应用名称: " + t.AppDisplayname)
	fmt.Println("使用者名称: " + t.Name)
	fmt.Println("使用者邮箱: " + t.UniqueName)
	tm := time.Unix(t.Iat, 0)
	fmt.Println("access token签发时间: " + tm.Format(timeLayout))
	tm = time.Unix(t.Exp, 0)
	fmt.Println("access token过期时间: " + tm.Format(timeLayout))
}

func NewJwtTokenInfo(tokenStr string) (*JwtTokenInfo, error) {
	pos := strings.Index(tokenStr, ".")
	if pos < 0 {
		return nil, errors.New("invalid access_token")
	}
	tmpStr := tokenStr[pos+1:]
	pos = strings.Index(tmpStr, ".")
	if pos < 0 {
		return nil, errors.New("invalid access_token")
	}
	tmpStr = tmpStr[:pos]
	data, err := base64.RawURLEncoding.DecodeString(tmpStr)
	if err != nil {
		return nil, err
	}
	jwtTokenInfo := new(JwtTokenInfo)
	if err := json.Unmarshal(data, jwtTokenInfo); err != nil {
		return nil, err
	}
	return jwtTokenInfo, nil
}
