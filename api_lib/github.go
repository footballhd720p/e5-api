package api_lib

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
)

type GitHubInfo struct {
	ApiURL     string //api地址
	Repository string //仓库名称
	Sha        string //提交hash
	Token      string //token
}

func NewGitHubInfo() *GitHubInfo {
	return &GitHubInfo{
		ApiURL:     os.Getenv("GITHUB_API_URL"),
		Repository: os.Getenv("GITHUB_REPOSITORY"),
		Sha:        os.Getenv("GITHUB_SHA"),
		Token:      os.Getenv("GITHUB_TOKEN"),
	}
}

func (t *GitHubInfo) InGitHub() bool {
	return t.Repository != ""
}

func (t *GitHubInfo) UpdateSecret(secretName string, secretValue string) error {
	client := resty.New()
	client.SetAuthScheme("Bearer")
	client.SetAuthToken(t.Token)
	apiUrl := t.ApiURL + "/repos/" + t.Repository + "/actions/secrets/" + secretName
	resp, err := client.R().
		EnableTrace().
		SetBody(map[string]string{
			"encrypted_value": secretValue,
		}).Put(apiUrl)
	if err != nil {
		return err
	}
	// Explore response object
	fmt.Println("code", resp.StatusCode())
	fmt.Println(apiUrl)
	fmt.Println(resp.Request.Header)
	fmt.Println(resp)
	return nil
}
