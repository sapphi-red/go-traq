/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PatchGroupMemberRequest ユーザーグループメンバー編集リクエスト
type PatchGroupMemberRequest struct {
	// ユーザーの役割
	Role string `json:"role"`
}
