package api_lib

import "os"

type GitHubInfo struct {
	ApiURL     string //api地址
	Repository string //仓库名称
	Sha        string //提交hash
	Token      string //token
}

func (t *GitHubInfo) InGitHub() bool {
	return t.Repository != ""
}

func NewGitHubInfo() *GitHubInfo {
	return &GitHubInfo{
		ApiURL:     os.Getenv("GITHUB_API_URL"),
		Repository: os.Getenv("GITHUB_REPOSITORY"),
		Sha:        os.Getenv("GITHUB_SHA"),
		Token:      os.Getenv("GITHUB_TOKEN"),
	}
}
