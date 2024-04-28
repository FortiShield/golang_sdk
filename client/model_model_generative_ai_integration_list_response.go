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

// checks if the ModelGenerativeAiIntegrationListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelGenerativeAiIntegrationListResponse{}

// ModelGenerativeAiIntegrationListResponse struct for ModelGenerativeAiIntegrationListResponse
type ModelGenerativeAiIntegrationListResponse struct {
	DefaultIntegration *bool `json:"default_integration,omitempty"`
	Id *int32 `json:"id,omitempty"`
	IntegrationType *string `json:"integration_type,omitempty"`
	Label *string `json:"label,omitempty"`
	LastErrorMsg *string `json:"last_error_msg,omitempty"`
}

// NewModelGenerativeAiIntegrationListResponse instantiates a new ModelGenerativeAiIntegrationListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelGenerativeAiIntegrationListResponse() *ModelGenerativeAiIntegrationListResponse {
	this := ModelGenerativeAiIntegrationListResponse{}
	return &this
}

// NewModelGenerativeAiIntegrationListResponseWithDefaults instantiates a new ModelGenerativeAiIntegrationListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelGenerativeAiIntegrationListResponseWithDefaults() *ModelGenerativeAiIntegrationListResponse {
	this := ModelGenerativeAiIntegrationListResponse{}
	return &this
}

// GetDefaultIntegration returns the DefaultIntegration field value if set, zero value otherwise.
func (o *ModelGenerativeAiIntegrationListResponse) GetDefaultIntegration() bool {
	if o == nil || IsNil(o.DefaultIntegration) {
		var ret bool
		return ret
	}
	return *o.DefaultIntegration
}

// GetDefaultIntegrationOk returns a tuple with the DefaultIntegration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationListResponse) GetDefaultIntegrationOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultIntegration) {
		return nil, false
	}
	return o.DefaultIntegration, true
}

// HasDefaultIntegration returns a boolean if a field has been set.
func (o *ModelGenerativeAiIntegrationListResponse) HasDefaultIntegration() bool {
	if o != nil && !IsNil(o.DefaultIntegration) {
		return true
	}

	return false
}

// SetDefaultIntegration gets a reference to the given bool and assigns it to the DefaultIntegration field.
func (o *ModelGenerativeAiIntegrationListResponse) SetDefaultIntegration(v bool) {
	o.DefaultIntegration = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ModelGenerativeAiIntegrationListResponse) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationListResponse) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ModelGenerativeAiIntegrationListResponse) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ModelGenerativeAiIntegrationListResponse) SetId(v int32) {
	o.Id = &v
}

// GetIntegrationType returns the IntegrationType field value if set, zero value otherwise.
func (o *ModelGenerativeAiIntegrationListResponse) GetIntegrationType() string {
	if o == nil || IsNil(o.IntegrationType) {
		var ret string
		return ret
	}
	return *o.IntegrationType
}

// GetIntegrationTypeOk returns a tuple with the IntegrationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationListResponse) GetIntegrationTypeOk() (*string, bool) {
	if o == nil || IsNil(o.IntegrationType) {
		return nil, false
	}
	return o.IntegrationType, true
}

// HasIntegrationType returns a boolean if a field has been set.
func (o *ModelGenerativeAiIntegrationListResponse) HasIntegrationType() bool {
	if o != nil && !IsNil(o.IntegrationType) {
		return true
	}

	return false
}

// SetIntegrationType gets a reference to the given string and assigns it to the IntegrationType field.
func (o *ModelGenerativeAiIntegrationListResponse) SetIntegrationType(v string) {
	o.IntegrationType = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ModelGenerativeAiIntegrationListResponse) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationListResponse) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ModelGenerativeAiIntegrationListResponse) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ModelGenerativeAiIntegrationListResponse) SetLabel(v string) {
	o.Label = &v
}

// GetLastErrorMsg returns the LastErrorMsg field value if set, zero value otherwise.
func (o *ModelGenerativeAiIntegrationListResponse) GetLastErrorMsg() string {
	if o == nil || IsNil(o.LastErrorMsg) {
		var ret string
		return ret
	}
	return *o.LastErrorMsg
}

// GetLastErrorMsgOk returns a tuple with the LastErrorMsg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelGenerativeAiIntegrationListResponse) GetLastErrorMsgOk() (*string, bool) {
	if o == nil || IsNil(o.LastErrorMsg) {
		return nil, false
	}
	return o.LastErrorMsg, true
}

// HasLastErrorMsg returns a boolean if a field has been set.
func (o *ModelGenerativeAiIntegrationListResponse) HasLastErrorMsg() bool {
	if o != nil && !IsNil(o.LastErrorMsg) {
		return true
	}

	return false
}

// SetLastErrorMsg gets a reference to the given string and assigns it to the LastErrorMsg field.
func (o *ModelGenerativeAiIntegrationListResponse) SetLastErrorMsg(v string) {
	o.LastErrorMsg = &v
}

func (o ModelGenerativeAiIntegrationListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelGenerativeAiIntegrationListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DefaultIntegration) {
		toSerialize["default_integration"] = o.DefaultIntegration
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.IntegrationType) {
		toSerialize["integration_type"] = o.IntegrationType
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.LastErrorMsg) {
		toSerialize["last_error_msg"] = o.LastErrorMsg
	}
	return toSerialize, nil
}

type NullableModelGenerativeAiIntegrationListResponse struct {
	value *ModelGenerativeAiIntegrationListResponse
	isSet bool
}

func (v NullableModelGenerativeAiIntegrationListResponse) Get() *ModelGenerativeAiIntegrationListResponse {
	return v.value
}

func (v *NullableModelGenerativeAiIntegrationListResponse) Set(val *ModelGenerativeAiIntegrationListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableModelGenerativeAiIntegrationListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableModelGenerativeAiIntegrationListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelGenerativeAiIntegrationListResponse(val *ModelGenerativeAiIntegrationListResponse) *NullableModelGenerativeAiIntegrationListResponse {
	return &NullableModelGenerativeAiIntegrationListResponse{value: val, isSet: true}
}

func (v NullableModelGenerativeAiIntegrationListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelGenerativeAiIntegrationListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


