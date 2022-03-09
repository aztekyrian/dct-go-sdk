# VDBGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the entity. | [readonly] 
**Name** | **string** | A unique name for the entity. | 
**VdbIds** | **[]string** | The list of VDB IDs in this VDBGroup. | 

## Methods

### NewVDBGroup

`func NewVDBGroup(id string, name string, vdbIds []string, ) *VDBGroup`

NewVDBGroup instantiates a new VDBGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVDBGroupWithDefaults

`func NewVDBGroupWithDefaults() *VDBGroup`

NewVDBGroupWithDefaults instantiates a new VDBGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VDBGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VDBGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VDBGroup) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VDBGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VDBGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VDBGroup) SetName(v string)`

SetName sets Name field to given value.


### GetVdbIds

`func (o *VDBGroup) GetVdbIds() []string`

GetVdbIds returns the VdbIds field if non-nil, zero value otherwise.

### GetVdbIdsOk

`func (o *VDBGroup) GetVdbIdsOk() (*[]string, bool)`

GetVdbIdsOk returns a tuple with the VdbIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdbIds

`func (o *VDBGroup) SetVdbIds(v []string)`

SetVdbIds sets VdbIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


