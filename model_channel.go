/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// Channel チャンネル
type Channel struct {
	// チャンネルUUID
	Id string `json:"id"`
	// 親チャンネルUUID
	ParentId *string `json:"parentId"`
	// チャンネルがアーカイブされているかどうか
	Archived bool `json:"archived"`
	// 強制通知チャンネルかどうか
	Force bool `json:"force"`
	// チャンネルトピック
	Topic string `json:"topic"`
	// チャンネル名
	Name string `json:"name"`
	// 子チャンネルのUUID配列
	Children []string `json:"children"`
}
