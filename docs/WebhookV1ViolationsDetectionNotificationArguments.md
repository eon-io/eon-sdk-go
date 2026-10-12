# WebhookV1ViolationsDetectionNotificationArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceIds** | Pointer to **[]string** | IDs of the affected resources. | [optional] 
**TotalResources** | Pointer to **int32** | Total number of affected resources. | [optional] 
**ControlName** | Pointer to **string** | Name of the control. | [optional] 
**Detected** | Pointer to **bool** | Whether the violations were detected. Omitted when they were resolved. | [optional] 

## Methods

### NewWebhookV1ViolationsDetectionNotificationArguments

`func NewWebhookV1ViolationsDetectionNotificationArguments() *WebhookV1ViolationsDetectionNotificationArguments`

NewWebhookV1ViolationsDetectionNotificationArguments instantiates a new WebhookV1ViolationsDetectionNotificationArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1ViolationsDetectionNotificationArgumentsWithDefaults

`func NewWebhookV1ViolationsDetectionNotificationArgumentsWithDefaults() *WebhookV1ViolationsDetectionNotificationArguments`

NewWebhookV1ViolationsDetectionNotificationArgumentsWithDefaults instantiates a new WebhookV1ViolationsDetectionNotificationArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceIds

`func (o *WebhookV1ViolationsDetectionNotificationArguments) GetResourceIds() []string`

GetResourceIds returns the ResourceIds field if non-nil, zero value otherwise.

### GetResourceIdsOk

`func (o *WebhookV1ViolationsDetectionNotificationArguments) GetResourceIdsOk() (*[]string, bool)`

GetResourceIdsOk returns a tuple with the ResourceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIds

`func (o *WebhookV1ViolationsDetectionNotificationArguments) SetResourceIds(v []string)`

SetResourceIds sets ResourceIds field to given value.

### HasResourceIds

`func (o *WebhookV1ViolationsDetectionNotificationArguments) HasResourceIds() bool`

HasResourceIds returns a boolean if a field has been set.

### GetTotalResources

`func (o *WebhookV1ViolationsDetectionNotificationArguments) GetTotalResources() int32`

GetTotalResources returns the TotalResources field if non-nil, zero value otherwise.

### GetTotalResourcesOk

`func (o *WebhookV1ViolationsDetectionNotificationArguments) GetTotalResourcesOk() (*int32, bool)`

GetTotalResourcesOk returns a tuple with the TotalResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalResources

`func (o *WebhookV1ViolationsDetectionNotificationArguments) SetTotalResources(v int32)`

SetTotalResources sets TotalResources field to given value.

### HasTotalResources

`func (o *WebhookV1ViolationsDetectionNotificationArguments) HasTotalResources() bool`

HasTotalResources returns a boolean if a field has been set.

### GetControlName

`func (o *WebhookV1ViolationsDetectionNotificationArguments) GetControlName() string`

GetControlName returns the ControlName field if non-nil, zero value otherwise.

### GetControlNameOk

`func (o *WebhookV1ViolationsDetectionNotificationArguments) GetControlNameOk() (*string, bool)`

GetControlNameOk returns a tuple with the ControlName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlName

`func (o *WebhookV1ViolationsDetectionNotificationArguments) SetControlName(v string)`

SetControlName sets ControlName field to given value.

### HasControlName

`func (o *WebhookV1ViolationsDetectionNotificationArguments) HasControlName() bool`

HasControlName returns a boolean if a field has been set.

### GetDetected

`func (o *WebhookV1ViolationsDetectionNotificationArguments) GetDetected() bool`

GetDetected returns the Detected field if non-nil, zero value otherwise.

### GetDetectedOk

`func (o *WebhookV1ViolationsDetectionNotificationArguments) GetDetectedOk() (*bool, bool)`

GetDetectedOk returns a tuple with the Detected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetected

`func (o *WebhookV1ViolationsDetectionNotificationArguments) SetDetected(v bool)`

SetDetected sets Detected field to given value.

### HasDetected

`func (o *WebhookV1ViolationsDetectionNotificationArguments) HasDetected() bool`

HasDetected returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


