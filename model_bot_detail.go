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

// BotDetail BOT詳細情報
type BotDetail struct {
	// BOT UUID
	Id string `json:"id"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	State     BotState  `json:"state"`
	// BOTが購読しているイベントの配列
	SubscribeEvents []string `json:"subscribeEvents"`
	// BOT開発者UUID
	DeveloperId string `json:"developerId"`
	// 説明
	Description string `json:"description"`
	// BOTユーザーUUID
	BotUserId string    `json:"botUserId"`
	Tokens    BotTokens `json:"tokens"`
	// BOTサーバーエンドポイント
	Endpoint string `json:"endpoint"`
	// 特権BOTかどうか
	Privileged bool `json:"privileged"`
	// BOTが参加しているチャンネルのUUID配列
	Channels []string `json:"channels"`
}
