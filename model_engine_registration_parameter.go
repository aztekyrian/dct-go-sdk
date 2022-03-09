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

// EngineRegistrationParameter Parameters to register and authenticate an engine.
type EngineRegistrationParameter struct {
	Name string `json:"name"`
	Hostname string `json:"hostname"`
	Username NullableString `json:"username,omitempty"`
	Password NullableString `json:"password,omitempty"`
	// Arguments to pass to the Vault CLI tool to retrieve the username for the engine.
	HashicorpVaultUsernameCommandArgs []string `json:"hashicorp_vault_username_command_args,omitempty"`
	// Arguments to pass to the Vault CLI tool to retrieve the password for the engine.
	HashicorpVaultPasswordCommandArgs []string `json:"hashicorp_vault_password_command_args,omitempty"`
	// Reference to the Hashicorp vault to use to retrieve engine credentials.
	HashicorpVaultId NullableInt64 `json:"hashicorp_vault_id,omitempty"`
	// Allow connections to the engine over HTTPs without validating the TLS certificate. Even though the connection to the engine might be performed over HTTPs, setting this property eliminates the protection against a man-in-the-middle attach for connections to this engine. Instead, consider creating a truststore with a Certificate Authority to validate the engine's certificate, and set the truststore_path propery. 
	InsecureSsl *bool `json:"insecure_ssl,omitempty"`
	// Ignore validation of the name associated to the TLS certificate when connecting to the engine over HTTPs. Setting this value must only be done if the TLS certificate of the engine does not match the hostname, and the TLS configuration of the engine cannot be fixed. Setting this property reduces the protection against a man-in-the-middle attack for connections to this engine. This is ignored if insecure_ssl is set. 
	UnsafeSslHostnameCheck *bool `json:"unsafe_ssl_hostname_check,omitempty"`
	// File name of a truststore which can be used to validate the TLS certificate of the engine. The truststore must be available at /etc/config/certs/<truststore_filename> 
	TruststoreFilename NullableString `json:"truststore_filename,omitempty"`
	// Password to read the truststore. 
	TruststorePassword NullableString `json:"truststore_password,omitempty"`
}

// NewEngineRegistrationParameter instantiates a new EngineRegistrationParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineRegistrationParameter(name string, hostname string) *EngineRegistrationParameter {
	this := EngineRegistrationParameter{}
	this.Name = name
	this.Hostname = hostname
	var insecureSsl bool = false
	this.InsecureSsl = &insecureSsl
	var unsafeSslHostnameCheck bool = false
	this.UnsafeSslHostnameCheck = &unsafeSslHostnameCheck
	return &this
}

// NewEngineRegistrationParameterWithDefaults instantiates a new EngineRegistrationParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineRegistrationParameterWithDefaults() *EngineRegistrationParameter {
	this := EngineRegistrationParameter{}
	var insecureSsl bool = false
	this.InsecureSsl = &insecureSsl
	var unsafeSslHostnameCheck bool = false
	this.UnsafeSslHostnameCheck = &unsafeSslHostnameCheck
	return &this
}

// GetName returns the Name field value
func (o *EngineRegistrationParameter) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EngineRegistrationParameter) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EngineRegistrationParameter) SetName(v string) {
	o.Name = v
}

// GetHostname returns the Hostname field value
func (o *EngineRegistrationParameter) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *EngineRegistrationParameter) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *EngineRegistrationParameter) SetHostname(v string) {
	o.Hostname = v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineRegistrationParameter) GetUsername() string {
	if o == nil || o.Username.Get() == nil {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngineRegistrationParameter) GetUsernameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *EngineRegistrationParameter) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *EngineRegistrationParameter) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *EngineRegistrationParameter) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineRegistrationParameter) GetPassword() string {
	if o == nil || o.Password.Get() == nil {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngineRegistrationParameter) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *EngineRegistrationParameter) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *EngineRegistrationParameter) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *EngineRegistrationParameter) UnsetPassword() {
	o.Password.Unset()
}

// GetHashicorpVaultUsernameCommandArgs returns the HashicorpVaultUsernameCommandArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineRegistrationParameter) GetHashicorpVaultUsernameCommandArgs() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.HashicorpVaultUsernameCommandArgs
}

// GetHashicorpVaultUsernameCommandArgsOk returns a tuple with the HashicorpVaultUsernameCommandArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngineRegistrationParameter) GetHashicorpVaultUsernameCommandArgsOk() ([]string, bool) {
	if o == nil || o.HashicorpVaultUsernameCommandArgs == nil {
		return nil, false
	}
	return o.HashicorpVaultUsernameCommandArgs, true
}

// HasHashicorpVaultUsernameCommandArgs returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasHashicorpVaultUsernameCommandArgs() bool {
	if o != nil && o.HashicorpVaultUsernameCommandArgs != nil {
		return true
	}

	return false
}

