# DestinationProviderAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Azure** | Pointer to [**NullableDestinationProviderAttributesAzure**](DestinationProviderAttributesAzure.md) |  | [optional] 

## Methods

### NewDestinationProviderAttributes

`func NewDestinationProviderAttributes() *DestinationProviderAttributes`

NewDestinationProviderAttributes instantiates a new DestinationProviderAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationProviderAttributesWithDefaults

`func NewDestinationProviderAttributesWithDefaults() *DestinationProviderAttributes`

NewDestinationProviderAttributesWithDefaults instantiates a new DestinationProviderAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzure

`func (o *DestinationProviderAttributes) GetAzure() DestinationProviderAttributesAzure`

GetAzure returns the Azure field if non-nil, zero value otherwise.

### GetAzureOk

`func (o *DestinationProviderAttributes) GetAzureOk() (*DestinationProviderAttributesAzure, bool)`

GetAzureOk returns a tuple with the Azure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzure

`func (o *DestinationProviderAttributes) SetAzure(v DestinationProviderAttributesAzure)`

SetAzure sets Azure field to given value.

### HasAzure

`func (o *DestinationProviderAttributes) HasAzure() bool`

HasAzure returns a boolean if a field has been set.

### SetAzureNil

`func (o *DestinationProviderAttributes) SetAzureNil(b bool)`

 SetAzureNil sets the value for Azure to be an explicit nil

### UnsetAzure
`func (o *DestinationProviderAttributes) UnsetAzure()`

UnsetAzure ensures that no value is present for Azure, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


