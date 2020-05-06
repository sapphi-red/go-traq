/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PatchUserRequest ユーザー情報編集リクエスト
type PatchUserRequest struct {
	// 新しい表示名
	DisplayName string `json:"displayName,omitempty"`
	// TwitterID
	TwitterId string           `json:"twitterId,omitempty"`
	State     UserAccountState `json:"state,omitempty"`
	// ユーザーロール
	Role string `json:"role,omitempty"`
}