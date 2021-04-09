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
		fmt.Println(githubInfo)
	} else {
		fmt.Println("普通模式")
	}
	tokenInfo, err := api_lib.NewTokenInfo(githubInfo)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("access_token=" + tokenInfo.AccessToken + "\n" + "refresh_token" + tokenInfo.RefreshToken)
}
