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

type StampHistory struct {
	// スタンプID
	StampId string `json:"stampId,omitempty"`
	// そのスタンプが最後に押された日時
	Datetime time.Time `json:"datetime,omitempty"`
}