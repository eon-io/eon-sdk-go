# RestoreDestinationTagLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subject** | [**RestoreAccountTagSubject**](RestoreAccountTagSubject.md) |  | 
**Operator** | [**RestoreAccountTagOperator**](RestoreAccountTagOperator.md) |  | 
**Values** | **[]string** | Tag keys when &#x60;subject&#x60; is &#x60;TAG_KEYS&#x60;, &#x60;key&#x3D;value&#x60; pairs when it is &#x60;TAG_KEY_VALUES&#x60;.  | 

## Methods

### NewRestoreDestinationTagLimits

`func NewRestoreDestinationTagLimits(subject RestoreAccountTagSubject, operator RestoreAccountTagOperator, values []string, ) *RestoreDestinationTagLimits`

NewRestoreDestinationTagLimits instantiates a new RestoreDestinationTagLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreDestinationTagLimitsWithDefaults

`func NewRestoreDestinationTagLimitsWithDefaults() *RestoreDestinationTagLimits`

NewRestoreDestinationTagLimitsWithDefaults instantiates a new RestoreDestinationTagLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubject

`func (o *RestoreDestinationTagLimits) GetSubject() RestoreAccountTagSubject`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RestoreDestinationTagLimits) GetSubjectOk() (*RestoreAccountTagSubject, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RestoreDestinationTagLimits) SetSubject(v RestoreAccountTagSubject)`

SetSubject sets Subject field to given value.


### GetOperator

`func (o *RestoreDestinationTagLimits) GetOperator() RestoreAccountTagOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *RestoreDestinationTagLimits) GetOperatorOk() (*RestoreAccountTagOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *RestoreDestinationTagLimits) SetOperator(v RestoreAccountTagOperator)`

SetOperator sets Operator field to given value.


### GetValues

`func (o *RestoreDestinationTagLimits) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RestoreDestinationTagLimits) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RestoreDestinationTagLimits) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


