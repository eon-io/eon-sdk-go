# DestinationDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestoreAccountId** | **string** | Eon-assigned restore account ID. | 
**ProviderAccountId** | **string** | Cloud-provider-assigned restore account ID. | 
**CloudProvider** | [**Provider**](Provider.md) |  | 
**Region** | **string** | Region the data is restored to. | 

## Methods

### NewDestinationDetails

`func NewDestinationDetails(restoreAccountId string, providerAccountId string, cloudProvider Provider, region string, ) *DestinationDetails`

NewDestinationDetails instantiates a new DestinationDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationDetailsWithDefaults

`func NewDestinationDetailsWithDefaults() *DestinationDetails`

NewDestinationDetailsWithDefaults instantiates a new DestinationDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestoreAccountId

`func (o *DestinationDetails) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *DestinationDetails) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *DestinationDetails) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetProviderAccountId

`func (o *DestinationDetails) GetProviderAccountId() string`

GetProviderAccountId returns the ProviderAccountId field if non-nil, zero value otherwise.

### GetProviderAccountIdOk

`func (o *DestinationDetails) GetProviderAccountIdOk() (*string, bool)`

GetProviderAccountIdOk returns a tuple with the ProviderAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderAccountId

`func (o *DestinationDetails) SetProviderAccountId(v string)`

SetProviderAccountId sets ProviderAccountId field to given value.


### GetCloudProvider

`func (o *DestinationDetails) GetCloudProvider() Provider`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *DestinationDetails) GetCloudProviderOk() (*Provider, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *DestinationDetails) SetCloudProvider(v Provider)`

SetCloudProvider sets CloudProvider field to given value.


### GetRegion

`func (o *DestinationDetails) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *DestinationDetails) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *DestinationDetails) SetRegion(v string)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


