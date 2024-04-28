/*
Kengine ThreatMapper

Kengine Runtime API provides programmatic control over Kengine microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.2.0
Contact: community@kengine.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the IngestersSecretRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersSecretRule{}

// IngestersSecretRule struct for IngestersSecretRule
type IngestersSecretRule struct {
	Id *int32 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Part *string `json:"part,omitempty"`
	SignatureToMatch *string `json:"signature_to_match,omitempty"`
}

// NewIngestersSecretRule instantiates a new IngestersSecretRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersSecretRule() *IngestersSecretRule {
	this := IngestersSecretRule{}
	return &this
}

// NewIngestersSecretRuleWithDefaults instantiates a new IngestersSecretRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersSecretRuleWithDefaults() *IngestersSecretRule {
	this := IngestersSecretRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IngestersSecretRule) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretRule) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IngestersSecretRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *IngestersSecretRule) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IngestersSecretRule) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretRule) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IngestersSecretRule) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IngestersSecretRule) SetName(v string) {
	o.Name = &v
}

// GetPart returns the Part field value if set, zero value otherwise.
func (o *IngestersSecretRule) GetPart() string {
	if o == nil || IsNil(o.Part) {
		var ret string
		return ret
	}
	return *o.Part
}

// GetPartOk returns a tuple with the Part field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretRule) GetPartOk() (*string, bool) {
	if o == nil || IsNil(o.Part) {
		return nil, false
	}
	return o.Part, true
}

// HasPart returns a boolean if a field has been set.
func (o *IngestersSecretRule) HasPart() bool {
	if o != nil && !IsNil(o.Part) {
		return true
	}

	return false
}

// SetPart gets a reference to the given string and assigns it to the Part field.
func (o *IngestersSecretRule) SetPart(v string) {
	o.Part = &v
}

// GetSignatureToMatch returns the SignatureToMatch field value if set, zero value otherwise.
func (o *IngestersSecretRule) GetSignatureToMatch() string {
	if o == nil || IsNil(o.SignatureToMatch) {
		var ret string
		return ret
	}
	return *o.SignatureToMatch
}

// GetSignatureToMatchOk returns a tuple with the SignatureToMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecretRule) GetSignatureToMatchOk() (*string, bool) {
	if o == nil || IsNil(o.SignatureToMatch) {
		return nil, false
	}
	return o.SignatureToMatch, true
}

// HasSignatureToMatch returns a boolean if a field has been set.
func (o *IngestersSecretRule) HasSignatureToMatch() bool {
	if o != nil && !IsNil(o.SignatureToMatch) {
		return true
	}

	return false
}

// SetSignatureToMatch gets a reference to the given string and assigns it to the SignatureToMatch field.
func (o *IngestersSecretRule) SetSignatureToMatch(v string) {
	o.SignatureToMatch = &v
}

func (o IngestersSecretRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersSecretRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Part) {
		toSerialize["part"] = o.Part
	}
	if !IsNil(o.SignatureToMatch) {
		toSerialize["signature_to_match"] = o.SignatureToMatch
	}
	return toSerialize, nil
}

type NullableIngestersSecretRule struct {
	value *IngestersSecretRule
	isSet bool
}

func (v NullableIngestersSecretRule) Get() *IngestersSecretRule {
	return v.value
}

func (v *NullableIngestersSecretRule) Set(val *IngestersSecretRule) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersSecretRule) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersSecretRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersSecretRule(val *IngestersSecretRule) *NullableIngestersSecretRule {
	return &NullableIngestersSecretRule{value: val, isSet: true}
}

func (v NullableIngestersSecretRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersSecretRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


