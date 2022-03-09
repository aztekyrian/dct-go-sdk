# EngineConnectivityCheckRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EngineId** | **string** |  | 
**Host** | **string** |  | 
**Port** | **NullableInt32** |  | 

## Methods

### NewEngineConnectivityCheckRequest

`func NewEngineConnectivityCheckRequest(engineId string, host string, port NullableInt32, ) *EngineConnectivityCheckRequest`

NewEngineConnectivityCheckRequest instantiates a new EngineConnectivityCheckRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineConnectivityCheckRequestWithDefaults

`func NewEngineConnectivityCheckRequestWithDefaults() *EngineConnectivityCheckRequest`

NewEngineConnectivityCheckRequestWithDefaults instantiates a new EngineConnectivityCheckRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEngineId

`func (o *EngineConnectivityCheckRequest) GetEngineId() string`

GetEngineId returns the EngineId field if non-nil, zero value otherwise.

### GetEngineIdOk

`func (o *EngineConnectivityCheckRequest) GetEngineIdOk() (*string, bool)`

GetEngineIdOk returns a tuple with the EngineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineId

`func (o *EngineConnectivityCheckRequest) SetEngineId(v string)`

SetEngineId sets EngineId field to given value.


### GetHost

`func (o *EngineConnectivityCheckRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *EngineConnectivityCheckRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *EngineConnectivityCheckRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *EngineConnectivityCheckRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *EngineConnectivityCheckRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *EngineConnectivityCheckRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### SetPortNil

`func (o *EngineConnectivityCheckRequest) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *EngineConnectivityCheckRequest) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


