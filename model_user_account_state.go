/*
 * traQ v3
 *
 * traQ v3 API
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package traq

// UserAccountState ユーザーアカウント状態 0: 停止 1: 有効 2: 一時停止
type UserAccountState int32

// List of UserAccountState
const (
	USERACCOUNTSTATE_deactivated UserAccountState = 0
	USERACCOUNTSTATE_active      UserAccountState = 1
	USERACCOUNTSTATE_suspended   UserAccountState = 2
)
