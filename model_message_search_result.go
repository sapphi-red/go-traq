/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// MessageSearchResult メッセージ検索結果
type MessageSearchResult struct {
	// 検索にヒットしたメッセージ件数
	TotalHits int64 `json:"totalHits,omitempty"`
	// 検索にヒットしたメッセージの配列
	Hits []Message `json:"hits,omitempty"`
}
