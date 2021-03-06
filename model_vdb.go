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
	"time"
)

// VDB A Delphix virtual database or dataset.
type VDB struct {
	// The VDB object entity ID.
	Id *string `json:"id,omitempty"`
	// The database type of this VDB.
	DatabaseType NullableString `json:"database_type,omitempty"`
	// The container name of this VDB.
	Name NullableString `json:"name,omitempty"`
	// The database version of this VDB.
	DatabaseVersion NullableString `json:"database_version,omitempty"`
	// The total size of this VDB, in bytes.
	Size NullableInt64 `json:"size,omitempty"`
	// A reference to the Engine that this VDB belongs to.
	EngineId *string `json:"engine_id,omitempty"`
	// The runtime status of the VDB. 'Unknown' if all attempts to connect to the dataset failed.
	Status NullableString `json:"status,omitempty"`
	// A reference to the Environment that hosts this VDB.
	EnvironmentId NullableString `json:"environment_id,omitempty"`
	// The IP address of the VDB's host.
	IpAddress NullableString `json:"ip_address,omitempty"`
	// The FQDN of the VDB's host.
	Fqdn NullableString `json:"fqdn,omitempty"`
	// A reference to the parent dataset of this VDB.
	ParentId NullableString `json:"parent_id,omitempty"`
	// The name of the group containing this VDB.
	GroupName NullableString `json:"group_name,omitempty"`
	// The date this VDB was created.
	CreationDate NullableTime `json:"creation_date,omitempty"`
}

// NewVDB instantiates a new VDB object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVDB() *VDB {
	this := VDB{}
	return &this
}

// NewVDBWithDefaults instantiates a new VDB object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVDBWithDefaults() *VDB {
	this := VDB{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VDB) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDB) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VDB) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VDB) SetId(v string) {
	o.Id = &v
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetDatabaseType() string {
	if o == nil || o.DatabaseType.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatabaseType.Get()
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetDatabaseTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatabaseType.Get(), o.DatabaseType.IsSet()
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *VDB) HasDatabaseType() bool {
	if o != nil && o.DatabaseType.IsSet() {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given NullableString and assigns it to the DatabaseType field.
func (o *VDB) SetDatabaseType(v string) {
	o.DatabaseType.Set(&v)
}
// SetDatabaseTypeNil sets the value for DatabaseType to be an explicit nil
func (o *VDB) SetDatabaseTypeNil() {
	o.DatabaseType.Set(nil)
}

// UnsetDatabaseType ensures that no value is present for DatabaseType, not even an explicit nil
func (o *VDB) UnsetDatabaseType() {
	o.DatabaseType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *VDB) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *VDB) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *VDB) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *VDB) UnsetName() {
	o.Name.Unset()
}

// GetDatabaseVersion returns the DatabaseVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetDatabaseVersion() string {
	if o == nil || o.DatabaseVersion.Get() == nil {
		var ret string
		return ret
	}
	return *o.DatabaseVersion.Get()
}

// GetDatabaseVersionOk returns a tuple with the DatabaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetDatabaseVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.DatabaseVersion.Get(), o.DatabaseVersion.IsSet()
}

// HasDatabaseVersion returns a boolean if a field has been set.
func (o *VDB) HasDatabaseVersion() bool {
	if o != nil && o.DatabaseVersion.IsSet() {
		return true
	}

	return false
}

// SetDatabaseVersion gets a reference to the given NullableString and assigns it to the DatabaseVersion field.
func (o *VDB) SetDatabaseVersion(v string) {
	o.DatabaseVersion.Set(&v)
}
// SetDatabaseVersionNil sets the value for DatabaseVersion to be an explicit nil
func (o *VDB) SetDatabaseVersionNil() {
	o.DatabaseVersion.Set(nil)
}

// UnsetDatabaseVersion ensures that no value is present for DatabaseVersion, not even an explicit nil
func (o *VDB) UnsetDatabaseVersion() {
	o.DatabaseVersion.Unset()
}

// GetSize returns the Size field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetSize() int64 {
	if o == nil || o.Size.Get() == nil {
		var ret int64
		return ret
	}
	return *o.Size.Get()
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Size.Get(), o.Size.IsSet()
}

// HasSize returns a boolean if a field has been set.
func (o *VDB) HasSize() bool {
	if o != nil && o.Size.IsSet() {
		return true
	}

	return false
}

// SetSize gets a reference to the given NullableInt64 and assigns it to the Size field.
func (o *VDB) SetSize(v int64) {
	o.Size.Set(&v)
}
// SetSizeNil sets the value for Size to be an explicit nil
func (o *VDB) SetSizeNil() {
	o.Size.Set(nil)
}

