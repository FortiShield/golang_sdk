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

// checks if the ModelRegistryImagesReq type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelRegistryImagesReq{}

// ModelRegistryImagesReq struct for ModelRegistryImagesReq
type ModelRegistryImagesReq struct {
	ImageFilter ReportersFieldsFilters `json:"image_filter"`
	ImageStubFilter ReportersFieldsFilters `json:"image_stub_filter"`
	RegistryId string `json:"registry_id"`
	Window ModelFetchWindow `json:"window"`
}

type _ModelRegistryImagesReq ModelRegistryImagesReq

// NewModelRegistryImagesReq instantiates a new ModelRegistryImagesReq object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelRegistryImagesReq(imageFilter ReportersFieldsFilters, imageStubFilter ReportersFieldsFilters, registryId string, window ModelFetchWindow) *ModelRegistryImagesReq {
	this := ModelRegistryImagesReq{}
	this.ImageFilter = imageFilter
	this.ImageStubFilter = imageStubFilter
	this.RegistryId = registryId
	this.Window = window
	return &this
}

// NewModelRegistryImagesReqWithDefaults instantiates a new ModelRegistryImagesReq object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelRegistryImagesReqWithDefaults() *ModelRegistryImagesReq {
	this := ModelRegistryImagesReq{}
	return &this
}

// GetImageFilter returns the ImageFilter field value
func (o *ModelRegistryImagesReq) GetImageFilter() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.ImageFilter
}

// GetImageFilterOk returns a tuple with the ImageFilter field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryImagesReq) GetImageFilterOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageFilter, true
}

// SetImageFilter sets field value
func (o *ModelRegistryImagesReq) SetImageFilter(v ReportersFieldsFilters) {
	o.ImageFilter = v
}

// GetImageStubFilter returns the ImageStubFilter field value
func (o *ModelRegistryImagesReq) GetImageStubFilter() ReportersFieldsFilters {
	if o == nil {
		var ret ReportersFieldsFilters
		return ret
	}

	return o.ImageStubFilter
}

// GetImageStubFilterOk returns a tuple with the ImageStubFilter field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryImagesReq) GetImageStubFilterOk() (*ReportersFieldsFilters, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageStubFilter, true
}

// SetImageStubFilter sets field value
func (o *ModelRegistryImagesReq) SetImageStubFilter(v ReportersFieldsFilters) {
	o.ImageStubFilter = v
}

// GetRegistryId returns the RegistryId field value
func (o *ModelRegistryImagesReq) GetRegistryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryId
}

// GetRegistryIdOk returns a tuple with the RegistryId field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryImagesReq) GetRegistryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RegistryId, true
}

// SetRegistryId sets field value
func (o *ModelRegistryImagesReq) SetRegistryId(v string) {
	o.RegistryId = v
}

// GetWindow returns the Window field value
func (o *ModelRegistryImagesReq) GetWindow() ModelFetchWindow {
	if o == nil {
		var ret ModelFetchWindow
		return ret
	}

	return o.Window
}

// GetWindowOk returns a tuple with the Window field value
// and a boolean to check if the value has been set.
func (o *ModelRegistryImagesReq) GetWindowOk() (*ModelFetchWindow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Window, true
}

// SetWindow sets field value
func (o *ModelRegistryImagesReq) SetWindow(v ModelFetchWindow) {
	o.Window = v
}

func (o ModelRegistryImagesReq) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelRegistryImagesReq) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["image_filter"] = o.ImageFilter
	toSerialize["image_stub_filter"] = o.ImageStubFilter
	toSerialize["registry_id"] = o.RegistryId
	toSerialize["window"] = o.Window
	return toSerialize, nil
}

func (o *ModelRegistryImagesReq) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"image_filter",
		"image_stub_filter",
		"registry_id",
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

	varModelRegistryImagesReq := _ModelRegistryImagesReq{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelRegistryImagesReq)

	if err != nil {
		return err
	}

	*o = ModelRegistryImagesReq(varModelRegistryImagesReq)

	return err
}

type NullableModelRegistryImagesReq struct {
	value *ModelRegistryImagesReq
	isSet bool
}

func (v NullableModelRegistryImagesReq) Get() *ModelRegistryImagesReq {
	return v.value
}

func (v *NullableModelRegistryImagesReq) Set(val *ModelRegistryImagesReq) {
	v.value = val
	v.isSet = true
}

func (v NullableModelRegistryImagesReq) IsSet() bool {
	return v.isSet
}

func (v *NullableModelRegistryImagesReq) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelRegistryImagesReq(val *ModelRegistryImagesReq) *NullableModelRegistryImagesReq {
	return &NullableModelRegistryImagesReq{value: val, isSet: true}
}

func (v NullableModelRegistryImagesReq) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelRegistryImagesReq) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


