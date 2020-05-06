/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PostUserGroupRequest ユーザーグループ作成リクエスト
type PostUserGroupRequest struct {
	// グループ名
	Name string `json:"name"`
	// 説明
	Description string `json:"description"`
	// グループタイプ
	Type string `json:"type"`
}