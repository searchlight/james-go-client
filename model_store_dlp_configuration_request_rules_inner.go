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

// checks if the StoreDLPConfigurationRequestRulesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreDLPConfigurationRequestRulesInner{}

// StoreDLPConfigurationRequestRulesInner struct for StoreDLPConfigurationRequestRulesInner
type StoreDLPConfigurationRequestRulesInner struct {
	// Description of the configuration item
	Explanation *string `json:"explanation,omitempty"`
	// Regular expression to match contents of targets
	Expression *string `json:"expression,omitempty"`
	// Unique identifier of the configuration item
	Id *string `json:"id,omitempty"`
	// Indicates if the expression should be applied to the subject headers and textual bodies of the mail
	TargetsContent *bool `json:"targetsContent,omitempty"`
	// Indicates if the expression should be applied to the recipients of the mail
	TargetsRecipients *bool `json:"targetsRecipients,omitempty"`
	// Indicates if the expression should be applied to the sender and from headers of the mail
	TargetsSender *bool `json:"targetsSender,omitempty"`
}

// NewStoreDLPConfigurationRequestRulesInner instantiates a new StoreDLPConfigurationRequestRulesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreDLPConfigurationRequestRulesInner() *StoreDLPConfigurationRequestRulesInner {
	this := StoreDLPConfigurationRequestRulesInner{}
	return &this
}

// NewStoreDLPConfigurationRequestRulesInnerWithDefaults instantiates a new StoreDLPConfigurationRequestRulesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDLPConfigurationRequestRulesInnerWithDefaults() *StoreDLPConfigurationRequestRulesInner {
	this := StoreDLPConfigurationRequestRulesInner{}
	return &this
}

// GetExplanation returns the Explanation field value if set, zero value otherwise.
func (o *StoreDLPConfigurationRequestRulesInner) GetExplanation() string {
	if o == nil || IsNil(o.Explanation) {
		var ret string
		return ret
	}
	return *o.Explanation
}

// GetExplanationOk returns a tuple with the Explanation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDLPConfigurationRequestRulesInner) GetExplanationOk() (*string, bool) {
	if o == nil || IsNil(o.Explanation) {
		return nil, false
	}
	return o.Explanation, true
}

// HasExplanation returns a boolean if a field has been set.
func (o *StoreDLPConfigurationRequestRulesInner) HasExplanation() bool {
	if o != nil && !IsNil(o.Explanation) {
		return true
	}

	return false
}

// SetExplanation gets a reference to the given string and assigns it to the Explanation field.
func (o *StoreDLPConfigurationRequestRulesInner) SetExplanation(v string) {
	o.Explanation = &v
}

// GetExpression returns the Expression field value if set, zero value otherwise.
func (o *StoreDLPConfigurationRequestRulesInner) GetExpression() string {
	if o == nil || IsNil(o.Expression) {
		var ret string
		return ret
	}
	return *o.Expression
}

// GetExpressionOk returns a tuple with the Expression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDLPConfigurationRequestRulesInner) GetExpressionOk() (*string, bool) {
	if o == nil || IsNil(o.Expression) {
		return nil, false
	}
	return o.Expression, true
}

// HasExpression returns a boolean if a field has been set.
func (o *StoreDLPConfigurationRequestRulesInner) HasExpression() bool {
	if o != nil && !IsNil(o.Expression) {
		return true
	}

	return false
}

// SetExpression gets a reference to the given string and assigns it to the Expression field.
func (o *StoreDLPConfigurationRequestRulesInner) SetExpression(v string) {
	o.Expression = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StoreDLPConfigurationRequestRulesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDLPConfigurationRequestRulesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StoreDLPConfigurationRequestRulesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StoreDLPConfigurationRequestRulesInner) SetId(v string) {
	o.Id = &v
}

// GetTargetsContent returns the TargetsContent field value if set, zero value otherwise.
func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsContent() bool {
	if o == nil || IsNil(o.TargetsContent) {
		var ret bool
		return ret
	}
	return *o.TargetsContent
}

