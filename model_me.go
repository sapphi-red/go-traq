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
	"time"
)

type Me struct {
	// ユーザーUUID
	UserId string `json:"userId,omitempty"`
	// traQ ID
	Name string `json:"name,omitempty"`
	// 表示名
	DisplayName string `json:"displayName,omitempty"`
	// アイコンファイルUUID
	IconFileId string `json:"iconFileId,omitempty"`
	// BOTかどうか
	Bot bool `json:"bot,omitempty"`
	// ツイッターID
	TwitterId string `json:"twitterId,omitempty"`
	// 最終オンライン日時(オンラインに１度もなってない場合はnull)
	LastOnline time.Time `json:"lastOnline,omitempty"`
	// 現在オンラインかどうか
	IsOnline bool `json:"isOnline,omitempty"`
	// アカウントが停止中かどうか
	Suspended bool `json:"suspended,omitempty"`
	// アカウントの状態 (0:停止,1:有効,2:一時停止)
	AccountStatus int32 `json:"accountStatus,omitempty"`
	// ユーザーロール
	Role string `json:"role,omitempty"`
	// 所有している権限の配列
	Permissions []string `json:"permissions,omitempty"`
}
