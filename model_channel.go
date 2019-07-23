/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Channel struct {
	// チャンネルUUID
	ChannelId string `json:"channelId,omitempty"`
	// チャンネル名
	Name string `json:"name,omitempty"`
	// UUIDの配列
	Member []string `json:"member,omitempty"`
	// 親の階層のチャンネルUUID
	Parent string `json:"parent,omitempty"`
	// チャンネルトピック
	Topic string `json:"topic,omitempty"`
	// UUIDの配列
	Children []string `json:"children,omitempty"`
	// チャンネルの可視状態
	Visibility bool `json:"visibility,omitempty"`
	// 強制通知チャンネルか
	Force bool `json:"force,omitempty"`
	// プライベートチャンネルか
	Private bool `json:"private,omitempty"`
	// ダイレクトメッセージチャンネルか
	Dm bool `json:"dm,omitempty"`
}
