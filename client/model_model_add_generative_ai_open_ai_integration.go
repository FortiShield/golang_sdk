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

// checks if the ModelAddGenerativeAiOpenAIIntegration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelAddGenerativeAiOpenAIIntegration{}

// ModelAddGenerativeAiOpenAIIntegration struct for ModelAddGenerativeAiOpenAIIntegration
type ModelAddGenerativeAiOpenAIIntegration struct {
	ApiKey string `json:"api_key"`
	ModelId string `json:"model_id"`
}

type _ModelAddGenerativeAiOpenAIIntegration ModelAddGenerativeAiOpenAIIntegration

// NewModelAddGenerativeAiOpenAIIntegration instantiates a new ModelAddGenerativeAiOpenAIIntegration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelAddGenerativeAiOpenAIIntegration(apiKey string, modelId string) *ModelAddGenerativeAiOpenAIIntegration {
	this := ModelAddGenerativeAiOpenAIIntegration{}
	this.ApiKey = apiKey
	this.ModelId = modelId
	return &this
}

// NewModelAddGenerativeAiOpenAIIntegrationWithDefaults instantiates a new ModelAddGenerativeAiOpenAIIntegration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelAddGenerativeAiOpenAIIntegrationWithDefaults() *ModelAddGenerativeAiOpenAIIntegration {
	this := ModelAddGenerativeAiOpenAIIntegration{}
	return &this
}

// GetApiKey returns the ApiKey field value
func (o *ModelAddGenerativeAiOpenAIIntegration) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *ModelAddGenerativeAiOpenAIIntegration) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *ModelAddGenerativeAiOpenAIIntegration) SetApiKey(v string) {
	o.ApiKey = v
}

// GetModelId returns the ModelId field value
func (o *ModelAddGenerativeAiOpenAIIntegration) GetModelId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModelId
}

// GetModelIdOk returns a tuple with the ModelId field value
// and a boolean to check if the value has been set.
func (o *ModelAddGenerativeAiOpenAIIntegration) GetModelIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelId, true
}

// SetModelId sets field value
func (o *ModelAddGenerativeAiOpenAIIntegration) SetModelId(v string) {
	o.ModelId = v
}

func (o ModelAddGenerativeAiOpenAIIntegration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelAddGenerativeAiOpenAIIntegration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["api_key"] = o.ApiKey
	toSerialize["model_id"] = o.ModelId
	return toSerialize, nil
}

func (o *ModelAddGenerativeAiOpenAIIntegration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"api_key",
		"model_id",
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

	varModelAddGenerativeAiOpenAIIntegration := _ModelAddGenerativeAiOpenAIIntegration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelAddGenerativeAiOpenAIIntegration)

	if err != nil {
		return err
	}

	*o = ModelAddGenerativeAiOpenAIIntegration(varModelAddGenerativeAiOpenAIIntegration)

	return err
}

type NullableModelAddGenerativeAiOpenAIIntegration struct {
	value *ModelAddGenerativeAiOpenAIIntegration
	isSet bool
}

func (v NullableModelAddGenerativeAiOpenAIIntegration) Get() *ModelAddGenerativeAiOpenAIIntegration {
	return v.value
}

func (v *NullableModelAddGenerativeAiOpenAIIntegration) Set(val *ModelAddGenerativeAiOpenAIIntegration) {
	v.value = val
	v.isSet = true
}

func (v NullableModelAddGenerativeAiOpenAIIntegration) IsSet() bool {
	return v.isSet
}

func (v *NullableModelAddGenerativeAiOpenAIIntegration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelAddGenerativeAiOpenAIIntegration(val *ModelAddGenerativeAiOpenAIIntegration) *NullableModelAddGenerativeAiOpenAIIntegration {
	return &NullableModelAddGenerativeAiOpenAIIntegration{value: val, isSet: true}
}

func (v NullableModelAddGenerativeAiOpenAIIntegration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelAddGenerativeAiOpenAIIntegration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