// GetTargetsContentOk returns a tuple with the TargetsContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsContentOk() (*bool, bool) {
	if o == nil || IsNil(o.TargetsContent) {
		return nil, false
	}
	return o.TargetsContent, true
}

// HasTargetsContent returns a boolean if a field has been set.
func (o *StoreDLPConfigurationRequestRulesInner) HasTargetsContent() bool {
	if o != nil && !IsNil(o.TargetsContent) {
		return true
	}

	return false
}

// SetTargetsContent gets a reference to the given bool and assigns it to the TargetsContent field.
func (o *StoreDLPConfigurationRequestRulesInner) SetTargetsContent(v bool) {
	o.TargetsContent = &v
}

// GetTargetsRecipients returns the TargetsRecipients field value if set, zero value otherwise.
func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsRecipients() bool {
	if o == nil || IsNil(o.TargetsRecipients) {
		var ret bool
		return ret
	}
	return *o.TargetsRecipients
}

// GetTargetsRecipientsOk returns a tuple with the TargetsRecipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsRecipientsOk() (*bool, bool) {
	if o == nil || IsNil(o.TargetsRecipients) {
		return nil, false
	}
	return o.TargetsRecipients, true
}

// HasTargetsRecipients returns a boolean if a field has been set.
func (o *StoreDLPConfigurationRequestRulesInner) HasTargetsRecipients() bool {
	if o != nil && !IsNil(o.TargetsRecipients) {
		return true
	}

	return false
}

// SetTargetsRecipients gets a reference to the given bool and assigns it to the TargetsRecipients field.
func (o *StoreDLPConfigurationRequestRulesInner) SetTargetsRecipients(v bool) {
	o.TargetsRecipients = &v
}

// GetTargetsSender returns the TargetsSender field value if set, zero value otherwise.
func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsSender() bool {
	if o == nil || IsNil(o.TargetsSender) {
		var ret bool
		return ret
	}
	return *o.TargetsSender
}

// GetTargetsSenderOk returns a tuple with the TargetsSender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsSenderOk() (*bool, bool) {
	if o == nil || IsNil(o.TargetsSender) {
		return nil, false
	}
	return o.TargetsSender, true
}

// HasTargetsSender returns a boolean if a field has been set.
func (o *StoreDLPConfigurationRequestRulesInner) HasTargetsSender() bool {
	if o != nil && !IsNil(o.TargetsSender) {
		return true
	}

	return false
}

// SetTargetsSender gets a reference to the given bool and assigns it to the TargetsSender field.
func (o *StoreDLPConfigurationRequestRulesInner) SetTargetsSender(v bool) {
	o.TargetsSender = &v
}

func (o StoreDLPConfigurationRequestRulesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreDLPConfigurationRequestRulesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Explanation) {
		toSerialize["explanation"] = o.Explanation
	}
	if !IsNil(o.Expression) {
		toSerialize["expression"] = o.Expression
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TargetsContent) {
		toSerialize["targetsContent"] = o.TargetsContent
	}
	if !IsNil(o.TargetsRecipients) {
		toSerialize["targetsRecipients"] = o.TargetsRecipients
	}
	if !IsNil(o.TargetsSender) {
		toSerialize["targetsSender"] = o.TargetsSender
	}
	return toSerialize, nil
}

type NullableStoreDLPConfigurationRequestRulesInner struct {
	value *StoreDLPConfigurationRequestRulesInner
	isSet bool
}

func (v NullableStoreDLPConfigurationRequestRulesInner) Get() *StoreDLPConfigurationRequestRulesInner {
	return v.value
}

func (v *NullableStoreDLPConfigurationRequestRulesInner) Set(val *StoreDLPConfigurationRequestRulesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreDLPConfigurationRequestRulesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreDLPConfigurationRequestRulesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreDLPConfigurationRequestRulesInner(val *StoreDLPConfigurationRequestRulesInner) *NullableStoreDLPConfigurationRequestRulesInner {
	return &NullableStoreDLPConfigurationRequestRulesInner{value: val, isSet: true}
}

func (v NullableStoreDLPConfigurationRequestRulesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreDLPConfigurationRequestRulesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

