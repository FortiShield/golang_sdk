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

// checks if the ReportersCompareFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReportersCompareFilter{}

// ReportersCompareFilter struct for ReportersCompareFilter
type ReportersCompareFilter struct {
	FieldName string `json:"field_name"`
	FieldValue interface{} `json:"field_value"`
	GreaterThan bool `json:"greater_than"`
}

type _ReportersCompareFilter ReportersCompareFilter

// NewReportersCompareFilter instantiates a new ReportersCompareFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportersCompareFilter(fieldName string, fieldValue interface{}, greaterThan bool) *ReportersCompareFilter {
	this := ReportersCompareFilter{}
	this.FieldName = fieldName
	this.FieldValue = fieldValue
	this.GreaterThan = greaterThan
	return &this
}

// NewReportersCompareFilterWithDefaults instantiates a new ReportersCompareFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportersCompareFilterWithDefaults() *ReportersCompareFilter {
	this := ReportersCompareFilter{}
	return &this
}

// GetFieldName returns the FieldName field value
func (o *ReportersCompareFilter) GetFieldName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FieldName
}

// GetFieldNameOk returns a tuple with the FieldName field value
// and a boolean to check if the value has been set.
func (o *ReportersCompareFilter) GetFieldNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FieldName, true
}

// SetFieldName sets field value
func (o *ReportersCompareFilter) SetFieldName(v string) {
	o.FieldName = v
}

// GetFieldValue returns the FieldValue field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ReportersCompareFilter) GetFieldValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ReportersCompareFilter) GetFieldValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.FieldValue) {
		return nil, false
	}
	return &o.FieldValue, true
}

// SetFieldValue sets field value
func (o *ReportersCompareFilter) SetFieldValue(v interface{}) {
	o.FieldValue = v
}

// GetGreaterThan returns the GreaterThan field value
func (o *ReportersCompareFilter) GetGreaterThan() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.GreaterThan
}

// GetGreaterThanOk returns a tuple with the GreaterThan field value
// and a boolean to check if the value has been set.
func (o *ReportersCompareFilter) GetGreaterThanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GreaterThan, true
}

// SetGreaterThan sets field value
func (o *ReportersCompareFilter) SetGreaterThan(v bool) {
	o.GreaterThan = v
}

func (o ReportersCompareFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReportersCompareFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["field_name"] = o.FieldName
	if o.FieldValue != nil {
		toSerialize["field_value"] = o.FieldValue
	}
	toSerialize["greater_than"] = o.GreaterThan
	return toSerialize, nil
}

func (o *ReportersCompareFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"field_name",
		"field_value",
		"greater_than",
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

	varReportersCompareFilter := _ReportersCompareFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReportersCompareFilter)

	if err != nil {
		return err
	}

	*o = ReportersCompareFilter(varReportersCompareFilter)

	return err
}

type NullableReportersCompareFilter struct {
	value *ReportersCompareFilter
	isSet bool
}

func (v NullableReportersCompareFilter) Get() *ReportersCompareFilter {
	return v.value
}

func (v *NullableReportersCompareFilter) Set(val *ReportersCompareFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableReportersCompareFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableReportersCompareFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportersCompareFilter(val *ReportersCompareFilter) *NullableReportersCompareFilter {
	return &NullableReportersCompareFilter{value: val, isSet: true}
}

func (v NullableReportersCompareFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportersCompareFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


