# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [standard.proto](#standard.proto)
    - [CreateUnitDataRequest](#standard.CreateUnitDataRequest)
    - [CreateUnitDataResponse](#standard.CreateUnitDataResponse)
    - [DestroyUnitDataByKeyRequest](#standard.DestroyUnitDataByKeyRequest)
    - [DestroyUnitDataByKeyResponse](#standard.DestroyUnitDataByKeyResponse)
    - [GetUnitDataByKeyRequest](#standard.GetUnitDataByKeyRequest)
    - [GetUnitDataByKeyResponse](#standard.GetUnitDataByKeyResponse)
    - [UnitData](#standard.UnitData)
    - [UpdateUnitDataBodyByKeyRequest](#standard.UpdateUnitDataBodyByKeyRequest)
    - [UpdateUnitDataBodyByKeyResponse](#standard.UpdateUnitDataBodyByKeyResponse)
    - [UpdateUnitDataDestroyTimeByKeyRequest](#standard.UpdateUnitDataDestroyTimeByKeyRequest)
    - [UpdateUnitDataDestroyTimeByKeyResponse](#standard.UpdateUnitDataDestroyTimeByKeyResponse)
    - [UpdateUnitDataEffectiveTimeByKeyRequest](#standard.UpdateUnitDataEffectiveTimeByKeyRequest)
    - [UpdateUnitDataEffectiveTimeByKeyResponse](#standard.UpdateUnitDataEffectiveTimeByKeyResponse)
    - [UpdateUnitDataExpiryTimeByKeyRequest](#standard.UpdateUnitDataExpiryTimeByKeyRequest)
    - [UpdateUnitDataExpiryTimeByKeyResponse](#standard.UpdateUnitDataExpiryTimeByKeyResponse)
  
    - [State](#standard.State)
  
  
    - [Swaper](#standard.Swaper)
  

- [Scalar Value Types](#scalar-value-types)



<a name="standard.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## standard.proto



<a name="standard.CreateUnitDataRequest"></a>

### CreateUnitDataRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Body | [string](#string) |  | 数据 |
| ExpiryTime | [int64](#int64) |  | 失效时间 服务器当前时间 &#43; ExpiryTime ms |
| DestroyTime | [int64](#int64) |  | 销毁时间 服务器当前时间 &#43; DestroyTime ms |
| EffectiveTime | [int64](#int64) |  | 生效时间 服务器当前时间 &#43; EffectiveTime ms |






<a name="standard.CreateUnitDataResponse"></a>

### CreateUnitDataResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  |  |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.DestroyUnitDataByKeyRequest"></a>

### DestroyUnitDataByKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  |  |






<a name="standard.DestroyUnitDataByKeyResponse"></a>

### DestroyUnitDataByKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.GetUnitDataByKeyRequest"></a>

### GetUnitDataByKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  |  |






<a name="standard.GetUnitDataByKeyResponse"></a>

### GetUnitDataByKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |
| Data | [UnitData](#standard.UnitData) |  |  |






<a name="standard.UnitData"></a>

### UnitData
单位数据


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Body | [string](#string) |  | 数据 |
| ExpiryTime | [int64](#int64) |  | 失效时间 |
| CreatedTime | [int64](#int64) |  | 创建时间 剩余多少 ms |
| DestroyTime | [int64](#int64) |  | 销毁时间 剩余多少 ms |
| EffectiveTime | [int64](#int64) |  | 生效时间 剩余多少 ms |






<a name="standard.UpdateUnitDataBodyByKeyRequest"></a>

### UpdateUnitDataBodyByKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  |  |
| Body | [string](#string) |  |  |






<a name="standard.UpdateUnitDataBodyByKeyResponse"></a>

### UpdateUnitDataBodyByKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUnitDataDestroyTimeByKeyRequest"></a>

### UpdateUnitDataDestroyTimeByKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  |  |
| DestroyTime | [int64](#int64) |  |  |






<a name="standard.UpdateUnitDataDestroyTimeByKeyResponse"></a>

### UpdateUnitDataDestroyTimeByKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUnitDataEffectiveTimeByKeyRequest"></a>

### UpdateUnitDataEffectiveTimeByKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  |  |
| EffectiveTime | [int64](#int64) |  |  |






<a name="standard.UpdateUnitDataEffectiveTimeByKeyResponse"></a>

### UpdateUnitDataEffectiveTimeByKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |






<a name="standard.UpdateUnitDataExpiryTimeByKeyRequest"></a>

### UpdateUnitDataExpiryTimeByKeyRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| Key | [string](#string) |  |  |
| ExpiryTime | [int64](#int64) |  |  |






<a name="standard.UpdateUnitDataExpiryTimeByKeyResponse"></a>

### UpdateUnitDataExpiryTimeByKeyResponse



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| State | [State](#standard.State) |  |  |
| Message | [string](#string) |  |  |





 


<a name="standard.State"></a>

### State
状态

| Name | Number | Description |
| ---- | ------ | ----------- |
| UNKNOWN | 0 | 未知 |
| SUCCESS | 1 | 成功 |
| FAILURE | 2 | 失败 |
| SERVICE_ERROR | 3 | 服务错误 |
| PARAMS_INVALID | 4 | 参数不合法 |
| ILLEGAL_REQUEST | 5 | 非法请求 |
| UNIT_DATA_EXPIRED | 6 | 数据过期了 |
| UNIT_DATA_NOT_EXIST | 7 | 数据不存在 |
| UNIT_DATA_NOT_ACTIVE | 8 | 数据未到生效期 |
| DB_OPERATION_FATLURE | 9 | 数据库操作失败 |


 

 


<a name="standard.Swaper"></a>

### Swaper
在里面流通的是 `UnitData`
每个 `UnitData` 有自己的一些生命周期设定
可以创建一个 `UnitData` 并获得这个 `UnitData` 的对应加密 number
或者通过一个的特定的 `UnitData` 的 number 在有效的生命周期内获取这个数据

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateUnitData | [CreateUnitDataRequest](#standard.CreateUnitDataRequest) | [CreateUnitDataResponse](#standard.CreateUnitDataResponse) |  |
| GetUnitDataByKey | [GetUnitDataByKeyRequest](#standard.GetUnitDataByKeyRequest) | [GetUnitDataByKeyResponse](#standard.GetUnitDataByKeyResponse) |  |
| DestroyUnitDataByKey | [DestroyUnitDataByKeyRequest](#standard.DestroyUnitDataByKeyRequest) | [DestroyUnitDataByKeyResponse](#standard.DestroyUnitDataByKeyResponse) |  |
| UpdateUnitDataBodyByKey | [UpdateUnitDataBodyByKeyRequest](#standard.UpdateUnitDataBodyByKeyRequest) | [UpdateUnitDataBodyByKeyResponse](#standard.UpdateUnitDataBodyByKeyResponse) |  |
| UpdateUnitDataExpiryTimeByKey | [UpdateUnitDataExpiryTimeByKeyRequest](#standard.UpdateUnitDataExpiryTimeByKeyRequest) | [UpdateUnitDataExpiryTimeByKeyResponse](#standard.UpdateUnitDataExpiryTimeByKeyResponse) |  |
| UpdateUnitDataDestroyTimeByKey | [UpdateUnitDataDestroyTimeByKeyRequest](#standard.UpdateUnitDataDestroyTimeByKeyRequest) | [UpdateUnitDataDestroyTimeByKeyResponse](#standard.UpdateUnitDataDestroyTimeByKeyResponse) |  |
| UpdateUnitDataEffectiveTimeByKey | [UpdateUnitDataEffectiveTimeByKeyRequest](#standard.UpdateUnitDataEffectiveTimeByKeyRequest) | [UpdateUnitDataEffectiveTimeByKeyResponse](#standard.UpdateUnitDataEffectiveTimeByKeyResponse) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

