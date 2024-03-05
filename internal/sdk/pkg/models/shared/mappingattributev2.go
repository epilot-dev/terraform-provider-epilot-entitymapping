// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type MappingAttributeV2 struct {
	// Mapping operation nodes are either primitive values or operation node objects
	Operation OperationNode `json:"operation"`
	// Origin of an attribute.
	Origin *AttributeOrigin `json:"origin,omitempty"`
	// Target JSON path for the attribute to set
	Target string `json:"target"`
}

func (o *MappingAttributeV2) GetOperation() OperationNode {
	if o == nil {
		return OperationNode{}
	}
	return o.Operation
}

func (o *MappingAttributeV2) GetOrigin() *AttributeOrigin {
	if o == nil {
		return nil
	}
	return o.Origin
}

func (o *MappingAttributeV2) GetTarget() string {
	if o == nil {
		return ""
	}
	return o.Target
}