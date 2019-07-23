/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Linger please
var (
	_ context.Context
)

type AuthorizationApiService service

/*
AuthorizationApiService
OAuth2 認可承諾
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param submit 承諾する場合は\\\"approve\\\"
*/
func (a *AuthorizationApiService) Oauth2AuthorizeDecidePost(ctx context.Context, submit string) (*http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/authorize/decide"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("submit", parameterToString(submit, ""))
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
AuthorizationApiService
OAuth2 認可エンドポイント
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *Oauth2AuthorizeGetOpts - Optional Parameters:
 * @param "ResponseType" (optional.Interface of OAuth2ResponseType) -
 * @param "ClientId" (optional.String) -
 * @param "RedirectUri" (optional.String) -
 * @param "Scope" (optional.String) -
 * @param "State" (optional.String) -
 * @param "CodeChallenge" (optional.String) -
 * @param "CodeChallengeMethod" (optional.String) -
 * @param "Nonce" (optional.String) -
 * @param "Prompt" (optional.Interface of OAuth2Prompt) -
*/

type Oauth2AuthorizeGetOpts struct {
	ResponseType        optional.Interface
	ClientId            optional.String
	RedirectUri         optional.String
	Scope               optional.String
	State               optional.String
	CodeChallenge       optional.String
	CodeChallengeMethod optional.String
	Nonce               optional.String
	Prompt              optional.Interface
}

func (a *AuthorizationApiService) Oauth2AuthorizeGet(ctx context.Context, localVarOptionals *Oauth2AuthorizeGetOpts) (*http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ResponseType.IsSet() {
		localVarQueryParams.Add("response_type", parameterToString(localVarOptionals.ResponseType.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientId.IsSet() {
		localVarQueryParams.Add("client_id", parameterToString(localVarOptionals.ClientId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarQueryParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarQueryParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarQueryParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallenge.IsSet() {
		localVarQueryParams.Add("code_challenge", parameterToString(localVarOptionals.CodeChallenge.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallengeMethod.IsSet() {
		localVarQueryParams.Add("code_challenge_method", parameterToString(localVarOptionals.CodeChallengeMethod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nonce.IsSet() {
		localVarQueryParams.Add("nonce", parameterToString(localVarOptionals.Nonce.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prompt.IsSet() {
		localVarQueryParams.Add("prompt", parameterToString(localVarOptionals.Prompt.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
AuthorizationApiService
OAuth2 認可エンドポイント
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *Oauth2AuthorizePostOpts - Optional Parameters:
 * @param "ResponseType" (optional.Interface of OAuth2ResponseType) -
 * @param "ClientId" (optional.String) -
 * @param "RedirectUri" (optional.String) -
 * @param "Scope" (optional.String) -
 * @param "State" (optional.String) -
 * @param "CodeChallenge" (optional.String) -
 * @param "CodeChallengeMethod" (optional.String) -
 * @param "Nonce" (optional.String) -
 * @param "Prompt" (optional.Interface of OAuth2Prompt) -
*/

type Oauth2AuthorizePostOpts struct {
	ResponseType        optional.Interface
	ClientId            optional.String
	RedirectUri         optional.String
	Scope               optional.String
	State               optional.String
	CodeChallenge       optional.String
	CodeChallengeMethod optional.String
	Nonce               optional.String
	Prompt              optional.Interface
}

func (a *AuthorizationApiService) Oauth2AuthorizePost(ctx context.Context, localVarOptionals *Oauth2AuthorizePostOpts) (*http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/authorize"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if localVarOptionals != nil && localVarOptionals.ResponseType.IsSet() {
		paramJson, err := parameterToJson(localVarOptionals.ResponseType.Value())
		if err != nil {
			return nil, err
		}
		localVarFormParams.Add("response_type", paramJson)
	}
	if localVarOptionals != nil && localVarOptionals.ClientId.IsSet() {
		localVarFormParams.Add("client_id", parameterToString(localVarOptionals.ClientId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarFormParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarFormParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.State.IsSet() {
		localVarFormParams.Add("state", parameterToString(localVarOptionals.State.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallenge.IsSet() {
		localVarFormParams.Add("code_challenge", parameterToString(localVarOptionals.CodeChallenge.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeChallengeMethod.IsSet() {
		localVarFormParams.Add("code_challenge_method", parameterToString(localVarOptionals.CodeChallengeMethod.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Nonce.IsSet() {
		localVarFormParams.Add("nonce", parameterToString(localVarOptionals.Nonce.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Prompt.IsSet() {
		paramJson, err := parameterToJson(localVarOptionals.Prompt.Value())
		if err != nil {
			return nil, err
		}
		localVarFormParams.Add("prompt", paramJson)
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
AuthorizationApiService
OAuth2 トークンエンドポイント
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param grantType
 * @param optional nil or *Oauth2TokenPostOpts - Optional Parameters:
 * @param "Code" (optional.String) -
 * @param "RedirectUri" (optional.String) -
 * @param "ClientId" (optional.String) -
 * @param "CodeVerifier" (optional.String) -
 * @param "Username" (optional.String) -
 * @param "Password" (optional.String) -
 * @param "Scope" (optional.String) -
 * @param "RefreshToken" (optional.String) -
 * @param "ClientSecret" (optional.String) -
@return OAuth2Token
*/

type Oauth2TokenPostOpts struct {
	Code         optional.String
	RedirectUri  optional.String
	ClientId     optional.String
	CodeVerifier optional.String
	Username     optional.String
	Password     optional.String
	Scope        optional.String
	RefreshToken optional.String
	ClientSecret optional.String
}

func (a *AuthorizationApiService) Oauth2TokenPost(ctx context.Context, grantType string, localVarOptionals *Oauth2TokenPostOpts) (OAuth2Token, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  OAuth2Token
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/oauth2/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarFormParams.Add("grant_type", parameterToString(grantType, ""))
	if localVarOptionals != nil && localVarOptionals.Code.IsSet() {
		localVarFormParams.Add("code", parameterToString(localVarOptionals.Code.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RedirectUri.IsSet() {
		localVarFormParams.Add("redirect_uri", parameterToString(localVarOptionals.RedirectUri.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientId.IsSet() {
		localVarFormParams.Add("client_id", parameterToString(localVarOptionals.ClientId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CodeVerifier.IsSet() {
		localVarFormParams.Add("code_verifier", parameterToString(localVarOptionals.CodeVerifier.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Username.IsSet() {
		localVarFormParams.Add("username", parameterToString(localVarOptionals.Username.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Password.IsSet() {
		localVarFormParams.Add("password", parameterToString(localVarOptionals.Password.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Scope.IsSet() {
		localVarFormParams.Add("scope", parameterToString(localVarOptionals.Scope.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RefreshToken.IsSet() {
		localVarFormParams.Add("refresh_token", parameterToString(localVarOptionals.RefreshToken.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ClientSecret.IsSet() {
		localVarFormParams.Add("client_secret", parameterToString(localVarOptionals.ClientSecret.Value(), ""))
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v OAuth2Token
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
