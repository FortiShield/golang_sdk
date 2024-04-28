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

// checks if the ModelAgentUpgrade type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAgentUpgrade{}

// ModelAgentUpgrade struct for ModelAgentUpgrade
type ModelAgentUpgrade struct {
	NodeIds []string `json:"node_ids"`
	Version string `json:"version"`
}

type _ModelAgentUpgrade ModelAgentUpgrade

// NewModelAgentUpgrade instantiates a new ModelAgentUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAgentUpgrade(nodeIds []string, version string) *ModelAgentUpgrade {
	this := ModelAgentUpgrade{}
	this.NodeIds = nodeIds
	this.Version = version
	return &this
}

// NewModelAgentUpgradeWithDefaults instantiates a new ModelAgentUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAgentUpgradeWithDefaults() *ModelAgentUpgrade {
	this := ModelAgentUpgrade{}
	return &this
}

// GetNodeIds returns the NodeIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelAgentUpgrade) GetNodeIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NodeIds
}

// GetNodeIdsOk returns a tuple with the NodeIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelAgentUpgrade) GetNodeIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.NodeIds) {
		return nil, false
	}
	return o.NodeIds, true
}

// SetNodeIds sets field value
func (o *ModelAgentUpgrade) SetNodeIds(v []string) {
	o.NodeIds = v
}

// GetVersion returns the Version field value
func (o *ModelAgentUpgrade) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ModelAgentUpgrade) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ModelAgentUpgrade) SetVersion(v string) {
	o.Version = v
}

func (o ModelAgentUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAgentUpgrade) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeIds != nil {
		toSerialize["node_ids"] = o.NodeIds
	}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

func (o *ModelAgentUpgrade) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_ids",
		"version",
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

	varModelAgentUpgrade := _ModelAgentUpgrade{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelAgentUpgrade)

	if err != nil {
		return err
	}

	*o = ModelAgentUpgrade(varModelAgentUpgrade)

	return err
}

type NullableModelAgentUpgrade struct {
	value *ModelAgentUpgrade
	isSet bool
}

func (v NullableModelAgentUpgrade) Get() *ModelAgentUpgrade {
	return v.value
}

func (v *NullableModelAgentUpgrade) Set(val *ModelAgentUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAgentUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAgentUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAgentUpgrade(val *ModelAgentUpgrade) *NullableModelAgentUpgrade {
	return &NullableModelAgentUpgrade{value: val, isSet: true}
}

func (v NullableModelAgentUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAgentUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


