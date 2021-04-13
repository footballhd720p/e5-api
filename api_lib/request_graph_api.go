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
	Handler func(client *resty.Client, req *ApiRequest) error
}

//普通的GET请求
func defaultRequestHandler(client *resty.Client, req *ApiRequest) error {
	resp, err := client.R().Get(req.ApiPath)
	if err != nil {
		fmt.Println("GET " + req.Title + " error")
		return err
	}
	statusCode := resp.StatusCode()
	if statusCode != 200 {
		fmt.Println("GET " + req.Title + " error")
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
			ApiPath: "/me/drive/following",
			Title:   "列出onedrive关注的项",
		},
		{
			ApiPath: "/me/messages",
			Title:   "我的邮件",
		},
		{
			ApiPath: "/me/mailFolders",
			Title:   "列出邮件文件夹",
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
		{
			ApiPath: "/servicePrincipals",
			Title:   "获取服务主体列表",
		},
	}
	timeNow := time.Now()
	fileName := timeNow.Format("2006/01/02") + ".txt"
	requestList = append(requestList, &ApiRequest{
		ApiPath: "/me/drive/root:/e5/log-" + fileName + ":/content",
		Title:   "写入Onedrive文件",
		Handler: writeOneDriveHandler,
	})
	rand.Seed(timeNow.Unix())
	//随机 [5 ~ 10)轮
	sCount := rand.Intn(5) + 5
	fmt.Printf("本次随机调用次数：%d轮\n", sCount)
	for i := 0; i < sCount; i++ {
		fmt.Printf("\n开始第%d轮请求=================\n", i+1)
		//打乱请求顺序
		rand.Shuffle(len(requestList), func(i, j int) {
			requestList[i], requestList[j] = requestList[j], requestList[i]
		})
		//每次请求3/4的api
		for _, req := range requestList[0 : len(requestList)*3/4] {
			reqHandler := defaultRequestHandler
			if req.Handler != nil {
				reqHandler = req.Handler
			}
			if err := reqHandler(client, req); err != nil {
				fmt.Println("GET " + req.Title + " error")
				return err
			}
		}
	}
	return nil
}
