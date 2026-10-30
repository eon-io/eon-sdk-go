# RestoreTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of the restore template. | 
**Name** | **string** | Restore template display name. | 
**Description** | Pointer to **string** | Description of the restore template&#39;s purpose. | [optional] 
**ResourceType** | [**RestoreTemplateResourceType**](RestoreTemplateResourceType.md) |  | 
**Ec2Params** | Pointer to [**NullableEc2TemplateParams**](Ec2TemplateParams.md) |  | [optional] 
**S3Params** | Pointer to [**NullableS3TemplateParams**](S3TemplateParams.md) |  | [optional] 
**AttachedResourceCount** | Pointer to **int32** | Number of resources currently using this restore template. | [optional] 
**CreatedAt** | Pointer to **time.Time** | Date and time the restore template was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | Date and time the restore template was last updated. | [optional] 

## Methods

### NewRestoreTemplate

`func NewRestoreTemplate(id string, name string, resourceType RestoreTemplateResourceType, ) *RestoreTemplate`

NewRestoreTemplate instantiates a new RestoreTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreTemplateWithDefaults

`func NewRestoreTemplateWithDefaults() *RestoreTemplate`

NewRestoreTemplateWithDefaults instantiates a new RestoreTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RestoreTemplate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RestoreTemplate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RestoreTemplate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RestoreTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RestoreTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RestoreTemplate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *RestoreTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RestoreTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RestoreTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RestoreTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetResourceType

`func (o *RestoreTemplate) GetResourceType() RestoreTemplateResourceType`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *RestoreTemplate) GetResourceTypeOk() (*RestoreTemplateResourceType, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *RestoreTemplate) SetResourceType(v RestoreTemplateResourceType)`

SetResourceType sets ResourceType field to given value.


### GetEc2Params

`func (o *RestoreTemplate) GetEc2Params() Ec2TemplateParams`

GetEc2Params returns the Ec2Params field if non-nil, zero value otherwise.

### GetEc2ParamsOk

`func (o *RestoreTemplate) GetEc2ParamsOk() (*Ec2TemplateParams, bool)`

GetEc2ParamsOk returns a tuple with the Ec2Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEc2Params

`func (o *RestoreTemplate) SetEc2Params(v Ec2TemplateParams)`

SetEc2Params sets Ec2Params field to given value.

### HasEc2Params

`func (o *RestoreTemplate) HasEc2Params() bool`

HasEc2Params returns a boolean if a field has been set.

### SetEc2ParamsNil

`func (o *RestoreTemplate) SetEc2ParamsNil(b bool)`

 SetEc2ParamsNil sets the value for Ec2Params to be an explicit nil

### UnsetEc2Params
`func (o *RestoreTemplate) UnsetEc2Params()`

UnsetEc2Params ensures that no value is present for Ec2Params, not even an explicit nil
### GetS3Params

`func (o *RestoreTemplate) GetS3Params() S3TemplateParams`

GetS3Params returns the S3Params field if non-nil, zero value otherwise.

### GetS3ParamsOk

`func (o *RestoreTemplate) GetS3ParamsOk() (*S3TemplateParams, bool)`

GetS3ParamsOk returns a tuple with the S3Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3Params

`func (o *RestoreTemplate) SetS3Params(v S3TemplateParams)`

SetS3Params sets S3Params field to given value.

### HasS3Params

`func (o *RestoreTemplate) HasS3Params() bool`

HasS3Params returns a boolean if a field has been set.

### SetS3ParamsNil

`func (o *RestoreTemplate) SetS3ParamsNil(b bool)`

 SetS3ParamsNil sets the value for S3Params to be an explicit nil

### UnsetS3Params
`func (o *RestoreTemplate) UnsetS3Params()`

UnsetS3Params ensures that no value is present for S3Params, not even an explicit nil
### GetAttachedResourceCount

`func (o *RestoreTemplate) GetAttachedResourceCount() int32`

GetAttachedResourceCount returns the AttachedResourceCount field if non-nil, zero value otherwise.

### GetAttachedResourceCountOk

`func (o *RestoreTemplate) GetAttachedResourceCountOk() (*int32, bool)`

GetAttachedResourceCountOk returns a tuple with the AttachedResourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachedResourceCount

`func (o *RestoreTemplate) SetAttachedResourceCount(v int32)`

SetAttachedResourceCount sets AttachedResourceCount field to given value.

### HasAttachedResourceCount

`func (o *RestoreTemplate) HasAttachedResourceCount() bool`

HasAttachedResourceCount returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RestoreTemplate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RestoreTemplate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RestoreTemplate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RestoreTemplate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *RestoreTemplate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *RestoreTemplate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *RestoreTemplate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *RestoreTemplate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


