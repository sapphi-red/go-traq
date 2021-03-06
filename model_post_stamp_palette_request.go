/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// PostStampPaletteRequest スタンプパレット作成リクエスト
type PostStampPaletteRequest struct {
	// パレット内のスタンプのUUID配列
	Stamps []string `json:"stamps"`
	// パレット名
	Name string `json:"name"`
	// 説明
	Description string `json:"description"`
}
