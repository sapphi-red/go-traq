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

type Pin struct {
	// ピンUUID
	PinId string `json:"pinId,omitempty"`
	// チャンネルUUID
	ChannelId string `json:"channelId,omitempty"`
	// ピン留めしたユーザーのUUID
	UserId string `json:"userId,omitempty"`
	// ピン留めした日時
	DateTime time.Time `json:"dateTime,omitempty"`
	Message Message `json:"message,omitempty"`
}
