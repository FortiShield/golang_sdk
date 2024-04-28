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

// checks if the IngestersSecret type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersSecret{}

// IngestersSecret struct for IngestersSecret
type IngestersSecret struct {
	ImageLayerId *string `json:"ImageLayerId,omitempty"`
	Match *IngestersSecretMatch `json:"Match,omitempty"`
	Rule *IngestersSecretRule `json:"Rule,omitempty"`
	Severity *IngestersSecretSeverity `json:"Severity,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
}

// NewIngestersSecret instantiates a new IngestersSecret object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersSecret() *IngestersSecret {
	this := IngestersSecret{}
	return &this
}

// NewIngestersSecretWithDefaults instantiates a new IngestersSecret object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersSecretWithDefaults() *IngestersSecret {
	this := IngestersSecret{}
	return &this
}

// GetImageLayerId returns the ImageLayerId field value if set, zero value otherwise.
func (o *IngestersSecret) GetImageLayerId() string {
	if o == nil || IsNil(o.ImageLayerId) {
		var ret string
		return ret
	}
	return *o.ImageLayerId
}

// GetImageLayerIdOk returns a tuple with the ImageLayerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetImageLayerIdOk() (*string, bool) {
	if o == nil || IsNil(o.ImageLayerId) {
		return nil, false
	}
	return o.ImageLayerId, true
}

// HasImageLayerId returns a boolean if a field has been set.
func (o *IngestersSecret) HasImageLayerId() bool {
	if o != nil && !IsNil(o.ImageLayerId) {
		return true
	}

	return false
}

// SetImageLayerId gets a reference to the given string and assigns it to the ImageLayerId field.
func (o *IngestersSecret) SetImageLayerId(v string) {
	o.ImageLayerId = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *IngestersSecret) GetMatch() IngestersSecretMatch {
	if o == nil || IsNil(o.Match) {
		var ret IngestersSecretMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetMatchOk() (*IngestersSecretMatch, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *IngestersSecret) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given IngestersSecretMatch and assigns it to the Match field.
func (o *IngestersSecret) SetMatch(v IngestersSecretMatch) {
	o.Match = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *IngestersSecret) GetRule() IngestersSecretRule {
	if o == nil || IsNil(o.Rule) {
		var ret IngestersSecretRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetRuleOk() (*IngestersSecretRule, bool) {
	if o == nil || IsNil(o.Rule) {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *IngestersSecret) HasRule() bool {
	if o != nil && !IsNil(o.Rule) {
		return true
	}

	return false
}

// SetRule gets a reference to the given IngestersSecretRule and assigns it to the Rule field.
func (o *IngestersSecret) SetRule(v IngestersSecretRule) {
	o.Rule = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *IngestersSecret) GetSeverity() IngestersSecretSeverity {
	if o == nil || IsNil(o.Severity) {
		var ret IngestersSecretSeverity
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetSeverityOk() (*IngestersSecretSeverity, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// HasSeverity returns a boolean if a field has been set.
func (o *IngestersSecret) HasSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given IngestersSecretSeverity and assigns it to the Severity field.
func (o *IngestersSecret) SetSeverity(v IngestersSecretSeverity) {
	o.Severity = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *IngestersSecret) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersSecret) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *IngestersSecret) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *IngestersSecret) SetScanId(v string) {
	o.ScanId = &v
}

func (o IngestersSecret) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersSecret) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ImageLayerId) {
		toSerialize["ImageLayerId"] = o.ImageLayerId
	}
	if !IsNil(o.Match) {
		toSerialize["Match"] = o.Match
	}
	if !IsNil(o.Rule) {
		toSerialize["Rule"] = o.Rule
	}
	if !IsNil(o.Severity) {
		toSerialize["Severity"] = o.Severity
	}
	if !IsNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	return toSerialize, nil
}

type NullableIngestersSecret struct {
	value *IngestersSecret
	isSet bool
}

func (v NullableIngestersSecret) Get() *IngestersSecret {
	return v.value
}

func (v *NullableIngestersSecret) Set(val *IngestersSecret) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersSecret) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersSecret) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersSecret(val *IngestersSecret) *NullableIngestersSecret {
	return &NullableIngestersSecret{value: val, isSet: true}
}

func (v NullableIngestersSecret) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersSecret) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


