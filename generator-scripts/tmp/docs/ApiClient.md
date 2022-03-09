# ApiClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The entity ID of this API client. | [optional] [readonly] 
**ApiClientId** | Pointer to **string** | The unique ID which is used to identity the identity of an API request. The web server (nginx) configuration must be configured so as to include the external ID as the value of the X_CLIENT_ID HTTP request header when requests are proxied. For OAuth2/JWT based authentication, this typically corresponds to a value extracted from the JWT, uniquely identifying the API client. | [optional] 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewApiClient

`func NewApiClient() *ApiClient`

NewApiClient instantiates a new ApiClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiClientWithDefaults

`func NewApiClientWithDefaults() *ApiClient`

NewApiClientWithDefaults instantiates a new ApiClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiClient) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiClient) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiClient) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ApiClient) HasId() bool`

HasId returns a boolean if a field has been set.

### GetApiClientId

`func (o *ApiClient) GetApiClientId() string`

GetApiClientId returns the ApiClientId field if non-nil, zero value otherwise.

### GetApiClientIdOk

`func (o *ApiClient) GetApiClientIdOk() (*string, bool)`

GetApiClientIdOk returns a tuple with the ApiClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiClientId

`func (o *ApiClient) SetApiClientId(v string)`

SetApiClientId sets ApiClientId field to given value.

### HasApiClientId

`func (o *ApiClient) HasApiClientId() bool`

HasApiClientId returns a boolean if a field has been set.

### GetName

`func (o *ApiClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiClient) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


