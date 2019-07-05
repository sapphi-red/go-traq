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

type Message struct {
	// メッセージUUID
	MessageId string `json:"messageId,omitempty"`
	// 投稿者UUID
	UserId string `json:"userId,omitempty"`
	// 投稿先チャンネルUUID
	ParentChannelId string `json:"parentChannelId,omitempty"`
	// メッセージ本体
	Content string `json:"content,omitempty"`
	// メッセージ投稿日時
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// メッセージ更新日時
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// ピン留めされているかどうか
	Pin bool `json:"pin,omitempty"`
	// 自分が通報しているかどうか
	Reported bool `json:"reported,omitempty"`
	// メッセージスタンプ配列
	StampList []MessageStamp `json:"stampList,omitempty"`
}