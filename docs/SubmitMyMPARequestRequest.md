# SubmitMyMPARequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | [**MPASubmitAction**](MPASubmitAction.md) |  | 
**Comment** | Pointer to **string** | Optional comment to set on the request alongside the action. | [optional] 

## Methods

### NewSubmitMyMPARequestRequest

`func NewSubmitMyMPARequestRequest(action MPASubmitAction, ) *SubmitMyMPARequestRequest`

NewSubmitMyMPARequestRequest instantiates a new SubmitMyMPARequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitMyMPARequestRequestWithDefaults

`func NewSubmitMyMPARequestRequestWithDefaults() *SubmitMyMPARequestRequest`

NewSubmitMyMPARequestRequestWithDefaults instantiates a new SubmitMyMPARequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SubmitMyMPARequestRequest) GetAction() MPASubmitAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SubmitMyMPARequestRequest) GetActionOk() (*MPASubmitAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SubmitMyMPARequestRequest) SetAction(v MPASubmitAction)`

SetAction sets Action field to given value.


### GetComment

`func (o *SubmitMyMPARequestRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *SubmitMyMPARequestRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *SubmitMyMPARequestRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *SubmitMyMPARequestRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


