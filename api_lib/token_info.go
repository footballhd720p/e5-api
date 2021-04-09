package api_lib

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

const (
	//本地运行时配置文件路径
	tokenFilePath = "token.json"

	//GitHub环境变量键名
	clientIdKey     = "APP_CLIENT_ID"
	clientSecretKey = "APP_CLIENT_SECRET"
	accessTokenKey  = "APP_TOKEN"
	refreshTokenKey = "APP_R_TOKEN"
)

type TokenInfo struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewTokenInfo(githubInfo *GitHubInfo) (*TokenInfo, error) {
	if !githubInfo.InGitHub() {
		return newTokenFromFile()
	}
	//读取Github Secret
	appToken := os.Getenv(accessTokenKey)
	if appToken == "" {
		return nil, errors.New(accessTokenKey + " 未定义")
	}
	appRefreshToken := os.Getenv(refreshTokenKey)
	if appRefreshToken == "" {
		return nil, errors.New(refreshTokenKey + " 未定义")
	}

	clientId := os.Getenv(clientIdKey)
	if clientId == "" {
		return nil, errors.New(clientIdKey + " 未定义")
	}
	clientSecret := os.Getenv(clientSecretKey)
	if clientSecret == "" {
		return nil, errors.New(clientSecretKey + " 未定义")
	}
	return &TokenInfo{
		ClientId:     clientId,
		ClientSecret: clientSecret,
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
