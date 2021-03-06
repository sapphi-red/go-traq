/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// OgpMedia OGPに含まれる画像の情報
type OgpMedia struct {
	Url       string  `json:"url"`
	SecureUrl *string `json:"secureUrl"`
	Type      *string `json:"type"`
	Width     *int32  `json:"width"`
	Height    *int32  `json:"height"`
}
