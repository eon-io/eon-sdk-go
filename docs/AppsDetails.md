# AppsDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]App**](App.md) | List of app details. | [optional] 
**IsOverridden** | Pointer to **bool** | Whether the apps are manually overridden. If &#x60;true&#x60;, the list of apps is user-defined and remains static. If &#x60;false&#x60;, the apps are automatically detected and listed by Eon.  | [optional] 

## Methods

### NewAppsDetails

`func NewAppsDetails() *AppsDetails`

NewAppsDetails instantiates a new AppsDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsDetailsWithDefaults

`func NewAppsDetailsWithDefaults() *AppsDetails`

NewAppsDetailsWithDefaults instantiates a new AppsDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *AppsDetails) GetApps() []App`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AppsDetails) GetAppsOk() (*[]App, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AppsDetails) SetApps(v []App)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AppsDetails) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetIsOverridden

`func (o *AppsDetails) GetIsOverridden() bool`

GetIsOverridden returns the IsOverridden field if non-nil, zero value otherwise.

### GetIsOverriddenOk

`func (o *AppsDetails) GetIsOverriddenOk() (*bool, bool)`

GetIsOverriddenOk returns a tuple with the IsOverridden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOverridden

`func (o *AppsDetails) SetIsOverridden(v bool)`

SetIsOverridden sets IsOverridden field to given value.

### HasIsOverridden

`func (o *AppsDetails) HasIsOverridden() bool`

HasIsOverridden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


