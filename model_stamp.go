/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)

type Stamp struct {
	// スタンプUUID
	Id string `json:"id,omitempty"`
	// スタンプ名
	Name string `json:"name,omitempty"`
	// スタンプ作成者UUID
	CreatorId string `json:"creatorId,omitempty"`
	// スタンプファイルUUID
	FileId string `json:"fileId,omitempty"`
	// スタンプ作成日時
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// スタンプ更新日時
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}