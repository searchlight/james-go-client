/*
Apache JAMES Web Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GetMailRepository200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMailRepository200Response{}

// GetMailRepository200Response struct for GetMailRepository200Response
type GetMailRepository200Response struct {
	Path *string `json:"path,omitempty"`
	Repository *string `json:"repository,omitempty"`
	Size *int32 `json:"size,omitempty"`
}

// NewGetMailRepository200Response instantiates a new GetMailRepository200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMailRepository200Response() *GetMailRepository200Response {
	this := GetMailRepository200Response{}
	return &this
}

// NewGetMailRepository200ResponseWithDefaults instantiates a new GetMailRepository200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMailRepository200ResponseWithDefaults() *GetMailRepository200Response {
	this := GetMailRepository200Response{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *GetMailRepository200Response) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMailRepository200Response) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *GetMailRepository200Response) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *GetMailRepository200Response) SetPath(v string) {
	o.Path = &v
}

// GetRepository returns the Repository field value if set, zero value otherwise.
func (o *GetMailRepository200Response) GetRepository() string {
	if o == nil || IsNil(o.Repository) {
		var ret string
		return ret
	}
	return *o.Repository
}

// GetRepositoryOk returns a tuple with the Repository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMailRepository200Response) GetRepositoryOk() (*string, bool) {
	if o == nil || IsNil(o.Repository) {
		return nil, false
	}
	return o.Repository, true
}

// HasRepository returns a boolean if a field has been set.
func (o *GetMailRepository200Response) HasRepository() bool {
	if o != nil && !IsNil(o.Repository) {
		return true
	}

	return false
}

// SetRepository gets a reference to the given string and assigns it to the Repository field.
func (o *GetMailRepository200Response) SetRepository(v string) {
	o.Repository = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *GetMailRepository200Response) GetSize() int32 {
	if o == nil || IsNil(o.Size) {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetMailRepository200Response) GetSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *GetMailRepository200Response) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *GetMailRepository200Response) SetSize(v int32) {
	o.Size = &v
}

func (o GetMailRepository200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMailRepository200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Repository) {
		toSerialize["repository"] = o.Repository
	}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	return toSerialize, nil
}

type NullableGetMailRepository200Response struct {
	value *GetMailRepository200Response
	isSet bool
}

func (v NullableGetMailRepository200Response) Get() *GetMailRepository200Response {
	return v.value
}

func (v *NullableGetMailRepository200Response) Set(val *GetMailRepository200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMailRepository200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMailRepository200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMailRepository200Response(val *GetMailRepository200Response) *NullableGetMailRepository200Response {
	return &NullableGetMailRepository200Response{value: val, isSet: true}
}

func (v NullableGetMailRepository200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMailRepository200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


