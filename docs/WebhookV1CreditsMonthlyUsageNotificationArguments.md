# WebhookV1CreditsMonthlyUsageNotificationArguments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCredits** | Pointer to **int64** | Total credits in the agreement. | [optional] 
**ExpectedMonthlyCredits** | Pointer to **int64** | Expected monthly credit usage. | [optional] 
**ActualMonthlyCredits** | Pointer to **int64** | Actual monthly credit usage. | [optional] 
**ContractMonths** | Pointer to **int64** | Length of the contract, in months. | [optional] 
**AgreementId** | Pointer to **string** | Agreement ID. | [optional] 
**ThresholdPercentage** | Pointer to **int64** | Usage threshold percentage that was crossed. | [optional] 

## Methods

### NewWebhookV1CreditsMonthlyUsageNotificationArguments

`func NewWebhookV1CreditsMonthlyUsageNotificationArguments() *WebhookV1CreditsMonthlyUsageNotificationArguments`

NewWebhookV1CreditsMonthlyUsageNotificationArguments instantiates a new WebhookV1CreditsMonthlyUsageNotificationArguments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookV1CreditsMonthlyUsageNotificationArgumentsWithDefaults

`func NewWebhookV1CreditsMonthlyUsageNotificationArgumentsWithDefaults() *WebhookV1CreditsMonthlyUsageNotificationArguments`

NewWebhookV1CreditsMonthlyUsageNotificationArgumentsWithDefaults instantiates a new WebhookV1CreditsMonthlyUsageNotificationArguments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetTotalCredits() int64`

GetTotalCredits returns the TotalCredits field if non-nil, zero value otherwise.

### GetTotalCreditsOk

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetTotalCreditsOk() (*int64, bool)`

GetTotalCreditsOk returns a tuple with the TotalCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) SetTotalCredits(v int64)`

SetTotalCredits sets TotalCredits field to given value.

### HasTotalCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) HasTotalCredits() bool`

HasTotalCredits returns a boolean if a field has been set.

### GetExpectedMonthlyCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetExpectedMonthlyCredits() int64`

GetExpectedMonthlyCredits returns the ExpectedMonthlyCredits field if non-nil, zero value otherwise.

### GetExpectedMonthlyCreditsOk

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetExpectedMonthlyCreditsOk() (*int64, bool)`

GetExpectedMonthlyCreditsOk returns a tuple with the ExpectedMonthlyCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedMonthlyCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) SetExpectedMonthlyCredits(v int64)`

SetExpectedMonthlyCredits sets ExpectedMonthlyCredits field to given value.

### HasExpectedMonthlyCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) HasExpectedMonthlyCredits() bool`

HasExpectedMonthlyCredits returns a boolean if a field has been set.

### GetActualMonthlyCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetActualMonthlyCredits() int64`

GetActualMonthlyCredits returns the ActualMonthlyCredits field if non-nil, zero value otherwise.

### GetActualMonthlyCreditsOk

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetActualMonthlyCreditsOk() (*int64, bool)`

GetActualMonthlyCreditsOk returns a tuple with the ActualMonthlyCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMonthlyCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) SetActualMonthlyCredits(v int64)`

SetActualMonthlyCredits sets ActualMonthlyCredits field to given value.

### HasActualMonthlyCredits

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) HasActualMonthlyCredits() bool`

HasActualMonthlyCredits returns a boolean if a field has been set.

### GetContractMonths

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetContractMonths() int64`

GetContractMonths returns the ContractMonths field if non-nil, zero value otherwise.

### GetContractMonthsOk

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetContractMonthsOk() (*int64, bool)`

GetContractMonthsOk returns a tuple with the ContractMonths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractMonths

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) SetContractMonths(v int64)`

SetContractMonths sets ContractMonths field to given value.

### HasContractMonths

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) HasContractMonths() bool`

HasContractMonths returns a boolean if a field has been set.

### GetAgreementId

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetAgreementId() string`

GetAgreementId returns the AgreementId field if non-nil, zero value otherwise.

### GetAgreementIdOk

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetAgreementIdOk() (*string, bool)`

GetAgreementIdOk returns a tuple with the AgreementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreementId

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) SetAgreementId(v string)`

SetAgreementId sets AgreementId field to given value.

### HasAgreementId

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) HasAgreementId() bool`

HasAgreementId returns a boolean if a field has been set.

### GetThresholdPercentage

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetThresholdPercentage() int64`

GetThresholdPercentage returns the ThresholdPercentage field if non-nil, zero value otherwise.

### GetThresholdPercentageOk

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) GetThresholdPercentageOk() (*int64, bool)`

GetThresholdPercentageOk returns a tuple with the ThresholdPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdPercentage

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) SetThresholdPercentage(v int64)`

SetThresholdPercentage sets ThresholdPercentage field to given value.

### HasThresholdPercentage

`func (o *WebhookV1CreditsMonthlyUsageNotificationArguments) HasThresholdPercentage() bool`

HasThresholdPercentage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