// UnsetSize ensures that no value is present for Size, not even an explicit nil
func (o *VDB) UnsetSize() {
	o.Size.Unset()
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *VDB) GetEngineId() string {
	if o == nil || o.EngineId == nil {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDB) GetEngineIdOk() (*string, bool) {
	if o == nil || o.EngineId == nil {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *VDB) HasEngineId() bool {
	if o != nil && o.EngineId != nil {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *VDB) SetEngineId(v string) {
	o.EngineId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetStatus() string {
	if o == nil || o.Status.Get() == nil {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *VDB) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *VDB) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *VDB) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *VDB) UnsetStatus() {
	o.Status.Unset()
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetEnvironmentId() string {
	if o == nil || o.EnvironmentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EnvironmentId.Get()
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetEnvironmentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EnvironmentId.Get(), o.EnvironmentId.IsSet()
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *VDB) HasEnvironmentId() bool {
	if o != nil && o.EnvironmentId.IsSet() {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given NullableString and assigns it to the EnvironmentId field.
func (o *VDB) SetEnvironmentId(v string) {
	o.EnvironmentId.Set(&v)
}
// SetEnvironmentIdNil sets the value for EnvironmentId to be an explicit nil
func (o *VDB) SetEnvironmentIdNil() {
	o.EnvironmentId.Set(nil)
}

// UnsetEnvironmentId ensures that no value is present for EnvironmentId, not even an explicit nil
func (o *VDB) UnsetEnvironmentId() {
	o.EnvironmentId.Unset()
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetIpAddress() string {
	if o == nil || o.IpAddress.Get() == nil {
		var ret string
		return ret
	}
	return *o.IpAddress.Get()
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IpAddress.Get(), o.IpAddress.IsSet()
}

// HasIpAddress returns a boolean if a field has been set.
func (o *VDB) HasIpAddress() bool {
	if o != nil && o.IpAddress.IsSet() {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given NullableString and assigns it to the IpAddress field.
func (o *VDB) SetIpAddress(v string) {
	o.IpAddress.Set(&v)
}
// SetIpAddressNil sets the value for IpAddress to be an explicit nil
func (o *VDB) SetIpAddressNil() {
	o.IpAddress.Set(nil)
}

// UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
func (o *VDB) UnsetIpAddress() {
	o.IpAddress.Unset()
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetFqdn() string {
	if o == nil || o.Fqdn.Get() == nil {
		var ret string
		return ret
	}
	return *o.Fqdn.Get()
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetFqdnOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Fqdn.Get(), o.Fqdn.IsSet()
}

// HasFqdn returns a boolean if a field has been set.
func (o *VDB) HasFqdn() bool {
	if o != nil && o.Fqdn.IsSet() {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given NullableString and assigns it to the Fqdn field.
func (o *VDB) SetFqdn(v string) {
	o.Fqdn.Set(&v)
}
// SetFqdnNil sets the value for Fqdn to be an explicit nil
func (o *VDB) SetFqdnNil() {
	o.Fqdn.Set(nil)
}

// UnsetFqdn ensures that no value is present for Fqdn, not even an explicit nil
func (o *VDB) UnsetFqdn() {
	o.Fqdn.Unset()
}

// GetParentId returns the ParentId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetParentId() string {
	if o == nil || o.ParentId.Get() == nil {
		var ret string
		return ret
	}
	return *o.ParentId.Get()
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetParentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.ParentId.Get(), o.ParentId.IsSet()
}

// HasParentId returns a boolean if a field has been set.
func (o *VDB) HasParentId() bool {
	if o != nil && o.ParentId.IsSet() {
		return true
	}

	return false
}

// SetParentId gets a reference to the given NullableString and assigns it to the ParentId field.
func (o *VDB) SetParentId(v string) {
	o.ParentId.Set(&v)
}
// SetParentIdNil sets the value for ParentId to be an explicit nil
func (o *VDB) SetParentIdNil() {
	o.ParentId.Set(nil)
}

// UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
func (o *VDB) UnsetParentId() {
	o.ParentId.Unset()
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetGroupName() string {
	if o == nil || o.GroupName.Get() == nil {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetGroupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *VDB) HasGroupName() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *VDB) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *VDB) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *VDB) UnsetGroupName() {
	o.GroupName.Unset()
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VDB) GetCreationDate() time.Time {
	if o == nil || o.CreationDate.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.CreationDate.Get()
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VDB) GetCreationDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return o.CreationDate.Get(), o.CreationDate.IsSet()
}

// HasCreationDate returns a boolean if a field has been set.
func (o *VDB) HasCreationDate() bool {
	if o != nil && o.CreationDate.IsSet() {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given NullableTime and assigns it to the CreationDate field.
func (o *VDB) SetCreationDate(v time.Time) {
	o.CreationDate.Set(&v)
}
// SetCreationDateNil sets the value for CreationDate to be an explicit nil
func (o *VDB) SetCreationDateNil() {
	o.CreationDate.Set(nil)
}

// UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
func (o *VDB) UnsetCreationDate() {
	o.CreationDate.Unset()
}

func (o VDB) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.DatabaseType.IsSet() {
		toSerialize["database_type"] = o.DatabaseType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DatabaseVersion.IsSet() {
		toSerialize["database_version"] = o.DatabaseVersion.Get()
	}
	if o.Size.IsSet() {
		toSerialize["size"] = o.Size.Get()
	}
	if o.EngineId != nil {
		toSerialize["engine_id"] = o.EngineId
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.EnvironmentId.IsSet() {
		toSerialize["environment_id"] = o.EnvironmentId.Get()
	}
	if o.IpAddress.IsSet() {
		toSerialize["ip_address"] = o.IpAddress.Get()
	}
	if o.Fqdn.IsSet() {
		toSerialize["fqdn"] = o.Fqdn.Get()
	}
	if o.ParentId.IsSet() {
		toSerialize["parent_id"] = o.ParentId.Get()
	}
	if o.GroupName.IsSet() {
		toSerialize["group_name"] = o.GroupName.Get()
	}
	if o.CreationDate.IsSet() {
		toSerialize["creation_date"] = o.CreationDate.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableVDB struct {
	value *VDB
	isSet bool
}

func (v NullableVDB) Get() *VDB {
	return v.value
}

func (v *NullableVDB) Set(val *VDB) {
	v.value = val
	v.isSet = true
}

func (v NullableVDB) IsSet() bool {
	return v.isSet
}

func (v *NullableVDB) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVDB(val *VDB) *NullableVDB {
	return &NullableVDB{value: val, isSet: true}
}

func (v NullableVDB) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVDB) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


