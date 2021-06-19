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

// FileInfo ファイル情報
type FileInfo struct {
	// ファイルUUID
	Id string `json:"id"`
	// ファイル名
	Name string `json:"name"`
	// MIMEタイプ
	Mime string `json:"mime"`
	// ファイルサイズ
	Size int64 `json:"size"`
	// MD5ハッシュ
	Md5 string `json:"md5"`
	// アニメーション画像かどうか
	IsAnimatedImage bool `json:"isAnimatedImage"`
	// アップロード日時
	CreatedAt  time.Time          `json:"createdAt"`
	Thumbnails []ThumbnailInfo    `json:"thumbnails"`
	Thumbnail  *FileInfoThumbnail `json:"thumbnail"`
	// 属しているチャンネルUUID
	ChannelId *string `json:"channelId"`
	// アップロード者UUID
	UploaderId *string `json:"uploaderId"`
}
