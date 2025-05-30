/**
 * LINE Messaging API
 * This document describes LINE Messaging API.
 *
 * The version of the OpenAPI document: 0.0.1
 *
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

//go:generate python3 ../../generate-code.py

package module

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

type LineModuleAPI struct {
	httpClient   *http.Client
	endpoint     *url.URL
	channelToken string
	ctx          context.Context
}

// LineModuleAPIOption type
type LineModuleAPIOption func(*LineModuleAPI) error

// New returns a new bot client instance.
func NewLineModuleAPI(channelToken string, options ...LineModuleAPIOption) (*LineModuleAPI, error) {
	if channelToken == "" {
		return nil, errors.New("missing channel access token")
	}

	c := &LineModuleAPI{
		channelToken: channelToken,
		httpClient:   http.DefaultClient,
	}

	u, err := url.ParseRequestURI("https://api.line.me")
	if err != nil {
		return nil, err
	}
	c.endpoint = u

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}
	return c, nil
}

// WithContext method
func (call *LineModuleAPI) WithContext(ctx context.Context) *LineModuleAPI {
	call.ctx = ctx
	return call
}

func (client *LineModuleAPI) Do(req *http.Request) (*http.Response, error) {
	if client.channelToken != "" {
		req.Header.Set("Authorization", "Bearer "+client.channelToken)
	}
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+linebot.GetVersion())
	if client.ctx != nil {
		req = req.WithContext(client.ctx)
	}
	return client.httpClient.Do(req)
}

func (client *LineModuleAPI) Url(endpointPath string) string {
	newPath := path.Join(client.endpoint.Path, endpointPath)
	u := *client.endpoint
	u.Path = newPath
	return u.String()
}

// WithHTTPClient function
func WithHTTPClient(c *http.Client) LineModuleAPIOption {
	return func(client *LineModuleAPI) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpoint function
func WithEndpoint(endpoint string) LineModuleAPIOption {
	return func(client *LineModuleAPI) error {
		u, err := url.ParseRequestURI(endpoint)
		if err != nil {
			return err
		}
		client.endpoint = u
		return nil
	}
}

// AcquireChatControl
//
// If the Standby Channel wants to take the initiative (Chat Control), it calls the Acquire Control API. The channel that was previously an Active Channel will automatically switch to a Standby Channel.
// Parameters:
//        chatId             The `userId`, `roomId`, or `groupId`
//        acquireChatControlRequest

// https://developers.line.biz/en/reference/partner-docs/#acquire-control-api
func (client *LineModuleAPI) AcquireChatControl(

	chatId string,

	acquireChatControlRequest *AcquireChatControlRequest,

) (struct{}, error) {
	_, body, error := client.AcquireChatControlWithHttpInfo(

		chatId,

		acquireChatControlRequest,
	)
	return body, error
}

// AcquireChatControl
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// If the Standby Channel wants to take the initiative (Chat Control), it calls the Acquire Control API. The channel that was previously an Active Channel will automatically switch to a Standby Channel.
// Parameters:
//        chatId             The `userId`, `roomId`, or `groupId`
//        acquireChatControlRequest

// https://developers.line.biz/en/reference/partner-docs/#acquire-control-api
func (client *LineModuleAPI) AcquireChatControlWithHttpInfo(

	chatId string,

	acquireChatControlRequest *AcquireChatControlRequest,

) (*http.Response, struct{}, error) {
	path := "/v2/bot/chat/{chatId}/control/acquire"

	path = strings.Replace(path, "{chatId}", chatId, -1)

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(acquireChatControlRequest); err != nil {
		return nil, struct{}{}, err
	}
	req, err := http.NewRequest(http.MethodPost, client.Url(path), &buf)
	if err != nil {
		return nil, struct{}{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)

	if err != nil {
		return res, struct{}{}, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, struct{}{}, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, struct{}{}, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	return res, struct{}{}, nil

}

// DetachModule
//
// The module channel admin calls the Detach API to detach the module channel from a LINE Official Account.
// Parameters:
//        detachModuleRequest

// https://developers.line.biz/en/reference/partner-docs/#unlink-detach-module-channel-by-operation-mc-admin
func (client *LineModuleAPI) DetachModule(

	detachModuleRequest *DetachModuleRequest,

) (struct{}, error) {
	_, body, error := client.DetachModuleWithHttpInfo(

		detachModuleRequest,
	)
	return body, error
}

// DetachModule
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// The module channel admin calls the Detach API to detach the module channel from a LINE Official Account.
// Parameters:
//        detachModuleRequest

// https://developers.line.biz/en/reference/partner-docs/#unlink-detach-module-channel-by-operation-mc-admin
func (client *LineModuleAPI) DetachModuleWithHttpInfo(

	detachModuleRequest *DetachModuleRequest,

) (*http.Response, struct{}, error) {
	path := "/v2/bot/channel/detach"

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(detachModuleRequest); err != nil {
		return nil, struct{}{}, err
	}
	req, err := http.NewRequest(http.MethodPost, client.Url(path), &buf)
	if err != nil {
		return nil, struct{}{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := client.Do(req)

	if err != nil {
		return res, struct{}{}, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, struct{}{}, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, struct{}{}, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	return res, struct{}{}, nil

}

// GetModules
//
// Gets a list of basic information about the bots of multiple LINE Official Accounts that have attached module channels.
// Parameters:
//        start             Value of the continuation token found in the next property of the JSON object returned in the response. If you can't get all basic information about the bots in one request, include this parameter to get the remaining array.
//        limit             Specify the maximum number of bots that you get basic information from. The default value is 100. Max value: 100

// https://developers.line.biz/en/reference/partner-docs/#get-multiple-bot-info-api
func (client *LineModuleAPI) GetModules(

	start string,

	limit int32,

) (*GetModulesResponse, error) {
	_, body, error := client.GetModulesWithHttpInfo(

		start,

		limit,
	)
	return body, error
}

// GetModules
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Gets a list of basic information about the bots of multiple LINE Official Accounts that have attached module channels.
// Parameters:
//        start             Value of the continuation token found in the next property of the JSON object returned in the response. If you can't get all basic information about the bots in one request, include this parameter to get the remaining array.
//        limit             Specify the maximum number of bots that you get basic information from. The default value is 100. Max value: 100

// https://developers.line.biz/en/reference/partner-docs/#get-multiple-bot-info-api
func (client *LineModuleAPI) GetModulesWithHttpInfo(

	start string,

	limit int32,

) (*http.Response, *GetModulesResponse, error) {
	path := "/v2/bot/list"

	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, nil, err
	}

	query := url.Values{}
	if start != "" {
		query.Add("start", start)
	}
	query.Add("limit", strconv.FormatInt(int64(limit), 10))

	req.URL.RawQuery = query.Encode()

	res, err := client.Do(req)

	if err != nil {
		return res, nil, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, nil, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, nil, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	result := GetModulesResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// ReleaseChatControl
//
// To return the initiative (Chat Control) of Active Channel to Primary Channel, call the Release Control API.
// Parameters:
//        chatId             The `userId`, `roomId`, or `groupId`

// https://developers.line.biz/en/reference/partner-docs/#release-control-api
func (client *LineModuleAPI) ReleaseChatControl(

	chatId string,

) (struct{}, error) {
	_, body, error := client.ReleaseChatControlWithHttpInfo(

		chatId,
	)
	return body, error
}

// ReleaseChatControl
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// To return the initiative (Chat Control) of Active Channel to Primary Channel, call the Release Control API.
// Parameters:
//        chatId             The `userId`, `roomId`, or `groupId`

// https://developers.line.biz/en/reference/partner-docs/#release-control-api
func (client *LineModuleAPI) ReleaseChatControlWithHttpInfo(

	chatId string,

) (*http.Response, struct{}, error) {
	path := "/v2/bot/chat/{chatId}/control/release"

	path = strings.Replace(path, "{chatId}", chatId, -1)

	req, err := http.NewRequest(http.MethodPost, client.Url(path), nil)
	if err != nil {
		return nil, struct{}{}, err
	}

	res, err := client.Do(req)

	if err != nil {
		return res, struct{}{}, err
	}

	if res.StatusCode/100 != 2 {
		bodyBytes, err := io.ReadAll(res.Body)
		bodyReader := bytes.NewReader(bodyBytes)
		if err != nil {
			return res, struct{}{}, fmt.Errorf("failed to read response body: %w", err)
		}
		res.Body = io.NopCloser(bodyReader)
		return res, struct{}{}, fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	defer res.Body.Close()

	return res, struct{}{}, nil

}
