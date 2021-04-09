package api_lib

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
)

func (tokenInfo *TokenInfo) RefreshNew() (*TokenInfo, error) {
	requestURI := "https://login.microsoftonline.com/common/oauth2/v2.0/token"
	client := resty.New()

	resp, err := client.R().
		SetFormData(map[string]string{
			"client_id":     tokenInfo.ClientId,
			"client_secret": tokenInfo.ClientSecret,
			"scope":         "offline_access user.read mail.read",
			"grant_type":    "refresh_token",
			"refresh_token": tokenInfo.RefreshToken,
		}).
		Post(requestURI)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		return nil, errors.New(resp.String())
	}
	responseData := resp.Body()
	newTokenInfo := &TokenInfo{
		ClientId:     tokenInfo.ClientId,
		ClientSecret: tokenInfo.ClientSecret,
	}
	if err := json.Unmarshal(responseData, newTokenInfo); err != nil {
		return nil, err
	}
	return newTokenInfo, nil
}
