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

type Session struct {
	Id            string    `json:"id,omitempty"`
	LastIP        string    `json:"lastIP,omitempty"`
	LastUserAgent string    `json:"lastUserAgent,omitempty"`
	LastAccess    time.Time `json:"lastAccess,omitempty"`
	CreatedAt     time.Time `json:"createdAt,omitempty"`
}