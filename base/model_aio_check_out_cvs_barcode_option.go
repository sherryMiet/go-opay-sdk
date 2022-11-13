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

// AioCheckOutCvsBarcodeOption struct for AioCheckOutCvsBarcodeOption
type AioCheckOutCvsBarcodeOption struct {
	// **超商繳費截止時間**   注意事項：   `CVS`:以分鐘為單位   `BARCODE`:以天為單位   若未設定此參數，`CVS` 預設為 `10080` 分鐘(`7` 天)；BARCODE 預設為 `7` 天。   若需設定此參數，請於建立訂單時將此參數送給綠界。提醒您，CVS 帶入數值不可超過 `86400` 分鐘，超過時一律以 `86400` 分鐘計(`60` 天)   例：`08/01` 的 `20:15` 分購買商品，繳費期限為 `7` 天，表示 `8/08` 的 `20:15` 分前您必須前往超商繳費。   範例: `CVS`=`1440`(共 `1` 天)、`BARCODE`=`7`(共 `7` 天)
	StoreExpireDate *int `json:"StoreExpireDate,omitempty" form:"StoreExpireDate"`
	// **交易描述1** 會出現在超商繳費平台螢幕上
	Desc1 *string `json:"Desc_1,omitempty" form:"Desc_1"`
	// **交易描述2** 會出現在超商繳費平台螢幕上
	Desc2 *string `json:"Desc_2,omitempty" form:"Desc_2"`
	// **交易描述3** 會出現在超商繳費平台螢幕上
	Desc3 *string `json:"Desc_3,omitempty" form:"Desc_3"`
	// **交易描述4** 會出現在超商繳費平台螢幕上
	Desc4 *string `json:"Desc_4,omitempty" form:"Desc_4"`
	// **Server 端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)。   請參考[`ATM`、`CVS` 或 `BARCODE` 的取號結果通知.]   注意事項：   頁面將會停留在綠界，顯示繳費的相關資訊。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。
	PaymentInfoURL *string `json:"PaymentInfoURL,omitempty" form:"PaymentInfoURL"`
	// **Client端回傳付款方式相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會從 Client 端回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)且將頁面轉到特店指定的頁面。   請參考[`ATM`、`CVS` 或 `BARCODE` 的取號結果通知.]   注意事項：   若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。
	ClientRedirectURL *string `json:"ClientRedirectURL,omitempty" form:"ClientRedirectURL"`
}

// NewAioCheckOutCvsBarcodeOption instantiates a new AioCheckOutCvsBarcodeOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAioCheckOutCvsBarcodeOption() *AioCheckOutCvsBarcodeOption {
	this := AioCheckOutCvsBarcodeOption{}
	return &this
}

// NewAioCheckOutCvsBarcodeOptionWithDefaults instantiates a new AioCheckOutCvsBarcodeOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAioCheckOutCvsBarcodeOptionWithDefaults() *AioCheckOutCvsBarcodeOption {
	this := AioCheckOutCvsBarcodeOption{}
	return &this
}

// GetStoreExpireDate returns the StoreExpireDate field value if set, zero value otherwise.
func (o *AioCheckOutCvsBarcodeOption) GetStoreExpireDate() int {
	if o == nil || o.StoreExpireDate == nil {
		var ret int
		return ret
	}
	return *o.StoreExpireDate
}

// GetStoreExpireDateOk returns a tuple with the StoreExpireDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCvsBarcodeOption) GetStoreExpireDateOk() (*int, bool) {
	if o == nil || o.StoreExpireDate == nil {
		return nil, false
	}
	return o.StoreExpireDate, true
}

// HasStoreExpireDate returns a boolean if a field has been set.
func (o *AioCheckOutCvsBarcodeOption) HasStoreExpireDate() bool {
	if o != nil && o.StoreExpireDate != nil {
		return true
	}

	return false
}

