# CloudAccountIdFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **[]string** | Matches if any value in this list equals the account ID. | [optional] 
**NotIn** | Pointer to **[]string** | Matches if none of the values in this list equal the account ID. | [optional] 

## Methods

### NewCloudAccountIdFilters

`func NewCloudAccountIdFilters() *CloudAccountIdFilters`

NewCloudAccountIdFilters instantiates a new CloudAccountIdFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountIdFiltersWithDefaults

`func NewCloudAccountIdFiltersWithDefaults() *CloudAccountIdFilters`

NewCloudAccountIdFiltersWithDefaults instantiates a new CloudAccountIdFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *CloudAccountIdFilters) GetIn() []string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *CloudAccountIdFilters) GetInOk() (*[]string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *CloudAccountIdFilters) SetIn(v []string)`

SetIn sets In field to given value.

### HasIn

`func (o *CloudAccountIdFilters) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetNotIn

`func (o *CloudAccountIdFilters) GetNotIn() []string`

GetNotIn returns the NotIn field if non-nil, zero value otherwise.

### GetNotInOk

`func (o *CloudAccountIdFilters) GetNotInOk() (*[]string, bool)`

GetNotInOk returns a tuple with the NotIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotIn

`func (o *CloudAccountIdFilters) SetNotIn(v []string)`

SetNotIn sets NotIn field to given value.

### HasNotIn

`func (o *CloudAccountIdFilters) HasNotIn() bool`

HasNotIn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


