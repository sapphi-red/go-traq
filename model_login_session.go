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

// LoginSession ログインセッション情報
type LoginSession struct {
	// セッションUUID
	Id string `json:"id"`
	// 発行日時
	IssuedAt time.Time `json:"issuedAt"`
}
