# RestoreAwsEc2InstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreAccountId** | **string** | Eon-assigned ID of the restore account. | 
**Destination** | [**AwsEc2InstanceRestoreDestination**](AwsEc2InstanceRestoreDestination.md) |  | 
**TemplateId** | Pointer to **string** | When set, the restore service re-evaluates this restore template at submit time and applies the resolved params onto the request before restoring.  | [optional] 

## Methods

### NewRestoreAwsEc2InstanceRequest

`func NewRestoreAwsEc2InstanceRequest(restoreAccountId string, destination AwsEc2InstanceRestoreDestination, ) *RestoreAwsEc2InstanceRequest`

NewRestoreAwsEc2InstanceRequest instantiates a new RestoreAwsEc2InstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreAwsEc2InstanceRequestWithDefaults

`func NewRestoreAwsEc2InstanceRequestWithDefaults() *RestoreAwsEc2InstanceRequest`

NewRestoreAwsEc2InstanceRequestWithDefaults instantiates a new RestoreAwsEc2InstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreAccountId

`func (o *RestoreAwsEc2InstanceRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *RestoreAwsEc2InstanceRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *RestoreAwsEc2InstanceRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetDestination

`func (o *RestoreAwsEc2InstanceRequest) GetDestination() AwsEc2InstanceRestoreDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *RestoreAwsEc2InstanceRequest) GetDestinationOk() (*AwsEc2InstanceRestoreDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *RestoreAwsEc2InstanceRequest) SetDestination(v AwsEc2InstanceRestoreDestination)`

SetDestination sets Destination field to given value.


### GetTemplateId

`func (o *RestoreAwsEc2InstanceRequest) GetTemplateId() string`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *RestoreAwsEc2InstanceRequest) GetTemplateIdOk() (*string, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *RestoreAwsEc2InstanceRequest) SetTemplateId(v string)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *RestoreAwsEc2InstanceRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


