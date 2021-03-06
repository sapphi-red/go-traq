/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// ChildCreatedEvent 子チャンネル作成イベント
type ChildCreatedEvent struct {
	// 作成者UUID
	UserId string `json:"userId"`
	// チャンネルUUID
	ChannelId string `json:"channelId"`
}
