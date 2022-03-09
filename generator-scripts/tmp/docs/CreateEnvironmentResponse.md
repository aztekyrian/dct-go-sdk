# CreateEnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | The initiated job id. | [optional] 
**EnvironmentId** | Pointer to **string** | The id of environment created. | [optional] 

## Methods

### NewCreateEnvironmentResponse

`func NewCreateEnvironmentResponse() *CreateEnvironmentResponse`

NewCreateEnvironmentResponse instantiates a new CreateEnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateEnvironmentResponseWithDefaults

`func NewCreateEnvironmentResponseWithDefaults() *CreateEnvironmentResponse`

NewCreateEnvironmentResponseWithDefaults instantiates a new CreateEnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *CreateEnvironmentResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CreateEnvironmentResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CreateEnvironmentResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *CreateEnvironmentResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *CreateEnvironmentResponse) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *CreateEnvironmentResponse) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *CreateEnvironmentResponse) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *CreateEnvironmentResponse) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


