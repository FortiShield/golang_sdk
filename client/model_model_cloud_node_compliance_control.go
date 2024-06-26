/*
Khulnasoft Kengine

Khulnasoft Runtime API provides programmatic control over Khulnasoft microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@khulnasoft.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the ModelCloudNodeComplianceControl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCloudNodeComplianceControl{}

// ModelCloudNodeComplianceControl struct for ModelCloudNodeComplianceControl
type ModelCloudNodeComplianceControl struct {
	CategoryHierarchy []string `json:"category_hierarchy,omitempty"`
	ControlId *string `json:"control_id,omitempty"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Service *string `json:"service,omitempty"`
	Title *string `json:"title,omitempty"`
}

// NewModelCloudNodeComplianceControl instantiates a new ModelCloudNodeComplianceControl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCloudNodeComplianceControl() *ModelCloudNodeComplianceControl {
	this := ModelCloudNodeComplianceControl{}
	return &this
}

// NewModelCloudNodeComplianceControlWithDefaults instantiates a new ModelCloudNodeComplianceControl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCloudNodeComplianceControlWithDefaults() *ModelCloudNodeComplianceControl {
	this := ModelCloudNodeComplianceControl{}
	return &this
}

// GetCategoryHierarchy returns the CategoryHierarchy field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ModelCloudNodeComplianceControl) GetCategoryHierarchy() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.CategoryHierarchy
}

// GetCategoryHierarchyOk returns a tuple with the CategoryHierarchy field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelCloudNodeComplianceControl) GetCategoryHierarchyOk() ([]string, bool) {
	if o == nil || IsNil(o.CategoryHierarchy) {
		return nil, false
	}
	return o.CategoryHierarchy, true
}

// HasCategoryHierarchy returns a boolean if a field has been set.
func (o *ModelCloudNodeComplianceControl) HasCategoryHierarchy() bool {
	if o != nil && !IsNil(o.CategoryHierarchy) {
		return true
	}

	return false
}

// SetCategoryHierarchy gets a reference to the given []string and assigns it to the CategoryHierarchy field.
func (o *ModelCloudNodeComplianceControl) SetCategoryHierarchy(v []string) {
	o.CategoryHierarchy = v
}

// GetControlId returns the ControlId field value if set, zero value otherwise.
func (o *ModelCloudNodeComplianceControl) GetControlId() string {
	if o == nil || IsNil(o.ControlId) {
		var ret string
		return ret
	}
	return *o.ControlId
}

// GetControlIdOk returns a tuple with the ControlId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeComplianceControl) GetControlIdOk() (*string, bool) {
	if o == nil || IsNil(o.ControlId) {
		return nil, false
	}
	return o.ControlId, true
}

// HasControlId returns a boolean if a field has been set.
func (o *ModelCloudNodeComplianceControl) HasControlId() bool {
	if o != nil && !IsNil(o.ControlId) {
		return true
	}

	return false
}

// SetControlId gets a reference to the given string and assigns it to the ControlId field.
func (o *ModelCloudNodeComplianceControl) SetControlId(v string) {
	o.ControlId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ModelCloudNodeComplianceControl) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeComplianceControl) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ModelCloudNodeComplianceControl) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ModelCloudNodeComplianceControl) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ModelCloudNodeComplianceControl) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeComplianceControl) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ModelCloudNodeComplianceControl) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ModelCloudNodeComplianceControl) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *ModelCloudNodeComplianceControl) GetService() string {
	if o == nil || IsNil(o.Service) {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeComplianceControl) GetServiceOk() (*string, bool) {
	if o == nil || IsNil(o.Service) {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *ModelCloudNodeComplianceControl) HasService() bool {
	if o != nil && !IsNil(o.Service) {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *ModelCloudNodeComplianceControl) SetService(v string) {
	o.Service = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ModelCloudNodeComplianceControl) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCloudNodeComplianceControl) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ModelCloudNodeComplianceControl) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ModelCloudNodeComplianceControl) SetTitle(v string) {
	o.Title = &v
}

func (o ModelCloudNodeComplianceControl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCloudNodeComplianceControl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.CategoryHierarchy != nil {
		toSerialize["category_hierarchy"] = o.CategoryHierarchy
	}
	if !IsNil(o.ControlId) {
		toSerialize["control_id"] = o.ControlId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Service) {
		toSerialize["service"] = o.Service
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return toSerialize, nil
}

type NullableModelCloudNodeComplianceControl struct {
	value *ModelCloudNodeComplianceControl
	isSet bool
}

func (v NullableModelCloudNodeComplianceControl) Get() *ModelCloudNodeComplianceControl {
	return v.value
}

func (v *NullableModelCloudNodeComplianceControl) Set(val *ModelCloudNodeComplianceControl) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCloudNodeComplianceControl) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCloudNodeComplianceControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCloudNodeComplianceControl(val *ModelCloudNodeComplianceControl) *NullableModelCloudNodeComplianceControl {
	return &NullableModelCloudNodeComplianceControl{value: val, isSet: true}
}

func (v NullableModelCloudNodeComplianceControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCloudNodeComplianceControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


