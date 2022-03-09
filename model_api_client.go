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

// ApiClient struct for ApiClient
type ApiClient struct {
	// The entity ID of this API client.
	Id *int64 `json:"id,omitempty"`
	// The unique ID which is used to identity the identity of an API request. The web server (nginx) configuration must be configured so as to include the external ID as the value of the X_CLIENT_ID HTTP request header when requests are proxied. For OAuth2/JWT based authentication, this typically corresponds to a value extracted from the JWT, uniquely identifying the API client.
	ApiClientId *string `json:"api_client_id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewApiClient instantiates a new ApiClient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClient() *ApiClient {
	this := ApiClient{}
	return &this
}

// NewApiClientWithDefaults instantiates a new ApiClient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClientWithDefaults() *ApiClient {
	this := ApiClient{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiClient) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClient) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiClient) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ApiClient) SetId(v int64) {
	o.Id = &v
}

// GetApiClientId returns the ApiClientId field value if set, zero value otherwise.
func (o *ApiClient) GetApiClientId() string {
	if o == nil || o.ApiClientId == nil {
		var ret string
		return ret
	}
	return *o.ApiClientId
}

// GetApiClientIdOk returns a tuple with the ApiClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClient) GetApiClientIdOk() (*string, bool) {
	if o == nil || o.ApiClientId == nil {
		return nil, false
	}
	return o.ApiClientId, true
}

// HasApiClientId returns a boolean if a field has been set.
func (o *ApiClient) HasApiClientId() bool {
	if o != nil && o.ApiClientId != nil {
		return true
	}

	return false
}

// SetApiClientId gets a reference to the given string and assigns it to the ApiClientId field.
func (o *ApiClient) SetApiClientId(v string) {
	o.ApiClientId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiClient) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClient) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiClient) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiClient) SetName(v string) {
	o.Name = &v
}

func (o ApiClient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ApiClientId != nil {
		toSerialize["api_client_id"] = o.ApiClientId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableApiClient struct {
	value *ApiClient
	isSet bool
}

func (v NullableApiClient) Get() *ApiClient {
	return v.value
}

func (v *NullableApiClient) Set(val *ApiClient) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClient) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClient(val *ApiClient) *NullableApiClient {
	return &NullableApiClient{value: val, isSet: true}
}

func (v NullableApiClient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


