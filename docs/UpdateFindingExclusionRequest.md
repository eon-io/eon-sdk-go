# UpdateFindingExclusionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceId** | Pointer to **string** | Eon-assigned ID of the resource to apply the exclusion to. Omit it to apply the exclusion to every resource in the account.  | [optional] 
**Value** | **string** | What the exclusion matches, depending on &#x60;type&#x60;. For &#x60;PATH&#x60;, every file whose path starts with this value, so &#x60;/data/tmp&#x60; also matches &#x60;/data/tmp2/report.csv&#x60;. Linux paths are compared case-sensitively. Windows paths are compared as the backup stores them, which is currently lowercase, so write Windows prefixes in lowercase (for example &#x60;c:/users/app/cache&#x60;). For &#x60;TABLE&#x60; or &#x60;DATABASE&#x60;, the exact table or database name.  | 
**Type** | [**FindingObjectType**](FindingObjectType.md) |  | 
**Detector** | [**FindingExclusionDetectorType**](FindingExclusionDetectorType.md) |  | 

## Methods

### NewUpdateFindingExclusionRequest

`func NewUpdateFindingExclusionRequest(value string, type_ FindingObjectType, detector FindingExclusionDetectorType, ) *UpdateFindingExclusionRequest`

NewUpdateFindingExclusionRequest instantiates a new UpdateFindingExclusionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFindingExclusionRequestWithDefaults

`func NewUpdateFindingExclusionRequestWithDefaults() *UpdateFindingExclusionRequest`

NewUpdateFindingExclusionRequestWithDefaults instantiates a new UpdateFindingExclusionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceId

`func (o *UpdateFindingExclusionRequest) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *UpdateFindingExclusionRequest) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *UpdateFindingExclusionRequest) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *UpdateFindingExclusionRequest) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetValue

`func (o *UpdateFindingExclusionRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateFindingExclusionRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateFindingExclusionRequest) SetValue(v string)`

SetValue sets Value field to given value.


### GetType

`func (o *UpdateFindingExclusionRequest) GetType() FindingObjectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateFindingExclusionRequest) GetTypeOk() (*FindingObjectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateFindingExclusionRequest) SetType(v FindingObjectType)`

SetType sets Type field to given value.


### GetDetector

`func (o *UpdateFindingExclusionRequest) GetDetector() FindingExclusionDetectorType`

GetDetector returns the Detector field if non-nil, zero value otherwise.

### GetDetectorOk

`func (o *UpdateFindingExclusionRequest) GetDetectorOk() (*FindingExclusionDetectorType, bool)`

GetDetectorOk returns a tuple with the Detector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetector

`func (o *UpdateFindingExclusionRequest) SetDetector(v FindingExclusionDetectorType)`

SetDetector sets Detector field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


