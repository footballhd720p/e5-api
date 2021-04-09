package api_lib

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
)

func SaveTokenInfoToFile(tokenInfo *TokenInfo) error {
	jsonData, err := json.Marshal(tokenInfo)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(tokenFilePath, jsonData, 0644); err != nil {
		return err
	}
	return nil
}

func SaveTokenInfoToGithub(tokenInfo *TokenInfo, githubInfo *GitHubInfo) error {
	client := resty.New()
	client.SetAuthScheme("token")
	client.SetAuthToken(githubInfo.Token)
	return nil
}