// SetStoreExpireDate gets a reference to the given int and assigns it to the StoreExpireDate field.
func (o *AioCheckOutCvsBarcodeOption) SetStoreExpireDate(v int) {
	o.StoreExpireDate = &v
}

// GetDesc1 returns the Desc1 field value if set, zero value otherwise.
func (o *AioCheckOutCvsBarcodeOption) GetDesc1() string {
	if o == nil || o.Desc1 == nil {
		var ret string
		return ret
	}
	return *o.Desc1
}

// GetDesc1Ok returns a tuple with the Desc1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCvsBarcodeOption) GetDesc1Ok() (*string, bool) {
	if o == nil || o.Desc1 == nil {
		return nil, false
	}
	return o.Desc1, true
}

// HasDesc1 returns a boolean if a field has been set.
func (o *AioCheckOutCvsBarcodeOption) HasDesc1() bool {
	if o != nil && o.Desc1 != nil {
		return true
	}

	return false
}

// SetDesc1 gets a reference to the given string and assigns it to the Desc1 field.
func (o *AioCheckOutCvsBarcodeOption) SetDesc1(v string) {
	o.Desc1 = &v
}

// GetDesc2 returns the Desc2 field value if set, zero value otherwise.
func (o *AioCheckOutCvsBarcodeOption) GetDesc2() string {
	if o == nil || o.Desc2 == nil {
		var ret string
		return ret
	}
	return *o.Desc2
}

// GetDesc2Ok returns a tuple with the Desc2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCvsBarcodeOption) GetDesc2Ok() (*string, bool) {
	if o == nil || o.Desc2 == nil {
		return nil, false
	}
	return o.Desc2, true
}

// HasDesc2 returns a boolean if a field has been set.
func (o *AioCheckOutCvsBarcodeOption) HasDesc2() bool {
	if o != nil && o.Desc2 != nil {
		return true
	}

	return false
}

// SetDesc2 gets a reference to the given string and assigns it to the Desc2 field.
func (o *AioCheckOutCvsBarcodeOption) SetDesc2(v string) {
	o.Desc2 = &v
}

// GetDesc3 returns the Desc3 field value if set, zero value otherwise.
func (o *AioCheckOutCvsBarcodeOption) GetDesc3() string {
	if o == nil || o.Desc3 == nil {
		var ret string
		return ret
	}
	return *o.Desc3
}

// GetDesc3Ok returns a tuple with the Desc3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCvsBarcodeOption) GetDesc3Ok() (*string, bool) {
	if o == nil || o.Desc3 == nil {
		return nil, false
	}
	return o.Desc3, true
}

// HasDesc3 returns a boolean if a field has been set.
func (o *AioCheckOutCvsBarcodeOption) HasDesc3() bool {
	if o != nil && o.Desc3 != nil {
		return true
	}

	return false
}

// SetDesc3 gets a reference to the given string and assigns it to the Desc3 field.
func (o *AioCheckOutCvsBarcodeOption) SetDesc3(v string) {
	o.Desc3 = &v
}

// GetDesc4 returns the Desc4 field value if set, zero value otherwise.
func (o *AioCheckOutCvsBarcodeOption) GetDesc4() string {
	if o == nil || o.Desc4 == nil {
		var ret string
		return ret
	}
	return *o.Desc4
}

// GetDesc4Ok returns a tuple with the Desc4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCvsBarcodeOption) GetDesc4Ok() (*string, bool) {
	if o == nil || o.Desc4 == nil {
		return nil, false
	}
	return o.Desc4, true
}

// HasDesc4 returns a boolean if a field has been set.
func (o *AioCheckOutCvsBarcodeOption) HasDesc4() bool {
	if o != nil && o.Desc4 != nil {
		return true
	}

	return false
}

// SetDesc4 gets a reference to the given string and assigns it to the Desc4 field.
func (o *AioCheckOutCvsBarcodeOption) SetDesc4(v string) {
	o.Desc4 = &v
}

