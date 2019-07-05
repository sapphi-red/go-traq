/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject31 struct {
	// クライアント名(1-32文字)
	Name string `json:"name"`
	// クライアントの説明
	Description string `json:"description"`
	// リダイレクト先のURI
	RedirectUri string `json:"redirectUri"`
	// 要求するスコープ(必ず１つ以上)
	Scopes []string `json:"scopes"`
}