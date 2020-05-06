/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// PatchChannelRequest チャンネル情報変更リクエスト
type PatchChannelRequest struct {
	// チャンネル名
	Name string `json:"name,omitempty"`
	// アーカイブされているかどうか
	Archived bool `json:"archived,omitempty"`
	// 強制通知チャンネルかどうか
	Force bool `json:"force,omitempty"`
	// 親チャンネルUUID
	Parent string `json:"parent,omitempty"`
}