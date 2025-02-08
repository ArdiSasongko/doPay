/*
Fast Campus Pay (Wallet API Endpoint)

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0
Contact: azwar.nrst@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package repository

import (
	"encoding/json"
)

// checks if the V1UpdateBalanceResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1UpdateBalanceResponse{}

// V1UpdateBalanceResponse struct for V1UpdateBalanceResponse
type V1UpdateBalanceResponse struct {
	Message      *string  `json:"message,omitempty"`
	Success      *bool    `json:"success,omitempty"`
	FinalBalance *float64 `json:"finalBalance,omitempty"`
}

// NewV1UpdateBalanceResponse instantiates a new V1UpdateBalanceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1UpdateBalanceResponse() *V1UpdateBalanceResponse {
	this := V1UpdateBalanceResponse{}
	return &this
}

// NewV1UpdateBalanceResponseWithDefaults instantiates a new V1UpdateBalanceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1UpdateBalanceResponseWithDefaults() *V1UpdateBalanceResponse {
	this := V1UpdateBalanceResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *V1UpdateBalanceResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateBalanceResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *V1UpdateBalanceResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *V1UpdateBalanceResponse) SetMessage(v string) {
	o.Message = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *V1UpdateBalanceResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateBalanceResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *V1UpdateBalanceResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *V1UpdateBalanceResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetFinalBalance returns the FinalBalance field value if set, zero value otherwise.
func (o *V1UpdateBalanceResponse) GetFinalBalance() float64 {
	if o == nil || IsNil(o.FinalBalance) {
		var ret float64
		return ret
	}
	return *o.FinalBalance
}

// GetFinalBalanceOk returns a tuple with the FinalBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1UpdateBalanceResponse) GetFinalBalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.FinalBalance) {
		return nil, false
	}
	return o.FinalBalance, true
}

// HasFinalBalance returns a boolean if a field has been set.
func (o *V1UpdateBalanceResponse) HasFinalBalance() bool {
	if o != nil && !IsNil(o.FinalBalance) {
		return true
	}

	return false
}

// SetFinalBalance gets a reference to the given float64 and assigns it to the FinalBalance field.
func (o *V1UpdateBalanceResponse) SetFinalBalance(v float64) {
	o.FinalBalance = &v
}

func (o V1UpdateBalanceResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1UpdateBalanceResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.FinalBalance) {
		toSerialize["finalBalance"] = o.FinalBalance
	}
	return toSerialize, nil
}

type NullableV1UpdateBalanceResponse struct {
	value *V1UpdateBalanceResponse
	isSet bool
}

func (v NullableV1UpdateBalanceResponse) Get() *V1UpdateBalanceResponse {
	return v.value
}

func (v *NullableV1UpdateBalanceResponse) Set(val *V1UpdateBalanceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1UpdateBalanceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1UpdateBalanceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1UpdateBalanceResponse(val *V1UpdateBalanceResponse) *NullableV1UpdateBalanceResponse {
	return &NullableV1UpdateBalanceResponse{value: val, isSet: true}
}

func (v NullableV1UpdateBalanceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1UpdateBalanceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
