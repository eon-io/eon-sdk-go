# CreateRestoreTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Restore template display name. | 
**Description** | Pointer to **string** | Optional description of the restore template. | [optional] 
**ResourceType** | [**RestoreTemplateResourceType**](RestoreTemplateResourceType.md) |  | 
**Ec2Params** | Pointer to [**NullableEc2TemplateParams**](Ec2TemplateParams.md) |  | [optional] 
**S3Params** | Pointer to [**NullableS3TemplateParams**](S3TemplateParams.md) |  | [optional] 

## Methods

### NewCreateRestoreTemplateRequest

`func NewCreateRestoreTemplateRequest(name string, resourceType RestoreTemplateResourceType, ) *CreateRestoreTemplateRequest`

NewCreateRestoreTemplateRequest instantiates a new CreateRestoreTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRestoreTemplateRequestWithDefaults

`func NewCreateRestoreTemplateRequestWithDefaults() *CreateRestoreTemplateRequest`

NewCreateRestoreTemplateRequestWithDefaults instantiates a new CreateRestoreTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRestoreTemplateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRestoreTemplateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRestoreTemplateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRestoreTemplateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRestoreTemplateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRestoreTemplateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRestoreTemplateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceType

`func (o *CreateRestoreTemplateRequest) GetResourceType() RestoreTemplateResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *CreateRestoreTemplateRequest) GetResourceTypeOk() (*RestoreTemplateResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *CreateRestoreTemplateRequest) SetResourceType(v RestoreTemplateResourceType)`

SetResourceType sets ResourceType field to given value.


### GetEc2Params

`func (o *CreateRestoreTemplateRequest) GetEc2Params() Ec2TemplateParams`

GetEc2Params returns the Ec2Params field if non-nil, zero value otherwise.

### GetEc2ParamsOk

`func (o *CreateRestoreTemplateRequest) GetEc2ParamsOk() (*Ec2TemplateParams, bool)`

GetEc2ParamsOk returns a tuple with the Ec2Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2Params

`func (o *CreateRestoreTemplateRequest) SetEc2Params(v Ec2TemplateParams)`

SetEc2Params sets Ec2Params field to given value.

### HasEc2Params

`func (o *CreateRestoreTemplateRequest) HasEc2Params() bool`

HasEc2Params returns a boolean if a field has been set.

### SetEc2ParamsNil

`func (o *CreateRestoreTemplateRequest) SetEc2ParamsNil(b bool)`

 SetEc2ParamsNil sets the value for Ec2Params to be an explicit nil

### UnsetEc2Params
`func (o *CreateRestoreTemplateRequest) UnsetEc2Params()`

UnsetEc2Params ensures that no value is present for Ec2Params, not even an explicit nil
### GetS3Params

`func (o *CreateRestoreTemplateRequest) GetS3Params() S3TemplateParams`

GetS3Params returns the S3Params field if non-nil, zero value otherwise.

### GetS3ParamsOk

`func (o *CreateRestoreTemplateRequest) GetS3ParamsOk() (*S3TemplateParams, bool)`

GetS3ParamsOk returns a tuple with the S3Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Params

`func (o *CreateRestoreTemplateRequest) SetS3Params(v S3TemplateParams)`

SetS3Params sets S3Params field to given value.

### HasS3Params

`func (o *CreateRestoreTemplateRequest) HasS3Params() bool`

HasS3Params returns a boolean if a field has been set.

### SetS3ParamsNil

`func (o *CreateRestoreTemplateRequest) SetS3ParamsNil(b bool)`

 SetS3ParamsNil sets the value for S3Params to be an explicit nil

### UnsetS3Params
`func (o *CreateRestoreTemplateRequest) UnsetS3Params()`

UnsetS3Params ensures that no value is present for S3Params, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


