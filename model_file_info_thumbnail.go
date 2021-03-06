/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// FileInfoThumbnail サムネイル情報 サムネイルが存在しない場合はnullになります Deprecated: thumbnailsを参照してください
type FileInfoThumbnail struct {
	// MIMEタイプ
	Mime string `json:"mime"`
	// サムネイル幅
	Width int32 `json:"width,omitempty"`
	// サムネイル高さ
	Height int32 `json:"height,omitempty"`
}
