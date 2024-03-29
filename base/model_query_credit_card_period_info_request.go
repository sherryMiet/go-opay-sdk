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

// QueryCreditCardPeriodInfoRequest struct for QueryCreditCardPeriodInfoRequest
type QueryCreditCardPeriodInfoRequest struct {
	// **特店編號(由綠界提供)**
	MerchantID string `json:"MerchantID" form:"MerchantID"`
	// **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。
	MerchantTradeNo string `json:"MerchantTradeNo" form:"MerchantTradeNo"`
	// **驗證時間**   將當下的時間轉為UnixTimeStamp(見範例)用於驗證此次介接的時間區間。   綠界驗證時間區間暫訂為 3 分鐘內有效，超過則此次介接無效。
	TimeStamp int `json:"TimeStamp" form:"TimeStamp"`
	// **檢查碼** 請參考附錄檢查碼機制與產生檢查碼範例程式
	CheckMacValue string `json:"CheckMacValue" form:"CheckMacValue"`
}

// NewQueryCreditCardPeriodInfoRequest instantiates a new QueryCreditCardPeriodInfoRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueryCreditCardPeriodInfoRequest(merchantID string, merchantTradeNo string, timeStamp int, checkMacValue string) *QueryCreditCardPeriodInfoRequest {
	this := QueryCreditCardPeriodInfoRequest{}
	this.MerchantID = merchantID
	this.MerchantTradeNo = merchantTradeNo
	this.TimeStamp = timeStamp
	this.CheckMacValue = checkMacValue
	return &this
}

// NewQueryCreditCardPeriodInfoRequestWithDefaults instantiates a new QueryCreditCardPeriodInfoRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueryCreditCardPeriodInfoRequestWithDefaults() *QueryCreditCardPeriodInfoRequest {
	this := QueryCreditCardPeriodInfoRequest{}
	return &this
}

// GetMerchantID returns the MerchantID field value
func (o *QueryCreditCardPeriodInfoRequest) GetMerchantID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantID
}

// GetMerchantIDOk returns a tuple with the MerchantID field value
// and a boolean to check if the value has been set.
func (o *QueryCreditCardPeriodInfoRequest) GetMerchantIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantID, true
}

// SetMerchantID sets field value
func (o *QueryCreditCardPeriodInfoRequest) SetMerchantID(v string) {
	o.MerchantID = v
}

// GetMerchantTradeNo returns the MerchantTradeNo field value
func (o *QueryCreditCardPeriodInfoRequest) GetMerchantTradeNo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MerchantTradeNo
}

// GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field value
// and a boolean to check if the value has been set.
func (o *QueryCreditCardPeriodInfoRequest) GetMerchantTradeNoOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MerchantTradeNo, true
}

// SetMerchantTradeNo sets field value
func (o *QueryCreditCardPeriodInfoRequest) SetMerchantTradeNo(v string) {
	o.MerchantTradeNo = v
}

// GetTimeStamp returns the TimeStamp field value
func (o *QueryCreditCardPeriodInfoRequest) GetTimeStamp() int {
	if o == nil {
		var ret int
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *QueryCreditCardPeriodInfoRequest) GetTimeStampOk() (*int, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *QueryCreditCardPeriodInfoRequest) SetTimeStamp(v int) {
	o.TimeStamp = v
}

// GetCheckMacValue returns the CheckMacValue field value
func (o *QueryCreditCardPeriodInfoRequest) GetCheckMacValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CheckMacValue
}

// GetCheckMacValueOk returns a tuple with the CheckMacValue field value
// and a boolean to check if the value has been set.
func (o *QueryCreditCardPeriodInfoRequest) GetCheckMacValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CheckMacValue, true
}

// SetCheckMacValue sets field value
func (o *QueryCreditCardPeriodInfoRequest) SetCheckMacValue(v string) {
	o.CheckMacValue = v
}

func (o QueryCreditCardPeriodInfoRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["MerchantID"] = o.MerchantID
	}
	if true {
		toSerialize["MerchantTradeNo"] = o.MerchantTradeNo
	}
	if true {
		toSerialize["TimeStamp"] = o.TimeStamp
	}
	if true {
		toSerialize["CheckMacValue"] = o.CheckMacValue
	}
	return json.Marshal(toSerialize)
}

type NullableQueryCreditCardPeriodInfoRequest struct {
	value *QueryCreditCardPeriodInfoRequest
	isSet bool
}

func (v NullableQueryCreditCardPeriodInfoRequest) Get() *QueryCreditCardPeriodInfoRequest {
	return v.value
}

func (v *NullableQueryCreditCardPeriodInfoRequest) Set(val *QueryCreditCardPeriodInfoRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableQueryCreditCardPeriodInfoRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableQueryCreditCardPeriodInfoRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueryCreditCardPeriodInfoRequest(val *QueryCreditCardPeriodInfoRequest) *NullableQueryCreditCardPeriodInfoRequest {
	return &NullableQueryCreditCardPeriodInfoRequest{value: val, isSet: true}
}

func (v NullableQueryCreditCardPeriodInfoRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueryCreditCardPeriodInfoRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
