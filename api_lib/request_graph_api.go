package api_lib

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"time"
)

type ApiRequest struct {
	ApiPath string
	Title   string
}

func (req *ApiRequest) ProcessRequest(client *resty.Client) error {
	resp, err := client.R().Get(req.ApiPath)
	if err != nil {
		return err
	}
	statusCode := resp.StatusCode()
	if statusCode != 200 {
		return errors.New(resp.String())
	}
	fmt.Println("GET " + req.Title + " OK")
	//fmt.Println(resp.String())
	return nil
}

func RequestGraphApi(tokenInfo *TokenInfo) error {
	client := resty.New().
		SetAuthScheme("Bearer").
		SetAuthToken(tokenInfo.AccessToken).
		SetHostURL("https://graph.microsoft.com/v1.0")
	requestList := []*ApiRequest{
		{
			ApiPath: "/me/",
			Title:   "我的个人资料",
		},
		{
			ApiPath: "/users",
			Title:   "组织中的所有用户",
		},
		{
			ApiPath: "/me/drive/root/children",
			Title:   "我的驱动器中的所有项",
		},
		{
			ApiPath: "/me/drive/recent",
			Title:   "我最近使用的文件",
		},
		{
			ApiPath: "/me/messages",
			Title:   "我的邮件",
		},
		{
			ApiPath: "/me/mailFolders/inbox/messagerules",
			Title:   "我的收件箱规则",
		},
		{
			ApiPath: "/me/mailFolders/Inbox/messages/delta",
			Title:   "跟踪电子邮件更改",
		},
		{
			ApiPath: "/me/outlook/masterCategories",
			Title:   "我的 Outlook 类别",
		},
		{
			ApiPath: "/applications?$count=true",
			Title:   "列出所有应用及计数",
		},
	}
	rand.Seed(time.Now().Unix())
	//随机 [5 ~ 10)轮
	sCount := rand.Intn(5) + 5
	fmt.Printf("本次随机调用次数：%d轮\n", sCount)
	for i := 0; i < sCount; i++ {
		fmt.Printf("\n开始第%d轮请求=================\n", i+1)
		for _, req := range requestList {
			if err := req.ProcessRequest(client); err != nil {
				fmt.Println("GET " + req.Title + " error")
				return err
			}
		}
	}
	return nil
}
