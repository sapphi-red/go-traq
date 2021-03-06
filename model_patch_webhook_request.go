/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PatchWebhookRequest Webhook情報変更リクエスト
type PatchWebhookRequest struct {
	// Webhookユーザー表示名
	Name string `json:"name,omitempty"`
	// 説明
	Description string `json:"description,omitempty"`
	// デフォルトの投稿先チャンネルUUID
	ChannelId string `json:"channelId,omitempty"`
	// Webhookシークレット
	Secret string `json:"secret,omitempty"`
	// 移譲先のユーザーUUID
	OwnerId string `json:"ownerId,omitempty"`
}
