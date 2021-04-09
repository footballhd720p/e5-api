package api_lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func SaveTokenInfo(tokenInfo *TokenInfo, githubInfo *GitHubInfo) error {
	if githubInfo.InGitHub() {
		return saveTokenInfoToGithub(tokenInfo, githubInfo)
	}
	return saveTokenInfoToFile(tokenInfo)
}

func saveTokenInfoToFile(tokenInfo *TokenInfo) error {
	jsonData, err := json.Marshal(tokenInfo)
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(tokenFilePath, jsonData, 0644); err != nil {
		return err
	}
	return nil
}

func saveTokenInfoToGithub(tokenInfo *TokenInfo, githubInfo *GitHubInfo) error {
	envStr := tokenSecretKey + "=" + tokenInfo.AccessToken
	envStr += "\n" + refreshTokenSecretKey + "=" + tokenInfo.RefreshToken
	if err := githubInfo.WriteEnvData(envStr); err != nil {
		return err
	}
	//日志中添加掩码
	fmt.Println("::add-mask::" + tokenInfo.AccessToken)
	fmt.Println("::add-mask::" + tokenInfo.RefreshToken)
	return nil
}
