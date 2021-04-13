package api_lib

import (
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"time"
)

//写入OneDrive的handler
func writeOneDriveHandler(client *resty.Client, req *ApiRequest) error {
	timeNow := time.Now()
	bodyStr := "E5 API运行于" + timeNow.Format(TimeLayout)
	requestMethodType := "PUT"
	resp, err := client.R().
		SetHeader("Content-Type", "text/plain").
		SetBody(bodyStr).
		Put(req.ApiPath)
	if err != nil {
		fmt.Println(requestMethodType + " " + req.Title + " error")
		return err
	}
	statusCode := resp.StatusCode()
	if statusCode != 200 && statusCode != 201 {
		fmt.Println(requestMethodType + " " + req.Title + " error")
		return errors.New(resp.String())
	}
	fmt.Println(requestMethodType + " " + req.Title + " OK")
	//fmt.Println(resp.String())
	return nil
}
