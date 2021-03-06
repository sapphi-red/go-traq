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

// Webhook Webhook情報
type Webhook struct {
	// WebhookUUID
	Id string `json:"id"`
	// WebhookユーザーUUID
	BotUserId string `json:"botUserId"`
	// Webhookユーザー表示名
	DisplayName string `json:"displayName"`
	// 説明
	Description string `json:"description"`
	// セキュアWebhookかどうか
	Secure bool `json:"secure"`
	// デフォルトの投稿先チャンネルUUID
	ChannelId string `json:"channelId"`
	// オーナーUUID
	OwnerId string `json:"ownerId"`
	// 作成日時
	CreatedAt time.Time `json:"createdAt"`
	// 更新日時
	UpdatedAt time.Time `json:"updatedAt"`
}
