# QueryDBRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatabaseName** | **string** | Name of the database to query. | 
**Query** | **string** | SQL query to run against the database. Support is limited to Athena [&#x60;SELECT&#x60; queries](https://docs.aws.amazon.com/athena/latest/ug/select.html).  | 
**RestoreAccountId** | **string** | Eon-assigned ID of the [restore account](./list-restore-accounts). | 
**Destination** | [**QueryDBRestoreDestination**](QueryDBRestoreDestination.md) |  | 

## Methods

### NewQueryDBRequest

`func NewQueryDBRequest(databaseName string, query string, restoreAccountId string, destination QueryDBRestoreDestination, ) *QueryDBRequest`

NewQueryDBRequest instantiates a new QueryDBRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryDBRequestWithDefaults

`func NewQueryDBRequestWithDefaults() *QueryDBRequest`

NewQueryDBRequestWithDefaults instantiates a new QueryDBRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabaseName

`func (o *QueryDBRequest) GetDatabaseName() string`

GetDatabaseName returns the DatabaseName field if non-nil, zero value otherwise.

### GetDatabaseNameOk

`func (o *QueryDBRequest) GetDatabaseNameOk() (*string, bool)`

GetDatabaseNameOk returns a tuple with the DatabaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabaseName

`func (o *QueryDBRequest) SetDatabaseName(v string)`

SetDatabaseName sets DatabaseName field to given value.


### GetQuery

`func (o *QueryDBRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *QueryDBRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *QueryDBRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRestoreAccountId

`func (o *QueryDBRequest) GetRestoreAccountId() string`

GetRestoreAccountId returns the RestoreAccountId field if non-nil, zero value otherwise.

### GetRestoreAccountIdOk

`func (o *QueryDBRequest) GetRestoreAccountIdOk() (*string, bool)`

GetRestoreAccountIdOk returns a tuple with the RestoreAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestoreAccountId

`func (o *QueryDBRequest) SetRestoreAccountId(v string)`

SetRestoreAccountId sets RestoreAccountId field to given value.


### GetDestination

`func (o *QueryDBRequest) GetDestination() QueryDBRestoreDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *QueryDBRequest) GetDestinationOk() (*QueryDBRestoreDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *QueryDBRequest) SetDestination(v QueryDBRestoreDestination)`

SetDestination sets Destination field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