// SetHashicorpVaultUsernameCommandArgs gets a reference to the given []string and assigns it to the HashicorpVaultUsernameCommandArgs field.
func (o *EngineRegistrationParameter) SetHashicorpVaultUsernameCommandArgs(v []string) {
	o.HashicorpVaultUsernameCommandArgs = v
}

// GetHashicorpVaultPasswordCommandArgs returns the HashicorpVaultPasswordCommandArgs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineRegistrationParameter) GetHashicorpVaultPasswordCommandArgs() []string {
	if o == nil  {
		var ret []string
		return ret
	}
	return o.HashicorpVaultPasswordCommandArgs
}

// GetHashicorpVaultPasswordCommandArgsOk returns a tuple with the HashicorpVaultPasswordCommandArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngineRegistrationParameter) GetHashicorpVaultPasswordCommandArgsOk() ([]string, bool) {
	if o == nil || o.HashicorpVaultPasswordCommandArgs == nil {
		return nil, false
	}
	return o.HashicorpVaultPasswordCommandArgs, true
}

// HasHashicorpVaultPasswordCommandArgs returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasHashicorpVaultPasswordCommandArgs() bool {
	if o != nil && o.HashicorpVaultPasswordCommandArgs != nil {
		return true
	}

	return false
}

// SetHashicorpVaultPasswordCommandArgs gets a reference to the given []string and assigns it to the HashicorpVaultPasswordCommandArgs field.
func (o *EngineRegistrationParameter) SetHashicorpVaultPasswordCommandArgs(v []string) {
	o.HashicorpVaultPasswordCommandArgs = v
}

// GetHashicorpVaultId returns the HashicorpVaultId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineRegistrationParameter) GetHashicorpVaultId() int64 {
	if o == nil || o.HashicorpVaultId.Get() == nil {
		var ret int64
		return ret
	}
	return *o.HashicorpVaultId.Get()
}

// GetHashicorpVaultIdOk returns a tuple with the HashicorpVaultId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngineRegistrationParameter) GetHashicorpVaultIdOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.HashicorpVaultId.Get(), o.HashicorpVaultId.IsSet()
}

// HasHashicorpVaultId returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasHashicorpVaultId() bool {
	if o != nil && o.HashicorpVaultId.IsSet() {
		return true
	}

	return false
}

// SetHashicorpVaultId gets a reference to the given NullableInt64 and assigns it to the HashicorpVaultId field.
func (o *EngineRegistrationParameter) SetHashicorpVaultId(v int64) {
	o.HashicorpVaultId.Set(&v)
}
// SetHashicorpVaultIdNil sets the value for HashicorpVaultId to be an explicit nil
func (o *EngineRegistrationParameter) SetHashicorpVaultIdNil() {
	o.HashicorpVaultId.Set(nil)
}

// UnsetHashicorpVaultId ensures that no value is present for HashicorpVaultId, not even an explicit nil
func (o *EngineRegistrationParameter) UnsetHashicorpVaultId() {
	o.HashicorpVaultId.Unset()
}

// GetInsecureSsl returns the InsecureSsl field value if set, zero value otherwise.
func (o *EngineRegistrationParameter) GetInsecureSsl() bool {
	if o == nil || o.InsecureSsl == nil {
		var ret bool
		return ret
	}
	return *o.InsecureSsl
}

// GetInsecureSslOk returns a tuple with the InsecureSsl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineRegistrationParameter) GetInsecureSslOk() (*bool, bool) {
	if o == nil || o.InsecureSsl == nil {
		return nil, false
	}
	return o.InsecureSsl, true
}

// HasInsecureSsl returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasInsecureSsl() bool {
	if o != nil && o.InsecureSsl != nil {
		return true
	}

	return false
}

// SetInsecureSsl gets a reference to the given bool and assigns it to the InsecureSsl field.
func (o *EngineRegistrationParameter) SetInsecureSsl(v bool) {
	o.InsecureSsl = &v
}

// GetUnsafeSslHostnameCheck returns the UnsafeSslHostnameCheck field value if set, zero value otherwise.
func (o *EngineRegistrationParameter) GetUnsafeSslHostnameCheck() bool {
	if o == nil || o.UnsafeSslHostnameCheck == nil {
		var ret bool
		return ret
	}
	return *o.UnsafeSslHostnameCheck
}

// GetUnsafeSslHostnameCheckOk returns a tuple with the UnsafeSslHostnameCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EngineRegistrationParameter) GetUnsafeSslHostnameCheckOk() (*bool, bool) {
	if o == nil || o.UnsafeSslHostnameCheck == nil {
		return nil, false
	}
	return o.UnsafeSslHostnameCheck, true
}

