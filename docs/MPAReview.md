# MPAReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the review. | 
**ReviewerEmail** | **string** | Email address of the reviewer who submitted this review. | 
**ReviewerName** | Pointer to **string** | Display name of the reviewer at the time the review was submitted. | [optional] 
**Decision** | [**MPADecision**](MPADecision.md) |  | 
**Comment** | Pointer to **string** | Optional reviewer comment explaining the decision. | [optional] 
**ReviewedTime** | **time.Time** | Time when the review was submitted. | 

## Methods

### NewMPAReview

`func NewMPAReview(id string, reviewerEmail string, decision MPADecision, reviewedTime time.Time, ) *MPAReview`

NewMPAReview instantiates a new MPAReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPAReviewWithDefaults

`func NewMPAReviewWithDefaults() *MPAReview`

NewMPAReviewWithDefaults instantiates a new MPAReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MPAReview) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MPAReview) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MPAReview) SetId(v string)`

SetId sets Id field to given value.


### GetReviewerEmail

`func (o *MPAReview) GetReviewerEmail() string`

GetReviewerEmail returns the ReviewerEmail field if non-nil, zero value otherwise.

### GetReviewerEmailOk

`func (o *MPAReview) GetReviewerEmailOk() (*string, bool)`

GetReviewerEmailOk returns a tuple with the ReviewerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerEmail

`func (o *MPAReview) SetReviewerEmail(v string)`

SetReviewerEmail sets ReviewerEmail field to given value.


### GetReviewerName

`func (o *MPAReview) GetReviewerName() string`

GetReviewerName returns the ReviewerName field if non-nil, zero value otherwise.

### GetReviewerNameOk

`func (o *MPAReview) GetReviewerNameOk() (*string, bool)`

GetReviewerNameOk returns a tuple with the ReviewerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewerName

`func (o *MPAReview) SetReviewerName(v string)`

SetReviewerName sets ReviewerName field to given value.

### HasReviewerName

`func (o *MPAReview) HasReviewerName() bool`

HasReviewerName returns a boolean if a field has been set.

### GetDecision

`func (o *MPAReview) GetDecision() MPADecision`

GetDecision returns the Decision field if non-nil, zero value otherwise.

### GetDecisionOk

`func (o *MPAReview) GetDecisionOk() (*MPADecision, bool)`

GetDecisionOk returns a tuple with the Decision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecision

`func (o *MPAReview) SetDecision(v MPADecision)`

SetDecision sets Decision field to given value.


### GetComment

`func (o *MPAReview) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *MPAReview) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *MPAReview) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *MPAReview) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetReviewedTime

`func (o *MPAReview) GetReviewedTime() time.Time`

GetReviewedTime returns the ReviewedTime field if non-nil, zero value otherwise.

### GetReviewedTimeOk

`func (o *MPAReview) GetReviewedTimeOk() (*time.Time, bool)`

GetReviewedTimeOk returns a tuple with the ReviewedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedTime

`func (o *MPAReview) SetReviewedTime(v time.Time)`

SetReviewedTime sets ReviewedTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


