# FindingExclusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Finding exclusion ID. | 
**Scope** | [**FindingExclusionScope**](FindingExclusionScope.md) |  | 
**ProviderResourceId** | Pointer to **string** | Cloud provider ID of the resource the exclusion applies to, such as an EC2 instance ID or an Azure resource ID. Present when &#x60;scope&#x60; is &#x60;RESOURCE&#x60;.  | [optional] 
**Value** | **string** | What the exclusion matches, depending on &#x60;type&#x60;. For &#x60;PATH&#x60;, every file whose path starts with this value, so &#x60;/data/tmp&#x60; also matches &#x60;/data/tmp2/report.csv&#x60;. Linux paths are compared case-sensitively. Windows paths are compared as the backup stores them, which is currently lowercase, so write Windows prefixes in lowercase (for example &#x60;c:/users/app/cache&#x60;). For &#x60;TABLE&#x60; or &#x60;DATABASE&#x60;, the exact table or database name.  | 
**Type** | [**FindingObjectType**](FindingObjectType.md) |  | 
**Detector** | [**FindingExclusionDetectorType**](FindingExclusionDetectorType.md) |  | 
**UpdatedAt** | **time.Time** | When the exclusion was created or last updated. | 

## Methods

### NewFindingExclusion

`func NewFindingExclusion(id string, scope FindingExclusionScope, value string, type_ FindingObjectType, detector FindingExclusionDetectorType, updatedAt time.Time, ) *FindingExclusion`

NewFindingExclusion instantiates a new FindingExclusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFindingExclusionWithDefaults

`func NewFindingExclusionWithDefaults() *FindingExclusion`

NewFindingExclusionWithDefaults instantiates a new FindingExclusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FindingExclusion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FindingExclusion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FindingExclusion) SetId(v string)`

SetId sets Id field to given value.


### GetScope

`func (o *FindingExclusion) GetScope() FindingExclusionScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FindingExclusion) GetScopeOk() (*FindingExclusionScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FindingExclusion) SetScope(v FindingExclusionScope)`

SetScope sets Scope field to given value.


### GetProviderResourceId

`func (o *FindingExclusion) GetProviderResourceId() string`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *FindingExclusion) GetProviderResourceIdOk() (*string, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *FindingExclusion) SetProviderResourceId(v string)`

SetProviderResourceId sets ProviderResourceId field to given value.

### HasProviderResourceId

`func (o *FindingExclusion) HasProviderResourceId() bool`

HasProviderResourceId returns a boolean if a field has been set.

### GetValue

`func (o *FindingExclusion) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FindingExclusion) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FindingExclusion) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *FindingExclusion) GetType() FindingObjectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FindingExclusion) GetTypeOk() (*FindingObjectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FindingExclusion) SetType(v FindingObjectType)`

SetType sets Type field to given value.


### GetDetector

`func (o *FindingExclusion) GetDetector() FindingExclusionDetectorType`

GetDetector returns the Detector field if non-nil, zero value otherwise.

### GetDetectorOk

`func (o *FindingExclusion) GetDetectorOk() (*FindingExclusionDetectorType, bool)`

GetDetectorOk returns a tuple with the Detector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetector

`func (o *FindingExclusion) SetDetector(v FindingExclusionDetectorType)`

SetDetector sets Detector field to given value.


### GetUpdatedAt

`func (o *FindingExclusion) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FindingExclusion) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FindingExclusion) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


