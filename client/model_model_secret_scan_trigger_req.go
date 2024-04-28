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
	"bytes"
	"fmt"
)

// checks if the ModelSecretScanTriggerReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelSecretScanTriggerReq{}

// ModelSecretScanTriggerReq struct for ModelSecretScanTriggerReq
type ModelSecretScanTriggerReq struct {
	Filters ModelScanFilter `json:"filters"`
	IsPriority *bool `json:"is_priority,omitempty"`
	NodeIds []ModelNodeIdentifier `json:"node_ids"`
}

type _ModelSecretScanTriggerReq ModelSecretScanTriggerReq

// NewModelSecretScanTriggerReq instantiates a new ModelSecretScanTriggerReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelSecretScanTriggerReq(filters ModelScanFilter, nodeIds []ModelNodeIdentifier) *ModelSecretScanTriggerReq {
	this := ModelSecretScanTriggerReq{}
	this.Filters = filters
	this.NodeIds = nodeIds
	return &this
}

// NewModelSecretScanTriggerReqWithDefaults instantiates a new ModelSecretScanTriggerReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelSecretScanTriggerReqWithDefaults() *ModelSecretScanTriggerReq {
	this := ModelSecretScanTriggerReq{}
	return &this
}

// GetFilters returns the Filters field value
func (o *ModelSecretScanTriggerReq) GetFilters() ModelScanFilter {
	if o == nil {
		var ret ModelScanFilter
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ModelSecretScanTriggerReq) GetFiltersOk() (*ModelScanFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *ModelSecretScanTriggerReq) SetFilters(v ModelScanFilter) {
	o.Filters = v
}

// GetIsPriority returns the IsPriority field value if set, zero value otherwise.
func (o *ModelSecretScanTriggerReq) GetIsPriority() bool {
	if o == nil || IsNil(o.IsPriority) {
		var ret bool
		return ret
	}
	return *o.IsPriority
}

// GetIsPriorityOk returns a tuple with the IsPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelSecretScanTriggerReq) GetIsPriorityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPriority) {
		return nil, false
	}
	return o.IsPriority, true
}

// HasIsPriority returns a boolean if a field has been set.
func (o *ModelSecretScanTriggerReq) HasIsPriority() bool {
	if o != nil && !IsNil(o.IsPriority) {
		return true
	}

	return false
}

// SetIsPriority gets a reference to the given bool and assigns it to the IsPriority field.
func (o *ModelSecretScanTriggerReq) SetIsPriority(v bool) {
	o.IsPriority = &v
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []ModelNodeIdentifier will be returned
func (o *ModelSecretScanTriggerReq) GetNodeIds() []ModelNodeIdentifier {
	if o == nil {
		var ret []ModelNodeIdentifier
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelSecretScanTriggerReq) GetNodeIdsOk() ([]ModelNodeIdentifier, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *ModelSecretScanTriggerReq) SetNodeIds(v []ModelNodeIdentifier) {
	o.NodeIds = v
}

func (o ModelSecretScanTriggerReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelSecretScanTriggerReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filters"] = o.Filters
	if !IsNil(o.IsPriority) {
		toSerialize["is_priority"] = o.IsPriority
	}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	return toSerialize, nil
}

func (o *ModelSecretScanTriggerReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filters",
		"node_ids",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varModelSecretScanTriggerReq := _ModelSecretScanTriggerReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelSecretScanTriggerReq)

	if err != nil {
		return err
	}

	*o = ModelSecretScanTriggerReq(varModelSecretScanTriggerReq)

	return err
}

type NullableModelSecretScanTriggerReq struct {
	value *ModelSecretScanTriggerReq
	isSet bool
}

func (v NullableModelSecretScanTriggerReq) Get() *ModelSecretScanTriggerReq {
	return v.value
}

func (v *NullableModelSecretScanTriggerReq) Set(val *ModelSecretScanTriggerReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelSecretScanTriggerReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelSecretScanTriggerReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelSecretScanTriggerReq(val *ModelSecretScanTriggerReq) *NullableModelSecretScanTriggerReq {
	return &NullableModelSecretScanTriggerReq{value: val, isSet: true}
}

func (v NullableModelSecretScanTriggerReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelSecretScanTriggerReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

