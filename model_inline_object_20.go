/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineObject20 struct {
	// 通知をつける人のユーザーIDの配列
	On []string `json:"on,omitempty"`
	// 通知をつけない人のユーザーIDの配列
	Off []string `json:"off,omitempty"`
}