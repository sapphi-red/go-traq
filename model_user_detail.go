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

// UserDetail ユーザー詳細情報
type UserDetail struct {
	// ユーザーUUID
	Id    string           `json:"id"`
	State UserAccountState `json:"state"`
	// BOTかどうか
	Bot bool `json:"bot"`
	// アイコンファイルUUID
	IconFileId string `json:"iconFileId"`
	// ユーザー表示名
	DisplayName string `json:"displayName"`
	// ユーザー名
	Name string `json:"name"`
	// Twitter ID
	TwitterId string `json:"twitterId"`
	// 最終オンライン日時
	LastOnline *time.Time `json:"lastOnline"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
	// タグリスト
	Tags []UserTag `json:"tags"`
	// 所属グループのUUIDの配列
	Groups []string `json:"groups"`
	// 自己紹介(biography)
	Bio string `json:"bio"`
	// ホームチャンネル
	HomeChannel *string `json:"homeChannel"`
}
