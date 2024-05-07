/**
 * Channel Access Token API
 * This document describes Channel Access Token API.
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

package channel_access_token

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

type ChannelAccessTokenAPI struct {
	httpClient *http.Client
	endpoint   *url.URL
	ctx        context.Context
}

// ChannelAccessTokenAPIOption type
type ChannelAccessTokenAPIOption func(*ChannelAccessTokenAPI) error

// New returns a new bot client instance.
func NewChannelAccessTokenAPI(options ...ChannelAccessTokenAPIOption) (*ChannelAccessTokenAPI, error) {
	c := &ChannelAccessTokenAPI{
		httpClient: http.DefaultClient,
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
func (call *ChannelAccessTokenAPI) WithContext(ctx context.Context) *ChannelAccessTokenAPI {
	call.ctx = ctx
	return call
}

func (client *ChannelAccessTokenAPI) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", "LINE-BotSDK-Go/"+linebot.GetVersion())
	if client.ctx != nil {
		req = req.WithContext(client.ctx)
	}
	return client.httpClient.Do(req)
}

func (client *ChannelAccessTokenAPI) Url(endpointPath string) string {
	newPath := path.Join(client.endpoint.Path, endpointPath)
	u := *client.endpoint
	u.Path = newPath
	return u.String()
}

// WithHTTPClient function
func WithHTTPClient(c *http.Client) ChannelAccessTokenAPIOption {
	return func(client *ChannelAccessTokenAPI) error {
		client.httpClient = c
		return nil
	}
}

// WithEndpointClient function
func WithEndpoint(endpoint string) ChannelAccessTokenAPIOption {
	return func(client *ChannelAccessTokenAPI) error {
		u, err := url.ParseRequestURI(endpoint)
		if err != nil {
			return err
		}
		client.endpoint = u
		return nil
	}
}

// GetsAllValidChannelAccessTokenKeyIds
//
// Gets all valid channel access token key IDs.
// Parameters:
//        clientAssertionType             `urn:ietf:params:oauth:client-assertion-type:jwt-bearer`
//        clientAssertion             A JSON Web Token (JWT) (opens new window)the client needs to create and sign with the private key.

// https://developers.line.biz/en/reference/messaging-api/#get-all-valid-channel-access-token-key-ids-v2-1
func (client *ChannelAccessTokenAPI) GetsAllValidChannelAccessTokenKeyIds(

	clientAssertionType string,

	clientAssertion string,

) (*ChannelAccessTokenKeyIdsResponse, error) {
	_, body, error := client.GetsAllValidChannelAccessTokenKeyIdsWithHttpInfo(

		clientAssertionType,

		clientAssertion,
	)
	return body, error
}

// GetsAllValidChannelAccessTokenKeyIds
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Gets all valid channel access token key IDs.
// Parameters:
//        clientAssertionType             `urn:ietf:params:oauth:client-assertion-type:jwt-bearer`
//        clientAssertion             A JSON Web Token (JWT) (opens new window)the client needs to create and sign with the private key.

// https://developers.line.biz/en/reference/messaging-api/#get-all-valid-channel-access-token-key-ids-v2-1
func (client *ChannelAccessTokenAPI) GetsAllValidChannelAccessTokenKeyIdsWithHttpInfo(

	clientAssertionType string,

	clientAssertion string,

) (*http.Response, *ChannelAccessTokenKeyIdsResponse, error) {
	path := "/oauth2/v2.1/tokens/kid"

	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, nil, err
	}

	query := url.Values{}
	query.Add("clientAssertionType", clientAssertionType)
	query.Add("clientAssertion", clientAssertion)

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
	result := ChannelAccessTokenKeyIdsResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// IssueChannelToken
//
// Issue short-lived channel access token
// Parameters:
//        grantType             `client_credentials`
//        clientId             Channel ID.
//        clientSecret             Channel secret.

// https://developers.line.biz/en/reference/messaging-api/#issue-shortlived-channel-access-token
func (client *ChannelAccessTokenAPI) IssueChannelToken(

	grantType string,

	clientId string,

	clientSecret string,

) (*IssueShortLivedChannelAccessTokenResponse, error) {
	_, body, error := client.IssueChannelTokenWithHttpInfo(

		grantType,

		clientId,

		clientSecret,
	)
	return body, error
}

// IssueChannelToken
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Issue short-lived channel access token
// Parameters:
//        grantType             `client_credentials`
//        clientId             Channel ID.
//        clientSecret             Channel secret.

// https://developers.line.biz/en/reference/messaging-api/#issue-shortlived-channel-access-token
func (client *ChannelAccessTokenAPI) IssueChannelTokenWithHttpInfo(

	grantType string,

	clientId string,

	clientSecret string,

) (*http.Response, *IssueShortLivedChannelAccessTokenResponse, error) {
	path := "/v2/oauth/accessToken"

	vs := url.Values{
		"grant_type":    []string{string(grantType)},
		"client_id":     []string{string(clientId)},
		"client_secret": []string{string(clientSecret)},
	}
	buf := vs.Encode()
	body := bytes.NewBufferString(buf)

	req, err := http.NewRequest(http.MethodPost, client.Url(path), body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
	result := IssueShortLivedChannelAccessTokenResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// IssueChannelTokenByJWT
//
// Issues a channel access token that allows you to specify a desired expiration date. This method lets you use JWT assertion for authentication.
// Parameters:
//        grantType             client_credentials
//        clientAssertionType             urn:ietf:params:oauth:client-assertion-type:jwt-bearer
//        clientAssertion             A JSON Web Token the client needs to create and sign with the private key of the Assertion Signing Key.

// https://developers.line.biz/en/reference/messaging-api/#issue-channel-access-token-v2-1
func (client *ChannelAccessTokenAPI) IssueChannelTokenByJWT(

	grantType string,

	clientAssertionType string,

	clientAssertion string,

) (*IssueChannelAccessTokenResponse, error) {
	_, body, error := client.IssueChannelTokenByJWTWithHttpInfo(

		grantType,

		clientAssertionType,

		clientAssertion,
	)
	return body, error
}

// IssueChannelTokenByJWT
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Issues a channel access token that allows you to specify a desired expiration date. This method lets you use JWT assertion for authentication.
// Parameters:
//        grantType             client_credentials
//        clientAssertionType             urn:ietf:params:oauth:client-assertion-type:jwt-bearer
//        clientAssertion             A JSON Web Token the client needs to create and sign with the private key of the Assertion Signing Key.

// https://developers.line.biz/en/reference/messaging-api/#issue-channel-access-token-v2-1
func (client *ChannelAccessTokenAPI) IssueChannelTokenByJWTWithHttpInfo(

	grantType string,

	clientAssertionType string,

	clientAssertion string,

) (*http.Response, *IssueChannelAccessTokenResponse, error) {
	path := "/oauth2/v2.1/token"

	vs := url.Values{
		"grant_type":            []string{string(grantType)},
		"client_assertion_type": []string{string(clientAssertionType)},
		"client_assertion":      []string{string(clientAssertion)},
	}
	buf := vs.Encode()
	body := bytes.NewBufferString(buf)

	req, err := http.NewRequest(http.MethodPost, client.Url(path), body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
	result := IssueChannelAccessTokenResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// IssueStatelessChannelToken
//
// Issues a new stateless channel access token, which doesn't have max active token limit unlike the other token types. The newly issued token is only valid for 15 minutes but can not be revoked until it naturally expires.
// Parameters:
//        grantType             `client_credentials`
//        clientAssertionType             URL-encoded value of `urn:ietf:params:oauth:client-assertion-type:jwt-bearer`
//        clientAssertion             A JSON Web Token the client needs to create and sign with the private key of the Assertion Signing Key.
//        clientId             Channel ID.
//        clientSecret             Channel secret.

// https://developers.line.biz/en/reference/messaging-api/#issue-stateless-channel-access-token
func (client *ChannelAccessTokenAPI) IssueStatelessChannelToken(

	grantType string,

	clientAssertionType string,

	clientAssertion string,

	clientId string,

	clientSecret string,

) (*IssueStatelessChannelAccessTokenResponse, error) {
	_, body, error := client.IssueStatelessChannelTokenWithHttpInfo(

		grantType,

		clientAssertionType,

		clientAssertion,

		clientId,

		clientSecret,
	)
	return body, error
}

// IssueStatelessChannelToken
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Issues a new stateless channel access token, which doesn't have max active token limit unlike the other token types. The newly issued token is only valid for 15 minutes but can not be revoked until it naturally expires.
// Parameters:
//        grantType             `client_credentials`
//        clientAssertionType             URL-encoded value of `urn:ietf:params:oauth:client-assertion-type:jwt-bearer`
//        clientAssertion             A JSON Web Token the client needs to create and sign with the private key of the Assertion Signing Key.
//        clientId             Channel ID.
//        clientSecret             Channel secret.

// https://developers.line.biz/en/reference/messaging-api/#issue-stateless-channel-access-token
func (client *ChannelAccessTokenAPI) IssueStatelessChannelTokenWithHttpInfo(

	grantType string,

	clientAssertionType string,

	clientAssertion string,

	clientId string,

	clientSecret string,

) (*http.Response, *IssueStatelessChannelAccessTokenResponse, error) {
	path := "/oauth2/v3/token"

	vs := url.Values{
		"grant_type":            []string{string(grantType)},
		"client_assertion_type": []string{string(clientAssertionType)},
		"client_assertion":      []string{string(clientAssertion)},
		"client_id":             []string{string(clientId)},
		"client_secret":         []string{string(clientSecret)},
	}
	buf := vs.Encode()
	body := bytes.NewBufferString(buf)

	req, err := http.NewRequest(http.MethodPost, client.Url(path), body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
	result := IssueStatelessChannelAccessTokenResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// RevokeChannelToken
//
// Revoke short-lived or long-lived channel access token
// Parameters:
//        accessToken             Channel access token

// https://developers.line.biz/en/reference/messaging-api/#revoke-longlived-or-shortlived-channel-access-token
func (client *ChannelAccessTokenAPI) RevokeChannelToken(

	accessToken string,

) (struct{}, error) {
	_, body, error := client.RevokeChannelTokenWithHttpInfo(

		accessToken,
	)
	return body, error
}

// RevokeChannelToken
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Revoke short-lived or long-lived channel access token
// Parameters:
//        accessToken             Channel access token

// https://developers.line.biz/en/reference/messaging-api/#revoke-longlived-or-shortlived-channel-access-token
func (client *ChannelAccessTokenAPI) RevokeChannelTokenWithHttpInfo(

	accessToken string,

) (*http.Response, struct{}, error) {
	path := "/v2/oauth/revoke"

	vs := url.Values{
		"access_token": []string{string(accessToken)},
	}
	buf := vs.Encode()
	body := bytes.NewBufferString(buf)

	req, err := http.NewRequest(http.MethodPost, client.Url(path), body)
	if err != nil {
		return nil, struct{}{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

// RevokeChannelTokenByJWT
//
// Revoke channel access token v2.1
// Parameters:
//        clientId             Channel ID
//        clientSecret             Channel Secret
//        accessToken             Channel access token

// https://developers.line.biz/en/reference/messaging-api/#revoke-channel-access-token-v2-1
func (client *ChannelAccessTokenAPI) RevokeChannelTokenByJWT(

	clientId string,

	clientSecret string,

	accessToken string,

) (struct{}, error) {
	_, body, error := client.RevokeChannelTokenByJWTWithHttpInfo(

		clientId,

		clientSecret,

		accessToken,
	)
	return body, error
}

// RevokeChannelTokenByJWT
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Revoke channel access token v2.1
// Parameters:
//        clientId             Channel ID
//        clientSecret             Channel Secret
//        accessToken             Channel access token

// https://developers.line.biz/en/reference/messaging-api/#revoke-channel-access-token-v2-1
func (client *ChannelAccessTokenAPI) RevokeChannelTokenByJWTWithHttpInfo(

	clientId string,

	clientSecret string,

	accessToken string,

) (*http.Response, struct{}, error) {
	path := "/oauth2/v2.1/revoke"

	vs := url.Values{
		"client_id":     []string{string(clientId)},
		"client_secret": []string{string(clientSecret)},
		"access_token":  []string{string(accessToken)},
	}
	buf := vs.Encode()
	body := bytes.NewBufferString(buf)

	req, err := http.NewRequest(http.MethodPost, client.Url(path), body)
	if err != nil {
		return nil, struct{}{}, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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

// VerifyChannelToken
//
// Verify the validity of short-lived and long-lived channel access tokens
// Parameters:
//        accessToken             A short-lived or long-lived channel access token.

// https://developers.line.biz/en/reference/messaging-api/#verify-channel-access-token
func (client *ChannelAccessTokenAPI) VerifyChannelToken(

	accessToken string,

) (*VerifyChannelAccessTokenResponse, error) {
	_, body, error := client.VerifyChannelTokenWithHttpInfo(

		accessToken,
	)
	return body, error
}

// VerifyChannelToken
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// Verify the validity of short-lived and long-lived channel access tokens
// Parameters:
//        accessToken             A short-lived or long-lived channel access token.

// https://developers.line.biz/en/reference/messaging-api/#verify-channel-access-token
func (client *ChannelAccessTokenAPI) VerifyChannelTokenWithHttpInfo(

	accessToken string,

) (*http.Response, *VerifyChannelAccessTokenResponse, error) {
	path := "/v2/oauth/verify"

	vs := url.Values{
		"access_token": []string{string(accessToken)},
	}
	buf := vs.Encode()
	body := bytes.NewBufferString(buf)

	req, err := http.NewRequest(http.MethodPost, client.Url(path), body)
	if err != nil {
		return nil, nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
	result := VerifyChannelAccessTokenResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}

// VerifyChannelTokenByJWT
//
// You can verify whether a Channel access token with a user-specified expiration (Channel Access Token v2.1) is valid.
// Parameters:
//        accessToken             Channel access token with a user-specified expiration (Channel Access Token v2.1).

// https://developers.line.biz/en/reference/messaging-api/#verify-channel-access-token-v2-1
func (client *ChannelAccessTokenAPI) VerifyChannelTokenByJWT(

	accessToken string,

) (*VerifyChannelAccessTokenResponse, error) {
	_, body, error := client.VerifyChannelTokenByJWTWithHttpInfo(

		accessToken,
	)
	return body, error
}

// VerifyChannelTokenByJWT
// If you want to take advantage of the HTTPResponse object for status codes and headers, use this signature.
//
// You can verify whether a Channel access token with a user-specified expiration (Channel Access Token v2.1) is valid.
// Parameters:
//        accessToken             Channel access token with a user-specified expiration (Channel Access Token v2.1).

// https://developers.line.biz/en/reference/messaging-api/#verify-channel-access-token-v2-1
func (client *ChannelAccessTokenAPI) VerifyChannelTokenByJWTWithHttpInfo(

	accessToken string,

) (*http.Response, *VerifyChannelAccessTokenResponse, error) {
	path := "/oauth2/v2.1/verify"

	req, err := http.NewRequest(http.MethodGet, client.Url(path), nil)
	if err != nil {
		return nil, nil, err
	}

	query := url.Values{}
	query.Add("accessToken", accessToken)

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
	result := VerifyChannelAccessTokenResponse{}
	if err := decoder.Decode(&result); err != nil {
		return res, nil, fmt.Errorf("failed to decode JSON: %w", err)
	}
	return res, &result, nil

}
