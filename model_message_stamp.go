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

type MessageStamp struct {
	// ユーザーUUID
	UserId string `json:"userId,omitempty"`
	// スタンプUUID
	StampId string `json:"stampId,omitempty"`
	// 押された個数
	Count int32 `json:"count,omitempty"`
	// 最初に押した日時
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// 最後に押した日時
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}
