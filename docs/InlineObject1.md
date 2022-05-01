# InlineObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Phone** | **string** |  | 
**Account** | **string** | 通过接入jugugu申请 | 
**Secret** | **string** | 通过接入jugugu申请 | 
**Time** | **int32** | int64类型 | 
**Randomtoken** | **string** | 可以由sha256函数随机产生，避免2分钟内重复，可用时间+其他唯一参数作为种子，防止重放攻击 | 
**Paymentpassword** | Pointer to **string** | 可以自由设置也可以传“”空字符串，系统随机生成，建议让系统随机生成 | [optional] 
**Openid** | **string** | 项目方账户系统的唯一识别码，用于绑定登录jugugu系统 | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


