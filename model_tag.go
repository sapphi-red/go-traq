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

type Tag struct {
	// タグUUID
	TagId string `json:"tagId,omitempty"`
	// タグ文字列
	Tag string `json:"tag,omitempty"`
	// タグがロックされているかどうか
	IsLocked bool `json:"isLocked,omitempty"`
	// タグ付与日時
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// タグ更新日時
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
