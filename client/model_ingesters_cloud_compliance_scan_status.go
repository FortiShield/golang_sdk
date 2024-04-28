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
	"time"
)

// checks if the IngestersCloudComplianceScanStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngestersCloudComplianceScanStatus{}

// IngestersCloudComplianceScanStatus struct for IngestersCloudComplianceScanStatus
type IngestersCloudComplianceScanStatus struct {
	Timestamp *time.Time `json:"@timestamp,omitempty"`
	ComplianceCheckTypes []string `json:"compliance_check_types,omitempty"`
	Result *IngestersComplianceStats `json:"result,omitempty"`
	ScanId *string `json:"scan_id,omitempty"`
	ScanMessage *string `json:"scan_message,omitempty"`
	ScanStatus *string `json:"scan_status,omitempty"`
	TotalChecks *int32 `json:"total_checks,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewIngestersCloudComplianceScanStatus instantiates a new IngestersCloudComplianceScanStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestersCloudComplianceScanStatus() *IngestersCloudComplianceScanStatus {
	this := IngestersCloudComplianceScanStatus{}
	return &this
}

// NewIngestersCloudComplianceScanStatusWithDefaults instantiates a new IngestersCloudComplianceScanStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestersCloudComplianceScanStatusWithDefaults() *IngestersCloudComplianceScanStatus {
	this := IngestersCloudComplianceScanStatus{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *IngestersCloudComplianceScanStatus) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudComplianceScanStatus) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *IngestersCloudComplianceScanStatus) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *IngestersCloudComplianceScanStatus) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetComplianceCheckTypes returns the ComplianceCheckTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *IngestersCloudComplianceScanStatus) GetComplianceCheckTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.ComplianceCheckTypes
}

// GetComplianceCheckTypesOk returns a tuple with the ComplianceCheckTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *IngestersCloudComplianceScanStatus) GetComplianceCheckTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ComplianceCheckTypes) {
		return nil, false
	}
	return o.ComplianceCheckTypes, true
}

// HasComplianceCheckTypes returns a boolean if a field has been set.
func (o *IngestersCloudComplianceScanStatus) HasComplianceCheckTypes() bool {
	if o != nil && !IsNil(o.ComplianceCheckTypes) {
		return true
	}

	return false
}

// SetComplianceCheckTypes gets a reference to the given []string and assigns it to the ComplianceCheckTypes field.
func (o *IngestersCloudComplianceScanStatus) SetComplianceCheckTypes(v []string) {
	o.ComplianceCheckTypes = v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *IngestersCloudComplianceScanStatus) GetResult() IngestersComplianceStats {
	if o == nil || IsNil(o.Result) {
		var ret IngestersComplianceStats
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudComplianceScanStatus) GetResultOk() (*IngestersComplianceStats, bool) {
	if o == nil || IsNil(o.Result) {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *IngestersCloudComplianceScanStatus) HasResult() bool {
	if o != nil && !IsNil(o.Result) {
		return true
	}

	return false
}

// SetResult gets a reference to the given IngestersComplianceStats and assigns it to the Result field.
func (o *IngestersCloudComplianceScanStatus) SetResult(v IngestersComplianceStats) {
	o.Result = &v
}

// GetScanId returns the ScanId field value if set, zero value otherwise.
func (o *IngestersCloudComplianceScanStatus) GetScanId() string {
	if o == nil || IsNil(o.ScanId) {
		var ret string
		return ret
	}
	return *o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudComplianceScanStatus) GetScanIdOk() (*string, bool) {
	if o == nil || IsNil(o.ScanId) {
		return nil, false
	}
	return o.ScanId, true
}

// HasScanId returns a boolean if a field has been set.
func (o *IngestersCloudComplianceScanStatus) HasScanId() bool {
	if o != nil && !IsNil(o.ScanId) {
		return true
	}

	return false
}

// SetScanId gets a reference to the given string and assigns it to the ScanId field.
func (o *IngestersCloudComplianceScanStatus) SetScanId(v string) {
	o.ScanId = &v
}

// GetScanMessage returns the ScanMessage field value if set, zero value otherwise.
func (o *IngestersCloudComplianceScanStatus) GetScanMessage() string {
	if o == nil || IsNil(o.ScanMessage) {
		var ret string
		return ret
	}
	return *o.ScanMessage
}

// GetScanMessageOk returns a tuple with the ScanMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudComplianceScanStatus) GetScanMessageOk() (*string, bool) {
	if o == nil || IsNil(o.ScanMessage) {
		return nil, false
	}
	return o.ScanMessage, true
}

// HasScanMessage returns a boolean if a field has been set.
func (o *IngestersCloudComplianceScanStatus) HasScanMessage() bool {
	if o != nil && !IsNil(o.ScanMessage) {
		return true
	}

	return false
}

// SetScanMessage gets a reference to the given string and assigns it to the ScanMessage field.
func (o *IngestersCloudComplianceScanStatus) SetScanMessage(v string) {
	o.ScanMessage = &v
}

// GetScanStatus returns the ScanStatus field value if set, zero value otherwise.
func (o *IngestersCloudComplianceScanStatus) GetScanStatus() string {
	if o == nil || IsNil(o.ScanStatus) {
		var ret string
		return ret
	}
	return *o.ScanStatus
}

// GetScanStatusOk returns a tuple with the ScanStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudComplianceScanStatus) GetScanStatusOk() (*string, bool) {
	if o == nil || IsNil(o.ScanStatus) {
		return nil, false
	}
	return o.ScanStatus, true
}

// HasScanStatus returns a boolean if a field has been set.
func (o *IngestersCloudComplianceScanStatus) HasScanStatus() bool {
	if o != nil && !IsNil(o.ScanStatus) {
		return true
	}

	return false
}

// SetScanStatus gets a reference to the given string and assigns it to the ScanStatus field.
func (o *IngestersCloudComplianceScanStatus) SetScanStatus(v string) {
	o.ScanStatus = &v
}

// GetTotalChecks returns the TotalChecks field value if set, zero value otherwise.
func (o *IngestersCloudComplianceScanStatus) GetTotalChecks() int32 {
	if o == nil || IsNil(o.TotalChecks) {
		var ret int32
		return ret
	}
	return *o.TotalChecks
}

// GetTotalChecksOk returns a tuple with the TotalChecks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudComplianceScanStatus) GetTotalChecksOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalChecks) {
		return nil, false
	}
	return o.TotalChecks, true
}

// HasTotalChecks returns a boolean if a field has been set.
func (o *IngestersCloudComplianceScanStatus) HasTotalChecks() bool {
	if o != nil && !IsNil(o.TotalChecks) {
		return true
	}

	return false
}

// SetTotalChecks gets a reference to the given int32 and assigns it to the TotalChecks field.
func (o *IngestersCloudComplianceScanStatus) SetTotalChecks(v int32) {
	o.TotalChecks = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IngestersCloudComplianceScanStatus) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestersCloudComplianceScanStatus) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IngestersCloudComplianceScanStatus) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *IngestersCloudComplianceScanStatus) SetType(v string) {
	o.Type = &v
}

func (o IngestersCloudComplianceScanStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngestersCloudComplianceScanStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["@timestamp"] = o.Timestamp
	}
	if o.ComplianceCheckTypes != nil {
		toSerialize["compliance_check_types"] = o.ComplianceCheckTypes
	}
	if !IsNil(o.Result) {
		toSerialize["result"] = o.Result
	}
	if !IsNil(o.ScanId) {
		toSerialize["scan_id"] = o.ScanId
	}
	if !IsNil(o.ScanMessage) {
		toSerialize["scan_message"] = o.ScanMessage
	}
	if !IsNil(o.ScanStatus) {
		toSerialize["scan_status"] = o.ScanStatus
	}
	if !IsNil(o.TotalChecks) {
		toSerialize["total_checks"] = o.TotalChecks
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIngestersCloudComplianceScanStatus struct {
	value *IngestersCloudComplianceScanStatus
	isSet bool
}

func (v NullableIngestersCloudComplianceScanStatus) Get() *IngestersCloudComplianceScanStatus {
	return v.value
}

func (v *NullableIngestersCloudComplianceScanStatus) Set(val *IngestersCloudComplianceScanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestersCloudComplianceScanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestersCloudComplianceScanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestersCloudComplianceScanStatus(val *IngestersCloudComplianceScanStatus) *NullableIngestersCloudComplianceScanStatus {
	return &NullableIngestersCloudComplianceScanStatus{value: val, isSet: true}
}

func (v NullableIngestersCloudComplianceScanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestersCloudComplianceScanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


