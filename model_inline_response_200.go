/*
 * Jugugu密钥Fast模式
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse200 struct for InlineResponse200
type InlineResponse200 struct {
	// -1代表失败并错误，0代表失败并提示，1代表成功
	State string `json:"state"`
	// 详细信息
	Msg string `json:"msg"`
}