// HasUnsafeSslHostnameCheck returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasUnsafeSslHostnameCheck() bool {
	if o != nil && o.UnsafeSslHostnameCheck != nil {
		return true
	}

	return false
}

// SetUnsafeSslHostnameCheck gets a reference to the given bool and assigns it to the UnsafeSslHostnameCheck field.
func (o *EngineRegistrationParameter) SetUnsafeSslHostnameCheck(v bool) {
	o.UnsafeSslHostnameCheck = &v
}

// GetTruststoreFilename returns the TruststoreFilename field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineRegistrationParameter) GetTruststoreFilename() string {
	if o == nil || o.TruststoreFilename.Get() == nil {
		var ret string
		return ret
	}
	return *o.TruststoreFilename.Get()
}

// GetTruststoreFilenameOk returns a tuple with the TruststoreFilename field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngineRegistrationParameter) GetTruststoreFilenameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TruststoreFilename.Get(), o.TruststoreFilename.IsSet()
}

// HasTruststoreFilename returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasTruststoreFilename() bool {
	if o != nil && o.TruststoreFilename.IsSet() {
		return true
	}

	return false
}

// SetTruststoreFilename gets a reference to the given NullableString and assigns it to the TruststoreFilename field.
func (o *EngineRegistrationParameter) SetTruststoreFilename(v string) {
	o.TruststoreFilename.Set(&v)
}
// SetTruststoreFilenameNil sets the value for TruststoreFilename to be an explicit nil
func (o *EngineRegistrationParameter) SetTruststoreFilenameNil() {
	o.TruststoreFilename.Set(nil)
}

// UnsetTruststoreFilename ensures that no value is present for TruststoreFilename, not even an explicit nil
func (o *EngineRegistrationParameter) UnsetTruststoreFilename() {
	o.TruststoreFilename.Unset()
}

// GetTruststorePassword returns the TruststorePassword field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EngineRegistrationParameter) GetTruststorePassword() string {
	if o == nil || o.TruststorePassword.Get() == nil {
		var ret string
		return ret
	}
	return *o.TruststorePassword.Get()
}

// GetTruststorePasswordOk returns a tuple with the TruststorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngineRegistrationParameter) GetTruststorePasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.TruststorePassword.Get(), o.TruststorePassword.IsSet()
}

// HasTruststorePassword returns a boolean if a field has been set.
func (o *EngineRegistrationParameter) HasTruststorePassword() bool {
	if o != nil && o.TruststorePassword.IsSet() {
		return true
	}

	return false
}

// SetTruststorePassword gets a reference to the given NullableString and assigns it to the TruststorePassword field.
func (o *EngineRegistrationParameter) SetTruststorePassword(v string) {
	o.TruststorePassword.Set(&v)
}
// SetTruststorePasswordNil sets the value for TruststorePassword to be an explicit nil
func (o *EngineRegistrationParameter) SetTruststorePasswordNil() {
	o.TruststorePassword.Set(nil)
}

// UnsetTruststorePassword ensures that no value is present for TruststorePassword, not even an explicit nil
func (o *EngineRegistrationParameter) UnsetTruststorePassword() {
	o.TruststorePassword.Unset()
}

func (o EngineRegistrationParameter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.HashicorpVaultUsernameCommandArgs != nil {
		toSerialize["hashicorp_vault_username_command_args"] = o.HashicorpVaultUsernameCommandArgs
	}
	if o.HashicorpVaultPasswordCommandArgs != nil {
		toSerialize["hashicorp_vault_password_command_args"] = o.HashicorpVaultPasswordCommandArgs
	}
	if o.HashicorpVaultId.IsSet() {
		toSerialize["hashicorp_vault_id"] = o.HashicorpVaultId.Get()
	}
	if o.InsecureSsl != nil {
		toSerialize["insecure_ssl"] = o.InsecureSsl
	}
	if o.UnsafeSslHostnameCheck != nil {
		toSerialize["unsafe_ssl_hostname_check"] = o.UnsafeSslHostnameCheck
	}
	if o.TruststoreFilename.IsSet() {
		toSerialize["truststore_filename"] = o.TruststoreFilename.Get()
	}
	if o.TruststorePassword.IsSet() {
		toSerialize["truststore_password"] = o.TruststorePassword.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEngineRegistrationParameter struct {
	value *EngineRegistrationParameter
	isSet bool
}

func (v NullableEngineRegistrationParameter) Get() *EngineRegistrationParameter {
	return v.value
}

func (v *NullableEngineRegistrationParameter) Set(val *EngineRegistrationParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineRegistrationParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineRegistrationParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineRegistrationParameter(val *EngineRegistrationParameter) *NullableEngineRegistrationParameter {
	return &NullableEngineRegistrationParameter{value: val, isSet: true}
}

func (v NullableEngineRegistrationParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineRegistrationParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

