# ListAwsVpcsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vpcs** | [**[]ListAwsVpcsResponseVpcsInner**](ListAwsVpcsResponseVpcsInner.md) | List of VPCs. | 

## Methods

### NewListAwsVpcsResponse

`func NewListAwsVpcsResponse(vpcs []ListAwsVpcsResponseVpcsInner, ) *ListAwsVpcsResponse`

NewListAwsVpcsResponse instantiates a new ListAwsVpcsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAwsVpcsResponseWithDefaults

`func NewListAwsVpcsResponseWithDefaults() *ListAwsVpcsResponse`

NewListAwsVpcsResponseWithDefaults instantiates a new ListAwsVpcsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVpcs

`func (o *ListAwsVpcsResponse) GetVpcs() []ListAwsVpcsResponseVpcsInner`

GetVpcs returns the Vpcs field if non-nil, zero value otherwise.

### GetVpcsOk

`func (o *ListAwsVpcsResponse) GetVpcsOk() (*[]ListAwsVpcsResponseVpcsInner, bool)`

GetVpcsOk returns a tuple with the Vpcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpcs

`func (o *ListAwsVpcsResponse) SetVpcs(v []ListAwsVpcsResponseVpcsInner)`

SetVpcs sets Vpcs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


