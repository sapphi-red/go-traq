/*
 * traQ API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 2.6.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse20010 struct {
	TagId string `json:"tagId,omitempty"`
	Tag string `json:"tag,omitempty"`
	// UUIDの配列
	Users []string `json:"users,omitempty"`
}
