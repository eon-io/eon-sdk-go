# GcpSourceAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccount** | **string** | Email address of the GCP service account Eon uses to access the project. | 

## Methods

### NewGcpSourceAccountAttributesInput

`func NewGcpSourceAccountAttributesInput(serviceAccount string, ) *GcpSourceAccountAttributesInput`

NewGcpSourceAccountAttributesInput instantiates a new GcpSourceAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpSourceAccountAttributesInputWithDefaults

`func NewGcpSourceAccountAttributesInputWithDefaults() *GcpSourceAccountAttributesInput`

NewGcpSourceAccountAttributesInputWithDefaults instantiates a new GcpSourceAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccount

`func (o *GcpSourceAccountAttributesInput) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *GcpSourceAccountAttributesInput) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *GcpSourceAccountAttributesInput) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


