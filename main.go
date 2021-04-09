package main

import (
	"fmt"
	"github.com/liuguangw/e5-api/api_lib"
	"log"
	"time"
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
	//读取token信息
	jwtTokenInfo, err := api_lib.NewJwtTokenInfo(tokenInfo.AccessToken)
	if err != nil {
		log.Fatalln(err)
	}
	jwtTokenInfo.ShowTokenInfo()
	if time.Now().Unix() > jwtTokenInfo.Exp {
		fmt.Println("token已过期,准备刷新token")
		//刷新token
		tokenInfo, err := tokenInfo.RefreshNew()
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("获取新token成功")
		//保存新的token
		if err := api_lib.SaveTokenInfo(tokenInfo, githubInfo); err != nil {
			log.Fatalln(err)
		}
		fmt.Println("保存新token成功")
	} else {
		fmt.Println("token未过期,无需刷新")
	}
	//请求Microsoft365 api
	//fmt.Println("access_token=" + newTokenInfo.AccessToken + "\n" + "refresh_token" + newTokenInfo.RefreshToken)
}
