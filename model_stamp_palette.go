/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"time"
)

// StampPalette スタンプパレット情報
type StampPalette struct {
	// スタンプパレットUUID
	Id string `json:"id"`
	// パレット名
	Name string `json:"name"`
	// パレット内のスタンプのUUID配列
	Stamps []string `json:"stamps"`
	// 作成者UUID
	CreatorId string `json:"creatorId"`
	// パレット作成日時
	CreatedAt time.Time `json:"createdAt"`
	// パレット更新日時
	UpdatedAt time.Time `json:"updatedAt"`
	// パレット説明
	Description string `json:"description"`
}