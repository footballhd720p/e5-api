package api_lib

import (
	"fmt"
	"os"
)

type GitHubInfo struct {
	Repository string //仓库名称
	Sha        string //提交hash
	EnvFile    string //GitHub Action环境变量文件路径
}

func NewGitHubInfo() *GitHubInfo {
	return &GitHubInfo{
		Repository: os.Getenv("GITHUB_REPOSITORY"),
		Sha:        os.Getenv("GITHUB_SHA"),
		EnvFile:    os.Getenv("GITHUB_ENV"),
	}
}

func (t *GitHubInfo) InGitHub() bool {
	return t.Repository != ""
}

func (t *GitHubInfo) ShowInfo() {
	fmt.Println("仓库名: " + t.Repository)
	fmt.Println("提交SHA: " + t.Sha)
}

func (t *GitHubInfo) WriteEnvData(envStr string) error {
	file, err := os.OpenFile(t.EnvFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString("\n" + envStr); err != nil {
		return err
	}
	return nil
}
