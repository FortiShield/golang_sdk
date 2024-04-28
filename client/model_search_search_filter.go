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

// checks if the SearchSearchFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSearchFilter{}

// SearchSearchFilter struct for SearchSearchFilter
type SearchSearchFilter struct {
	Filters ReportersFieldsFilters `json:"filters"`
	InFieldFilter []string `json:"in_field_filter"`
	Window ModelFetchWindow `json:"window"`
}

type _SearchSearchFilter SearchSearchFilter

// NewSearchSearchFilter instantiates a new SearchSearchFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSearchFilter(filters ReportersFieldsFilters, inFieldFilter []string, window ModelFetchWindow) *SearchSearchFilter {
	this := SearchSearchFilter{}
	this.Filters = filters
	this.InFieldFilter = inFieldFilter
	this.Window = window
	return &this
}

// NewSearchSearchFilterWithDefaults instantiates a new SearchSearchFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSearchFilterWithDefaults() *SearchSearchFilter {
	this := SearchSearchFilter{}
	return &this
}

// GetFilters returns the Filters field value
func (o *SearchSearchFilter) GetFilters() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value
// and a boolean to check if the value has been set.
func (o *SearchSearchFilter) GetFiltersOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filters, true
}

// SetFilters sets field value
func (o *SearchSearchFilter) SetFilters(v ReportersFieldsFilters) {
	o.Filters = v
}

// GetInFieldFilter returns the InFieldFilter field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *SearchSearchFilter) GetInFieldFilter() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.InFieldFilter
}

// GetInFieldFilterOk returns a tuple with the InFieldFilter field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SearchSearchFilter) GetInFieldFilterOk() ([]string, bool) {
	if o == nil || IsNil(o.InFieldFilter) {
		return nil, false
	}
	return o.InFieldFilter, true
}

// SetInFieldFilter sets field value
func (o *SearchSearchFilter) SetInFieldFilter(v []string) {
	o.InFieldFilter = v
}

// GetWindow returns the Window field value
func (o *SearchSearchFilter) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *SearchSearchFilter) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *SearchSearchFilter) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o SearchSearchFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSearchFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["filters"] = o.Filters
	if o.InFieldFilter != nil {
		toSerialize["in_field_filter"] = o.InFieldFilter
	}
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *SearchSearchFilter) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"filters",
		"in_field_filter",
		"window",
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

	varSearchSearchFilter := _SearchSearchFilter{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSearchSearchFilter)

	if err != nil {
		return err
	}

	*o = SearchSearchFilter(varSearchSearchFilter)

	return err
}

type NullableSearchSearchFilter struct {
	value *SearchSearchFilter
	isSet bool
}

func (v NullableSearchSearchFilter) Get() *SearchSearchFilter {
	return v.value
}

func (v *NullableSearchSearchFilter) Set(val *SearchSearchFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSearchFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSearchFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSearchFilter(val *SearchSearchFilter) *NullableSearchSearchFilter {
	return &NullableSearchSearchFilter{value: val, isSet: true}
}

func (v NullableSearchSearchFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSearchFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

