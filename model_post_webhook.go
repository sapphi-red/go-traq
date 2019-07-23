/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostWebhook struct {
	// webhookユーザーの表示名
	Name string `json:"name"`
	// webhookの説明
	Description string `json:"description"`
	// デフォルトの投稿先チャンネル(パブリックチャンネルのみ)
	ChannelId string `json:"channelId"`
	// webhookシークレット
	Secret string `json:"secret,omitempty"`
}