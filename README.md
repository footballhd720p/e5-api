# e5-api
用于调用Microsoft 365 api的测试工具



需要设置以下几个`Secret`

| 键名              | 说明                                                         |
| ----------------- | ------------------------------------------------------------ |
| APP_CLIENT_ID     | Azure的应用程序(客户端) ID                                   |
| APP_CLIENT_SECRET | Azure的应用 - 证书和密码 - 客户端密码 - 配置值               |
| APP_TOKEN         | 用户授权登录得到的access_token                               |
| APP_R_TOKEN       | 用户授权登录得到的refresh_token                              |
| REPO_ACCESS_TOKEN | [Github token](https://github.com/settings/tokens) 页面添加一个新的token，勾选repo权限(此token可以用于更新`access_token` , `refresh_token`数据) |

`access_token`, `refresh_token`可以使用该工具获取 [https://github.com/liuguangw/e5-worker](https://github.com/liuguangw/e5-worker)



> `config.json`是本地跑api时才需要修改,如果使用GitHub action, 无需修改此文件。