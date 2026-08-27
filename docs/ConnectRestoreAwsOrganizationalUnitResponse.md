# ConnectRestoreAwsOrganizationalUnitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrganizationalUnit** | Pointer to [**RestoreAwsOrganizationalUnit**](RestoreAwsOrganizationalUnit.md) |  | [optional] 
**ActionApprovalRequest** | Pointer to [**NullableMPARequest**](MPARequest.md) |  | [optional] 

## Methods

### NewConnectRestoreAwsOrganizationalUnitResponse

`func NewConnectRestoreAwsOrganizationalUnitResponse() *ConnectRestoreAwsOrganizationalUnitResponse`

NewConnectRestoreAwsOrganizationalUnitResponse instantiates a new ConnectRestoreAwsOrganizationalUnitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectRestoreAwsOrganizationalUnitResponseWithDefaults

`func NewConnectRestoreAwsOrganizationalUnitResponseWithDefaults() *ConnectRestoreAwsOrganizationalUnitResponse`

NewConnectRestoreAwsOrganizationalUnitResponseWithDefaults instantiates a new ConnectRestoreAwsOrganizationalUnitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizationalUnit

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) GetOrganizationalUnit() RestoreAwsOrganizationalUnit`

GetOrganizationalUnit returns the OrganizationalUnit field if non-nil, zero value otherwise.

### GetOrganizationalUnitOk

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) GetOrganizationalUnitOk() (*RestoreAwsOrganizationalUnit, bool)`

GetOrganizationalUnitOk returns a tuple with the OrganizationalUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalUnit

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) SetOrganizationalUnit(v RestoreAwsOrganizationalUnit)`

SetOrganizationalUnit sets OrganizationalUnit field to given value.

### HasOrganizationalUnit

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) HasOrganizationalUnit() bool`

HasOrganizationalUnit returns a boolean if a field has been set.

### GetActionApprovalRequest

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) GetActionApprovalRequest() MPARequest`

GetActionApprovalRequest returns the ActionApprovalRequest field if non-nil, zero value otherwise.

### GetActionApprovalRequestOk

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) GetActionApprovalRequestOk() (*MPARequest, bool)`

GetActionApprovalRequestOk returns a tuple with the ActionApprovalRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionApprovalRequest

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) SetActionApprovalRequest(v MPARequest)`

SetActionApprovalRequest sets ActionApprovalRequest field to given value.

### HasActionApprovalRequest

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) HasActionApprovalRequest() bool`

HasActionApprovalRequest returns a boolean if a field has been set.

### SetActionApprovalRequestNil

`func (o *ConnectRestoreAwsOrganizationalUnitResponse) SetActionApprovalRequestNil(b bool)`

 SetActionApprovalRequestNil sets the value for ActionApprovalRequest to be an explicit nil

### UnsetActionApprovalRequest
`func (o *ConnectRestoreAwsOrganizationalUnitResponse) UnsetActionApprovalRequest()`

UnsetActionApprovalRequest ensures that no value is present for ActionApprovalRequest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


