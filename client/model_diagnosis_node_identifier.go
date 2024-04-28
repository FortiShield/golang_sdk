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

// checks if the DiagnosisNodeIdentifier type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiagnosisNodeIdentifier{}

// DiagnosisNodeIdentifier struct for DiagnosisNodeIdentifier
type DiagnosisNodeIdentifier struct {
	NodeId string `json:"node_id"`
	NodeType string `json:"node_type"`
}

type _DiagnosisNodeIdentifier DiagnosisNodeIdentifier

// NewDiagnosisNodeIdentifier instantiates a new DiagnosisNodeIdentifier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiagnosisNodeIdentifier(nodeId string, nodeType string) *DiagnosisNodeIdentifier {
	this := DiagnosisNodeIdentifier{}
	this.NodeId = nodeId
	this.NodeType = nodeType
	return &this
}

// NewDiagnosisNodeIdentifierWithDefaults instantiates a new DiagnosisNodeIdentifier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiagnosisNodeIdentifierWithDefaults() *DiagnosisNodeIdentifier {
	this := DiagnosisNodeIdentifier{}
	return &this
}

// GetNodeId returns the NodeId field value
func (o *DiagnosisNodeIdentifier) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *DiagnosisNodeIdentifier) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *DiagnosisNodeIdentifier) SetNodeId(v string) {
	o.NodeId = v
}

// GetNodeType returns the NodeType field value
func (o *DiagnosisNodeIdentifier) GetNodeType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeType
}

// GetNodeTypeOk returns a tuple with the NodeType field value
// and a boolean to check if the value has been set.
func (o *DiagnosisNodeIdentifier) GetNodeTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeType, true
}

// SetNodeType sets field value
func (o *DiagnosisNodeIdentifier) SetNodeType(v string) {
	o.NodeType = v
}

func (o DiagnosisNodeIdentifier) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiagnosisNodeIdentifier) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_id"] = o.NodeId
	toSerialize["node_type"] = o.NodeType
	return toSerialize, nil
}

func (o *DiagnosisNodeIdentifier) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"node_id",
		"node_type",
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

	varDiagnosisNodeIdentifier := _DiagnosisNodeIdentifier{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDiagnosisNodeIdentifier)

	if err != nil {
		return err
	}

	*o = DiagnosisNodeIdentifier(varDiagnosisNodeIdentifier)

	return err
}

type NullableDiagnosisNodeIdentifier struct {
	value *DiagnosisNodeIdentifier
	isSet bool
}

func (v NullableDiagnosisNodeIdentifier) Get() *DiagnosisNodeIdentifier {
	return v.value
}

func (v *NullableDiagnosisNodeIdentifier) Set(val *DiagnosisNodeIdentifier) {
	v.value = val
	v.isSet = true
}

func (v NullableDiagnosisNodeIdentifier) IsSet() bool {
	return v.isSet
}

func (v *NullableDiagnosisNodeIdentifier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiagnosisNodeIdentifier(val *DiagnosisNodeIdentifier) *NullableDiagnosisNodeIdentifier {
	return &NullableDiagnosisNodeIdentifier{value: val, isSet: true}
}

func (v NullableDiagnosisNodeIdentifier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiagnosisNodeIdentifier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


