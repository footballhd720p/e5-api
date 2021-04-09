package api_lib

import (
	"encoding/json"
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
	if err := githubInfo.UpdateSecret(tokenSecretKey, tokenInfo.AccessToken); err != nil {
		return err
	}
	if err := githubInfo.UpdateSecret(refreshTokenSecretKey, tokenInfo.RefreshToken); err != nil {
		return err
	}
	return nil
}
