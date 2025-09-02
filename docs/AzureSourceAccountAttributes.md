# AzureSourceAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | ID of the Azure tenant the subscription belongs to. | 

## Methods

### NewAzureSourceAccountAttributes

`func NewAzureSourceAccountAttributes(tenantId string, ) *AzureSourceAccountAttributes`

NewAzureSourceAccountAttributes instantiates a new AzureSourceAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureSourceAccountAttributesWithDefaults

`func NewAzureSourceAccountAttributesWithDefaults() *AzureSourceAccountAttributes`

NewAzureSourceAccountAttributesWithDefaults instantiates a new AzureSourceAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *AzureSourceAccountAttributes) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AzureSourceAccountAttributes) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AzureSourceAccountAttributes) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


