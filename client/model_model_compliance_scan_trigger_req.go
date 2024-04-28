/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.2.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ModelComplianceScanTriggerReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelComplianceScanTriggerReq{}

// ModelComplianceScanTriggerReq struct for ModelComplianceScanTriggerReq
type ModelComplianceScanTriggerReq struct {
	BenchmarkTypes []string `json:"benchmark_types"`
	Filters ModelScanFilter `json:"filters"`
	IsPriority *bool `json:"is_priority,omitempty"`
	NodeIds []ModelNodeIdentifier `json:"node_ids"`
}

type _ModelComplianceScanTriggerReq ModelComplianceScanTriggerReq

// NewModelComplianceScanTriggerReq instantiates a new ModelComplianceScanTriggerReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelComplianceScanTriggerReq(benchmarkTypes []string, filters ModelScanFilter, nodeIds []ModelNodeIdentifier) *ModelComplianceScanTriggerReq {
	this := ModelComplianceScanTriggerReq{}
	this.BenchmarkTypes = benchmarkTypes
	this.Filters = filters
	this.NodeIds = nodeIds
	return &this
}

// NewModelComplianceScanTriggerReqWithDefaults instantiates a new ModelComplianceScanTriggerReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelComplianceScanTriggerReqWithDefaults() *ModelComplianceScanTriggerReq {
	this := ModelComplianceScanTriggerReq{}
	return &this
}

// GetBenchmarkTypes returns the BenchmarkTypes field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelComplianceScanTriggerReq) GetBenchmarkTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.BenchmarkTypes
}

// GetBenchmarkTypesOk returns a tuple with the BenchmarkTypes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelComplianceScanTriggerReq) GetBenchmarkTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.BenchmarkTypes) {
		return nil, false
	}
	return o.BenchmarkTypes, true
}

// SetBenchmarkTypes sets field value
func (o *ModelComplianceScanTriggerReq) SetBenchmarkTypes(v []string) {
	o.BenchmarkTypes = v
}

// GetFilters returns the Filters field value
func (o *ModelComplianceScanTriggerReq) GetFilters() ModelScanFilter {
	if o == nil {
		var ret ModelScanFilter
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanTriggerReq) GetFiltersOk() (*ModelScanFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *ModelComplianceScanTriggerReq) SetFilters(v ModelScanFilter) {
	o.Filters = v
}

// GetIsPriority returns the IsPriority field value if set, zero value otherwise.
func (o *ModelComplianceScanTriggerReq) GetIsPriority() bool {
	if o == nil || IsNil(o.IsPriority) {
		var ret bool
		return ret
	}
	return *o.IsPriority
}

// GetIsPriorityOk returns a tuple with the IsPriority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelComplianceScanTriggerReq) GetIsPriorityOk() (*bool, bool) {
	if o == nil || IsNil(o.IsPriority) {
		return nil, false
	}
	return o.IsPriority, true
}

// HasIsPriority returns a boolean if a field has been set.
func (o *ModelComplianceScanTriggerReq) HasIsPriority() bool {
	if o != nil && !IsNil(o.IsPriority) {
		return true
	}

	return false
}

// SetIsPriority gets a reference to the given bool and assigns it to the IsPriority field.
func (o *ModelComplianceScanTriggerReq) SetIsPriority(v bool) {
	o.IsPriority = &v
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []ModelNodeIdentifier will be returned
func (o *ModelComplianceScanTriggerReq) GetNodeIds() []ModelNodeIdentifier {
	if o == nil {
		var ret []ModelNodeIdentifier
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelComplianceScanTriggerReq) GetNodeIdsOk() ([]ModelNodeIdentifier, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *ModelComplianceScanTriggerReq) SetNodeIds(v []ModelNodeIdentifier) {
	o.NodeIds = v
}

func (o ModelComplianceScanTriggerReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelComplianceScanTriggerReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.BenchmarkTypes != nil {
		toSerialize["benchmark_types"] = o.BenchmarkTypes
	}
	toSerialize["filters"] = o.Filters
	if !IsNil(o.IsPriority) {
		toSerialize["is_priority"] = o.IsPriority
	}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	return toSerialize, nil
}

func (o *ModelComplianceScanTriggerReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"benchmark_types",
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

	varModelComplianceScanTriggerReq := _ModelComplianceScanTriggerReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelComplianceScanTriggerReq)

	if err != nil {
		return err
	}

	*o = ModelComplianceScanTriggerReq(varModelComplianceScanTriggerReq)

	return err
}

type NullableModelComplianceScanTriggerReq struct {
	value *ModelComplianceScanTriggerReq
	isSet bool
}

func (v NullableModelComplianceScanTriggerReq) Get() *ModelComplianceScanTriggerReq {
	return v.value
}

func (v *NullableModelComplianceScanTriggerReq) Set(val *ModelComplianceScanTriggerReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelComplianceScanTriggerReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelComplianceScanTriggerReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelComplianceScanTriggerReq(val *ModelComplianceScanTriggerReq) *NullableModelComplianceScanTriggerReq {
	return &NullableModelComplianceScanTriggerReq{value: val, isSet: true}
}

func (v NullableModelComplianceScanTriggerReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelComplianceScanTriggerReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