// GetPaymentInfoURL returns the PaymentInfoURL field value if set, zero value otherwise.
func (o *AioCheckOutCvsBarcodeOption) GetPaymentInfoURL() string {
	if o == nil || o.PaymentInfoURL == nil {
		var ret string
		return ret
	}
	return *o.PaymentInfoURL
}

// GetPaymentInfoURLOk returns a tuple with the PaymentInfoURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCvsBarcodeOption) GetPaymentInfoURLOk() (*string, bool) {
	if o == nil || o.PaymentInfoURL == nil {
		return nil, false
	}
	return o.PaymentInfoURL, true
}

// HasPaymentInfoURL returns a boolean if a field has been set.
func (o *AioCheckOutCvsBarcodeOption) HasPaymentInfoURL() bool {
	if o != nil && o.PaymentInfoURL != nil {
		return true
	}

	return false
}

// SetPaymentInfoURL gets a reference to the given string and assigns it to the PaymentInfoURL field.
func (o *AioCheckOutCvsBarcodeOption) SetPaymentInfoURL(v string) {
	o.PaymentInfoURL = &v
}

// GetClientRedirectURL returns the ClientRedirectURL field value if set, zero value otherwise.
func (o *AioCheckOutCvsBarcodeOption) GetClientRedirectURL() string {
	if o == nil || o.ClientRedirectURL == nil {
		var ret string
		return ret
	}
	return *o.ClientRedirectURL
}

// GetClientRedirectURLOk returns a tuple with the ClientRedirectURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AioCheckOutCvsBarcodeOption) GetClientRedirectURLOk() (*string, bool) {
	if o == nil || o.ClientRedirectURL == nil {
		return nil, false
	}
	return o.ClientRedirectURL, true
}

// HasClientRedirectURL returns a boolean if a field has been set.
func (o *AioCheckOutCvsBarcodeOption) HasClientRedirectURL() bool {
	if o != nil && o.ClientRedirectURL != nil {
		return true
	}

	return false
}

// SetClientRedirectURL gets a reference to the given string and assigns it to the ClientRedirectURL field.
func (o *AioCheckOutCvsBarcodeOption) SetClientRedirectURL(v string) {
	o.ClientRedirectURL = &v
}

func (o AioCheckOutCvsBarcodeOption) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StoreExpireDate != nil {
		toSerialize["StoreExpireDate"] = o.StoreExpireDate
	}
	if o.Desc1 != nil {
		toSerialize["Desc_1"] = o.Desc1
	}
	if o.Desc2 != nil {
		toSerialize["Desc_2"] = o.Desc2
	}
	if o.Desc3 != nil {
		toSerialize["Desc_3"] = o.Desc3
	}
	if o.Desc4 != nil {
		toSerialize["Desc_4"] = o.Desc4
	}
	if o.PaymentInfoURL != nil {
		toSerialize["PaymentInfoURL"] = o.PaymentInfoURL
	}
	if o.ClientRedirectURL != nil {
		toSerialize["ClientRedirectURL"] = o.ClientRedirectURL
	}
	return json.Marshal(toSerialize)
}

type NullableAioCheckOutCvsBarcodeOption struct {
	value *AioCheckOutCvsBarcodeOption
	isSet bool
}

func (v NullableAioCheckOutCvsBarcodeOption) Get() *AioCheckOutCvsBarcodeOption {
	return v.value
}

func (v *NullableAioCheckOutCvsBarcodeOption) Set(val *AioCheckOutCvsBarcodeOption) {
	v.value = val
	v.isSet = true
}

func (v NullableAioCheckOutCvsBarcodeOption) IsSet() bool {
	return v.isSet
}

func (v *NullableAioCheckOutCvsBarcodeOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAioCheckOutCvsBarcodeOption(val *AioCheckOutCvsBarcodeOption) *NullableAioCheckOutCvsBarcodeOption {
	return &NullableAioCheckOutCvsBarcodeOption{value: val, isSet: true}
}

func (v NullableAioCheckOutCvsBarcodeOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAioCheckOutCvsBarcodeOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}