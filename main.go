package main

import (
	"fmt"
	"github.com/liuguangw/e5-api/api_lib"
	"log"
)

func main() {
	githubInfo := api_lib.NewGitHubInfo()
	if githubInfo.InGitHub() {
		fmt.Println("GitHub模式")
		githubInfo.ShowInfo()
	} else {
		fmt.Println("普通模式")
	}
	tokenInfo, err := api_lib.NewTokenInfo(githubInfo)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("access_token=" + tokenInfo.AccessToken + "\n" + "refresh_token" + tokenInfo.RefreshToken)
	//测试更新token
	tokenInfo.AccessToken = "22222"
	tokenInfo.RefreshToken = "333333"
	if err := api_lib.SaveTokenInfo(tokenInfo, githubInfo); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("access_token=" + tokenInfo.AccessToken + "\n" + "refresh_token" + tokenInfo.RefreshToken)
}
