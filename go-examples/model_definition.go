/*
NCI EVS Rest API

Endpoints to support searching, metadata, and content retrieval for EVS terminologies. To learn more about how to interact with this api, see the <a href=\"https://github.com/NCIEVS/evsrestapi-client-SDK\">Github evsrestapi-client-SDK project.</a>

API version: 1.7.2.RELEASE
Contact: NCIAppSupport@nih.gov
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package evs

import (
	"encoding/json"
)

// checks if the Definition type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Definition{}

// Definition Represents a text definition for a concept
type Definition struct {
	// URI for this element in an rdf-based source file
	Uri *string `json:"uri,omitempty"`
	// Used to indicate the total amount of data in cases where a limit is being applied
	Ct *int32 `json:"ct,omitempty"`
	// Text definition value
	Definition *string `json:"definition,omitempty"`
	// Used by search calls to provide information for highlighting a view of results
	Highlight *string `json:"highlight,omitempty"`
	// Definition type
	Type *string `json:"type,omitempty"`
	// Definition source
	Source *string `json:"source,omitempty"`
	// Type/value qualifiers on the definition
	Qualifiers []Qualifier `json:"qualifiers,omitempty"`
}

// NewDefinition instantiates a new Definition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefinition() *Definition {
	this := Definition{}
	return &this
}

// NewDefinitionWithDefaults instantiates a new Definition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefinitionWithDefaults() *Definition {
	this := Definition{}
	return &this
}

// GetUri returns the Uri field value if set, zero value otherwise.
func (o *Definition) GetUri() string {
	if o == nil || IsNil(o.Uri) {
		var ret string
		return ret
	}
	return *o.Uri
}

// GetUriOk returns a tuple with the Uri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Definition) GetUriOk() (*string, bool) {
	if o == nil || IsNil(o.Uri) {
		return nil, false
	}
	return o.Uri, true
}

// HasUri returns a boolean if a field has been set.
func (o *Definition) HasUri() bool {
	if o != nil && !IsNil(o.Uri) {
		return true
	}

	return false
}

// SetUri gets a reference to the given string and assigns it to the Uri field.
func (o *Definition) SetUri(v string) {
	o.Uri = &v
}

// GetCt returns the Ct field value if set, zero value otherwise.
func (o *Definition) GetCt() int32 {
	if o == nil || IsNil(o.Ct) {
		var ret int32
		return ret
	}
	return *o.Ct
}

// GetCtOk returns a tuple with the Ct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Definition) GetCtOk() (*int32, bool) {
	if o == nil || IsNil(o.Ct) {
		return nil, false
	}
	return o.Ct, true
}

// HasCt returns a boolean if a field has been set.
func (o *Definition) HasCt() bool {
	if o != nil && !IsNil(o.Ct) {
		return true
	}

	return false
}

// SetCt gets a reference to the given int32 and assigns it to the Ct field.
func (o *Definition) SetCt(v int32) {
	o.Ct = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *Definition) GetDefinition() string {
	if o == nil || IsNil(o.Definition) {
		var ret string
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Definition) GetDefinitionOk() (*string, bool) {
	if o == nil || IsNil(o.Definition) {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *Definition) HasDefinition() bool {
	if o != nil && !IsNil(o.Definition) {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given string and assigns it to the Definition field.
func (o *Definition) SetDefinition(v string) {
	o.Definition = &v
}

// GetHighlight returns the Highlight field value if set, zero value otherwise.
func (o *Definition) GetHighlight() string {
	if o == nil || IsNil(o.Highlight) {
		var ret string
		return ret
	}
	return *o.Highlight
}

// GetHighlightOk returns a tuple with the Highlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Definition) GetHighlightOk() (*string, bool) {
	if o == nil || IsNil(o.Highlight) {
		return nil, false
	}
	return o.Highlight, true
}

// HasHighlight returns a boolean if a field has been set.
func (o *Definition) HasHighlight() bool {
	if o != nil && !IsNil(o.Highlight) {
		return true
	}

	return false
}

// SetHighlight gets a reference to the given string and assigns it to the Highlight field.
func (o *Definition) SetHighlight(v string) {
	o.Highlight = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Definition) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Definition) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Definition) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Definition) SetType(v string) {
	o.Type = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *Definition) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Definition) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *Definition) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *Definition) SetSource(v string) {
	o.Source = &v
}

// GetQualifiers returns the Qualifiers field value if set, zero value otherwise.
func (o *Definition) GetQualifiers() []Qualifier {
	if o == nil || IsNil(o.Qualifiers) {
		var ret []Qualifier
		return ret
	}
	return o.Qualifiers
}

// GetQualifiersOk returns a tuple with the Qualifiers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Definition) GetQualifiersOk() ([]Qualifier, bool) {
	if o == nil || IsNil(o.Qualifiers) {
		return nil, false
	}
	return o.Qualifiers, true
}

// HasQualifiers returns a boolean if a field has been set.
func (o *Definition) HasQualifiers() bool {
	if o != nil && !IsNil(o.Qualifiers) {
		return true
	}

	return false
}

// SetQualifiers gets a reference to the given []Qualifier and assigns it to the Qualifiers field.
func (o *Definition) SetQualifiers(v []Qualifier) {
	o.Qualifiers = v
}

func (o Definition) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Definition) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uri) {
		toSerialize["uri"] = o.Uri
	}
	if !IsNil(o.Ct) {
		toSerialize["ct"] = o.Ct
	}
	if !IsNil(o.Definition) {
		toSerialize["definition"] = o.Definition
	}
	if !IsNil(o.Highlight) {
		toSerialize["highlight"] = o.Highlight
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Qualifiers) {
		toSerialize["qualifiers"] = o.Qualifiers
	}
	return toSerialize, nil
}

type NullableDefinition struct {
	value *Definition
	isSet bool
}

func (v NullableDefinition) Get() *Definition {
	return v.value
}

func (v *NullableDefinition) Set(val *Definition) {
	v.value = val
	v.isSet = true
}

func (v NullableDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefinition(val *Definition) *NullableDefinition {
	return &NullableDefinition{value: val, isSet: true}
}

func (v NullableDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


