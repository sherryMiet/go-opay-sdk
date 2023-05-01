/*
 * ECPay API
 *
 * 綠界金流 API 定義文件
 *
 * API version: 0.0.24
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ecpayBase

import (
	"encoding/json"
)

// EncryptTypeEnum **CheckMacValue加密類型** 請固定填入 `1`，使用 SHA256 加密。
type EncryptTypeEnum int

// List of EncryptTypeEnum
const (
	ENCRYPTTYPEENUM_SHA256 EncryptTypeEnum = 1
)

// Ptr returns reference to EncryptTypeEnum value
func (v EncryptTypeEnum) Ptr() *EncryptTypeEnum {
	return &v
}

type NullableEncryptTypeEnum struct {
	value *EncryptTypeEnum
	isSet bool
}

func (v NullableEncryptTypeEnum) Get() *EncryptTypeEnum {
	return v.value
}

func (v *NullableEncryptTypeEnum) Set(val *EncryptTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableEncryptTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableEncryptTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEncryptTypeEnum(val *EncryptTypeEnum) *NullableEncryptTypeEnum {
	return &NullableEncryptTypeEnum{value: val, isSet: true}
}

func (v NullableEncryptTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEncryptTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
