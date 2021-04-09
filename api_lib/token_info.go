package api_lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type TokenInfo struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

const (
	tokenFilePath         = "token.json"
	tokenSecretKey        = "APP_TOKEN"
	refreshTokenSecretKey = "APP_R_TOKEN"
)

func NewTokenInfo(githubInfo *GitHubInfo) (*TokenInfo, error) {
	if !githubInfo.InGitHub() {
		return newTokenFromFile()
	}
	//读取Github Secret
	appToken := os.Getenv(tokenSecretKey)
	if appToken == "" {
		return nil, errors.New(tokenSecretKey + " 未定义")
	}
	appRefreshToken := os.Getenv(refreshTokenSecretKey)
	if appRefreshToken == "" {
		return nil, errors.New(refreshTokenSecretKey + " 未定义")
	}
	return &TokenInfo{
		AccessToken:  appToken,
		RefreshToken: appRefreshToken,
	}, nil
}

//从文件读取token信息
func newTokenFromFile() (*TokenInfo, error) {
	fileData, err := ioutil.ReadFile(tokenFilePath)
	if err != nil {
		return nil, err
	}
	tokenInfo := new(TokenInfo)
	if err := json.Unmarshal(fileData, tokenInfo); err != nil {
		return nil, err
	}
	return tokenInfo, nil
}
