# MPARequestDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | [**NullableMPARequest**](MPARequest.md) |  | 
**Reviews** | [**[]MPAReview**](MPAReview.md) |  | 

## Methods

### NewMPARequestDetail

`func NewMPARequestDetail(request NullableMPARequest, reviews []MPAReview, ) *MPARequestDetail`

NewMPARequestDetail instantiates a new MPARequestDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMPARequestDetailWithDefaults

`func NewMPARequestDetailWithDefaults() *MPARequestDetail`

NewMPARequestDetailWithDefaults instantiates a new MPARequestDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *MPARequestDetail) GetRequest() MPARequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *MPARequestDetail) GetRequestOk() (*MPARequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *MPARequestDetail) SetRequest(v MPARequest)`

SetRequest sets Request field to given value.


### SetRequestNil

`func (o *MPARequestDetail) SetRequestNil(b bool)`

 SetRequestNil sets the value for Request to be an explicit nil

### UnsetRequest
`func (o *MPARequestDetail) UnsetRequest()`

UnsetRequest ensures that no value is present for Request, not even an explicit nil
### GetReviews

`func (o *MPARequestDetail) GetReviews() []MPAReview`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *MPARequestDetail) GetReviewsOk() (*[]MPAReview, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *MPARequestDetail) SetReviews(v []MPAReview)`

SetReviews sets Reviews field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


