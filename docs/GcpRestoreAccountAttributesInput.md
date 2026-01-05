# GcpRestoreAccountAttributesInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccount** | **string** | Email address of the GCP service account Eon uses to access the project for restore operations. | 

## Methods

### NewGcpRestoreAccountAttributesInput

`func NewGcpRestoreAccountAttributesInput(serviceAccount string, ) *GcpRestoreAccountAttributesInput`

NewGcpRestoreAccountAttributesInput instantiates a new GcpRestoreAccountAttributesInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpRestoreAccountAttributesInputWithDefaults

`func NewGcpRestoreAccountAttributesInputWithDefaults() *GcpRestoreAccountAttributesInput`

NewGcpRestoreAccountAttributesInputWithDefaults instantiates a new GcpRestoreAccountAttributesInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccount

`func (o *GcpRestoreAccountAttributesInput) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *GcpRestoreAccountAttributesInput) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *GcpRestoreAccountAttributesInput) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


