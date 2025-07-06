# FilePath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | Absolute path to the file or directory to restore. | 
**IsDirectory** | **bool** | Whether &#x60;path&#x60; is a directory. If &#x60;true&#x60;, Eon restores all files in all subdirectories under the path. If &#x60;false&#x60;, Eon restores only the file at the path.  | 

## Methods

### NewFilePath

`func NewFilePath(path string, isDirectory bool, ) *FilePath`

NewFilePath instantiates a new FilePath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilePathWithDefaults

`func NewFilePathWithDefaults() *FilePath`

NewFilePathWithDefaults instantiates a new FilePath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *FilePath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *FilePath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *FilePath) SetPath(v string)`

SetPath sets Path field to given value.


### GetIsDirectory

`func (o *FilePath) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *FilePath) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *FilePath) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


