package teamcity

import (
	"encoding/json"
	"errors"
)

type CommonBuildFeature struct {
	id          string
	featureType string
	buildTypeID string
	vcsRootID   string
	disabled    bool
	properties *Properties
}

func (bf *CommonBuildFeature) ID() string {
	return bf.id
}

func (bf *CommonBuildFeature) SetID(value string) {
	bf.id = value
}

func (bf *CommonBuildFeature) Type() string {
	return bf.featureType
}

func (bf *CommonBuildFeature) VcsRootID() string {
	return bf.vcsRootID
}

func (bf *CommonBuildFeature) SetVcsRootID(value string) {
	bf.vcsRootID = value
}

func (bf *CommonBuildFeature) Properties() *Properties {
	return bf.properties
}

func (bf *CommonBuildFeature) BuildTypeID() string {
	return bf.buildTypeID
}

func (bf *CommonBuildFeature) SetBuildTypeID(value string) {
	bf.buildTypeID = value
}

func (bf *CommonBuildFeature) Disabled() bool {
	return bf.disabled
}

func (bf *CommonBuildFeature) SetDisabled(value bool) {
	bf.disabled = value
}

func (bf *CommonBuildFeature) MarshalJSON() ([]byte, error) {
	out := &buildFeatureJSON{
		ID:         bf.id,
		Disabled:   NewBool(bf.disabled),
		Properties: bf.properties,
		Inherited:  NewFalse(),
		Type:       bf.Type(),
	}

	if bf.vcsRootID != "" {
		out.Properties.AddOrReplaceValue("vcsRootId", bf.vcsRootID)
	}
	return json.Marshal(out)
}

func (bf *CommonBuildFeature) UnmarshalJSON(data []byte) error {
	var aux buildFeatureJSON
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	bf.id = aux.ID

	disabled := aux.Disabled
	if disabled == nil {
		disabled = NewFalse()
	}
	bf.disabled = *disabled

	if aux.Properties != nil {
		bf.properties = NewProperties(aux.Properties.Items...)
	}

	return nil
}

func NewCommonBuildFeature(featureType string, propertiesRaw []interface {}) (*CommonBuildFeature, error) {
	properties := NewPropertiesEmpty()
	for _, propertyRaw := range propertiesRaw {
		var name, value string
		var ok bool
		propertyMap := propertyRaw.(map[string]interface{})
		if name, ok = propertyMap["name"].(string); !ok {
			return nil, errors.New("missing name in property")
		}
		if value, ok = propertyMap["value"].(string); !ok {
			return nil, errors.New("missing value in property")
		}
		properties.Add(&Property{
			Name: name,
			Value: value,
		})
	}

	return &CommonBuildFeature{
		featureType: featureType,
		properties: properties,
	}, nil
}