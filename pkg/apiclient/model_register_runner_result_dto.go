/*
Daytona Server API

Daytona Server API

API version: v0.0.0-dev
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apiclient

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the RegisterRunnerResultDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegisterRunnerResultDTO{}

// RegisterRunnerResultDTO struct for RegisterRunnerResultDTO
type RegisterRunnerResultDTO struct {
	Alias    string          `json:"alias"`
	ApiKey   string          `json:"apiKey"`
	Id       string          `json:"id"`
	Metadata *RunnerMetadata `json:"metadata,omitempty"`
}

type _RegisterRunnerResultDTO RegisterRunnerResultDTO

// NewRegisterRunnerResultDTO instantiates a new RegisterRunnerResultDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegisterRunnerResultDTO(alias string, apiKey string, id string) *RegisterRunnerResultDTO {
	this := RegisterRunnerResultDTO{}
	this.Alias = alias
	this.ApiKey = apiKey
	this.Id = id
	return &this
}

// NewRegisterRunnerResultDTOWithDefaults instantiates a new RegisterRunnerResultDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegisterRunnerResultDTOWithDefaults() *RegisterRunnerResultDTO {
	this := RegisterRunnerResultDTO{}
	return &this
}

// GetAlias returns the Alias field value
func (o *RegisterRunnerResultDTO) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *RegisterRunnerResultDTO) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *RegisterRunnerResultDTO) SetAlias(v string) {
	o.Alias = v
}

// GetApiKey returns the ApiKey field value
func (o *RegisterRunnerResultDTO) GetApiKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiKey
}

// GetApiKeyOk returns a tuple with the ApiKey field value
// and a boolean to check if the value has been set.
func (o *RegisterRunnerResultDTO) GetApiKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiKey, true
}

// SetApiKey sets field value
func (o *RegisterRunnerResultDTO) SetApiKey(v string) {
	o.ApiKey = v
}

// GetId returns the Id field value
func (o *RegisterRunnerResultDTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RegisterRunnerResultDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RegisterRunnerResultDTO) SetId(v string) {
	o.Id = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RegisterRunnerResultDTO) GetMetadata() RunnerMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret RunnerMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegisterRunnerResultDTO) GetMetadataOk() (*RunnerMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RegisterRunnerResultDTO) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given RunnerMetadata and assigns it to the Metadata field.
func (o *RegisterRunnerResultDTO) SetMetadata(v RunnerMetadata) {
	o.Metadata = &v
}

func (o RegisterRunnerResultDTO) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegisterRunnerResultDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["apiKey"] = o.ApiKey
	toSerialize["id"] = o.Id
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	return toSerialize, nil
}

func (o *RegisterRunnerResultDTO) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"apiKey",
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRegisterRunnerResultDTO := _RegisterRunnerResultDTO{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRegisterRunnerResultDTO)

	if err != nil {
		return err
	}

	*o = RegisterRunnerResultDTO(varRegisterRunnerResultDTO)

	return err
}

type NullableRegisterRunnerResultDTO struct {
	value *RegisterRunnerResultDTO
	isSet bool
}

func (v NullableRegisterRunnerResultDTO) Get() *RegisterRunnerResultDTO {
	return v.value
}

func (v *NullableRegisterRunnerResultDTO) Set(val *RegisterRunnerResultDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableRegisterRunnerResultDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableRegisterRunnerResultDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegisterRunnerResultDTO(val *RegisterRunnerResultDTO) *NullableRegisterRunnerResultDTO {
	return &NullableRegisterRunnerResultDTO{value: val, isSet: true}
}

func (v NullableRegisterRunnerResultDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegisterRunnerResultDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}