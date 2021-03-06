/*
 * Jugugu密钥Fast模式
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineObject1 struct for InlineObject1
type InlineObject1 struct {
	Phone string `json:"phone"`
	// 通过接入jugugu申请
	Account string `json:"account"`
	// 通过接入jugugu申请
	Secret string `json:"secret"`
	// int64类型
	Time int32 `json:"time"`
	// 可以由sha256函数随机产生，避免2分钟内重复，可用时间+其他唯一参数作为种子，防止重放攻击
	Randomtoken string `json:"randomtoken"`
	// 可以自由设置也可以传“”空字符串，系统随机生成，建议让系统随机生成
	Paymentpassword *string `json:"paymentpassword,omitempty"`
	// 项目方账户系统的唯一识别码，用于绑定登录jugugu系统
	Openid string `json:"openid"`
}
