# UpdateFindingExclusionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | [**FindingExclusionScope**](FindingExclusionScope.md) |  | 
**ProviderResourceId** | Pointer to **string** | Cloud provider ID of the resource to apply the exclusion to, such as an EC2 instance ID or an Azure resource ID. Required when &#x60;scope&#x60; is &#x60;RESOURCE&#x60;, and must be omitted when it&#39;s &#x60;ACCOUNT&#x60;.  | [optional] 
**Value** | **string** | Matching value to exclude. The value set in this field depends on the specified &#x60;type&#x60;.  - &#x60;PATH&#x60;: Matches files whose path starts with the specified value, case-sensitive.   For example, &#x60;/data/tmp&#x60; matches &#x60;/data/tmp2/report.csv&#x60;.      Linux paths are case-sensitive.      Windows paths are case-insensitive, normalize for the file separator, and start with the drive letter.     For example, &#x60;C:\\Users\\App\\Cache&#x60; matches &#x60;c:/users/app/cache&#x60;.  - &#x60;TABLE&#x60;: Exact table name.  - &#x60;DATABASE&#x60;: Exact database name.  - &#x60;KMS_KEY&#x60;: Either a KMS key ARN, or an AWS account ID to trust all keys in the account.   Available for the ransomware detector only.  | 
**Type** | [**FindingObjectType**](FindingObjectType.md) |  | 
**Detector** | [**FindingExclusionDetectorType**](FindingExclusionDetectorType.md) |  | 

## Methods

### NewUpdateFindingExclusionRequest

`func NewUpdateFindingExclusionRequest(scope FindingExclusionScope, value string, type_ FindingObjectType, detector FindingExclusionDetectorType, ) *UpdateFindingExclusionRequest`

NewUpdateFindingExclusionRequest instantiates a new UpdateFindingExclusionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateFindingExclusionRequestWithDefaults

`func NewUpdateFindingExclusionRequestWithDefaults() *UpdateFindingExclusionRequest`

NewUpdateFindingExclusionRequestWithDefaults instantiates a new UpdateFindingExclusionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *UpdateFindingExclusionRequest) GetScope() FindingExclusionScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateFindingExclusionRequest) GetScopeOk() (*FindingExclusionScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateFindingExclusionRequest) SetScope(v FindingExclusionScope)`

SetScope sets Scope field to given value.


### GetProviderResourceId

`func (o *UpdateFindingExclusionRequest) GetProviderResourceId() string`

GetProviderResourceId returns the ProviderResourceId field if non-nil, zero value otherwise.

### GetProviderResourceIdOk

`func (o *UpdateFindingExclusionRequest) GetProviderResourceIdOk() (*string, bool)`

GetProviderResourceIdOk returns a tuple with the ProviderResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderResourceId

`func (o *UpdateFindingExclusionRequest) SetProviderResourceId(v string)`

SetProviderResourceId sets ProviderResourceId field to given value.

### HasProviderResourceId

`func (o *UpdateFindingExclusionRequest) HasProviderResourceId() bool`

HasProviderResourceId returns a boolean if a field has been set.

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


