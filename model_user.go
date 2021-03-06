/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

import (
	"time"
)

// User ユーザー情報
type User struct {
	// ユーザーUUID
	Id string `json:"id"`
	// ユーザー名
	Name string `json:"name"`
	// ユーザー表示名
	DisplayName string `json:"displayName"`
	// アイコンファイルUUID
	IconFileId string `json:"iconFileId"`
	// BOTかどうか
	Bot   bool             `json:"bot"`
	State UserAccountState `json:"state"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}
