/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// WebRtcChannelState struct for WebRtcChannelState
type WebRtcChannelState struct {
	// 接続ユーザーの状態の配列
	Users []WebRtcUserState `json:"users,omitempty"`
}