/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// VisibilityChanged
type VisibilityChangedEvent struct {
	// 作成者UUID
	UserId string `json:"userId,omitempty"`
	// 可視状態
	Visibility bool `json:"visibility,omitempty"`
}
