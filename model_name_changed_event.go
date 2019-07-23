/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// NameChanged
type NameChangedEvent struct {
	// 変更者UUID
	UserId string `json:"userId,omitempty"`
	// 変更前名前
	Before string `json:"before,omitempty"`
	// 変更後名前
	After string `json:"after,omitempty"`
}