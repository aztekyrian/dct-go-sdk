/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// DisableEnvironmentResponse struct for DisableEnvironmentResponse
type DisableEnvironmentResponse struct {
	// The initiated job id, in most cases it will be null
	JobId *string `json:"job_id,omitempty"`
}

// NewDisableEnvironmentResponse instantiates a new DisableEnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDisableEnvironmentResponse() *DisableEnvironmentResponse {
	this := DisableEnvironmentResponse{}
	return &this
}

// NewDisableEnvironmentResponseWithDefaults instantiates a new DisableEnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDisableEnvironmentResponseWithDefaults() *DisableEnvironmentResponse {
	this := DisableEnvironmentResponse{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *DisableEnvironmentResponse) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DisableEnvironmentResponse) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *DisableEnvironmentResponse) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *DisableEnvironmentResponse) SetJobId(v string) {
	o.JobId = &v
}

func (o DisableEnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	return json.Marshal(toSerialize)
}

type NullableDisableEnvironmentResponse struct {
	value *DisableEnvironmentResponse
	isSet bool
}

func (v NullableDisableEnvironmentResponse) Get() *DisableEnvironmentResponse {
	return v.value
}

func (v *NullableDisableEnvironmentResponse) Set(val *DisableEnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDisableEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDisableEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDisableEnvironmentResponse(val *DisableEnvironmentResponse) *NullableDisableEnvironmentResponse {
	return &NullableDisableEnvironmentResponse{value: val, isSet: true}
}

func (v NullableDisableEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDisableEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

