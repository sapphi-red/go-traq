/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject32 struct {
	// クライアント名(1-32文字)
	Name string `json:"name,omitempty"`
	// 説明
	Description string `json:"description,omitempty"`
	// リダイレクト先のURI
	RedirectUri string `json:"redirectUri,omitempty"`
}
