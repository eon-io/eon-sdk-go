# GcpSourceAccountAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccount** | **string** |  | 
**OrganizationId** | Pointer to **string** | Cloud-provider-assigned ID of the GCP project&#39;s parent organization. | [optional] 
**FolderId** | Pointer to **string** | Cloud-provider-assigned ID of the GCP project&#39;s parent folder. | [optional] 

## Methods

### NewGcpSourceAccountAttributes

`func NewGcpSourceAccountAttributes(serviceAccount string, ) *GcpSourceAccountAttributes`

NewGcpSourceAccountAttributes instantiates a new GcpSourceAccountAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcpSourceAccountAttributesWithDefaults

`func NewGcpSourceAccountAttributesWithDefaults() *GcpSourceAccountAttributes`

NewGcpSourceAccountAttributesWithDefaults instantiates a new GcpSourceAccountAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccount

`func (o *GcpSourceAccountAttributes) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *GcpSourceAccountAttributes) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *GcpSourceAccountAttributes) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.


### GetOrganizationId

`func (o *GcpSourceAccountAttributes) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *GcpSourceAccountAttributes) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *GcpSourceAccountAttributes) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *GcpSourceAccountAttributes) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetFolderId

`func (o *GcpSourceAccountAttributes) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *GcpSourceAccountAttributes) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *GcpSourceAccountAttributes) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *GcpSourceAccountAttributes) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


