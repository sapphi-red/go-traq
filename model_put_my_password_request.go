/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PutMyPasswordRequest パスワード変更リクエスト
type PutMyPasswordRequest struct {
	// 現在のパスワード
	Password string `json:"password"`
	// 新しいパスワード
	NewPassword string `json:"newPassword"`
}
