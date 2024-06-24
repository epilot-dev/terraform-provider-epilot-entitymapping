// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
)

func (r *EntityMappingDataSourceModel) RefreshFromSharedMappingConfigV2(resp *shared.MappingConfigV2) {
	if resp != nil {
		r.ID = types.StringValue(resp.ID)
		if resp.Source.Config == nil {
			r.Source.Config = nil
		} else {
			r.Source.Config = &tfTypes.Config{}
			if resp.Source.Config.EntityRef != nil {
				r.Source.Config.EntityRef = &tfTypes.EntityRef{}
				r.Source.Config.EntityRef.EntityID = types.StringValue(resp.Source.Config.EntityRef.EntityID)
				r.Source.Config.EntityRef.EntitySchema = types.StringPointerValue(resp.Source.Config.EntityRef.EntitySchema)
			}
			if resp.Source.Config.JourneyRef != nil {
				r.Source.Config.JourneyRef = &tfTypes.JourneyRef{}
				r.Source.Config.JourneyRef.JourneyID = types.StringPointerValue(resp.Source.Config.JourneyRef.JourneyID)
			}
		}
		if resp.Source.Type != nil {
			r.Source.Type = types.StringValue(string(*resp.Source.Type))
		} else {
			r.Source.Type = types.StringNull()
		}
		r.Targets = []tfTypes.TargetConfig{}
		if len(r.Targets) > len(resp.Targets) {
			r.Targets = r.Targets[:len(resp.Targets)]
		}
		for targetsCount, targetsItem := range resp.Targets {
			var targets1 tfTypes.TargetConfig
			targets1.AllowFailure = types.BoolPointerValue(targetsItem.AllowFailure)
			if targetsItem.ConditionMode == nil {
				targets1.ConditionMode = types.StringNull()
			} else {
				conditionModeResult, _ := json.Marshal(targetsItem.ConditionMode)
				targets1.ConditionMode = types.StringValue(string(conditionModeResult))
			}
			if targetsItem.Conditions == nil {
				targets1.Conditions = types.StringNull()
			} else {
				conditionsResult, _ := json.Marshal(targetsItem.Conditions)
				targets1.Conditions = types.StringValue(string(conditionsResult))
			}
			targets1.ID = types.StringPointerValue(targetsItem.ID)
			targets1.LinkbackRelationAttribute = types.StringPointerValue(targetsItem.LinkbackRelationAttribute)
			targets1.LinkbackRelationTags = []types.String{}
			for _, v := range targetsItem.LinkbackRelationTags {
				targets1.LinkbackRelationTags = append(targets1.LinkbackRelationTags, types.StringValue(v))
			}
			if targetsItem.LoopConfig == nil {
				targets1.LoopConfig = nil
			} else {
				targets1.LoopConfig = &tfTypes.LoopConfig{}
				if targetsItem.LoopConfig.Length != nil {
					targets1.LoopConfig.Length = types.NumberValue(big.NewFloat(float64(*targetsItem.LoopConfig.Length)))
				} else {
					targets1.LoopConfig.Length = types.NumberNull()
				}
				targets1.LoopConfig.SourcePath = types.StringPointerValue(targetsItem.LoopConfig.SourcePath)
			}
			targets1.MappingAttributes = []tfTypes.MappingAttributes{}
			for mappingAttributesCount, mappingAttributesItem := range targetsItem.MappingAttributes {
				var mappingAttributes1 tfTypes.MappingAttributes
				if mappingAttributesItem.MappingAttribute != nil {
					mappingAttributes1.MappingAttribute = &tfTypes.MappingAttribute{}
					if mappingAttributesItem.MappingAttribute.AppendValueMapper != nil {
						mappingAttributes1.MappingAttribute.AppendValueMapper = &tfTypes.AppendValueMapper{}
						mappingAttributes1.MappingAttribute.AppendValueMapper.Mode = types.StringValue(string(mappingAttributesItem.MappingAttribute.AppendValueMapper.Mode))
						mappingAttributes1.MappingAttribute.AppendValueMapper.Source = types.StringPointerValue(mappingAttributesItem.MappingAttribute.AppendValueMapper.Source)
						mappingAttributes1.MappingAttribute.AppendValueMapper.Target = types.StringValue(mappingAttributesItem.MappingAttribute.AppendValueMapper.Target)
						mappingAttributes1.MappingAttribute.AppendValueMapper.TargetUnique = []types.String{}
						for _, v := range mappingAttributesItem.MappingAttribute.AppendValueMapper.TargetUnique {
							mappingAttributes1.MappingAttribute.AppendValueMapper.TargetUnique = append(mappingAttributes1.MappingAttribute.AppendValueMapper.TargetUnique, types.StringValue(v))
						}
						mappingAttributes1.MappingAttribute.AppendValueMapper.ValueJSON = types.StringValue(mappingAttributesItem.MappingAttribute.AppendValueMapper.ValueJSON)
					}
					if mappingAttributesItem.MappingAttribute.CopyValueMapper != nil {
						mappingAttributes1.MappingAttribute.CopyValueMapper = &tfTypes.CopyValueMapper{}
						mappingAttributes1.MappingAttribute.CopyValueMapper.Mode = types.StringValue(string(mappingAttributesItem.MappingAttribute.CopyValueMapper.Mode))
						mappingAttributes1.MappingAttribute.CopyValueMapper.Source = types.StringValue(mappingAttributesItem.MappingAttribute.CopyValueMapper.Source)
						mappingAttributes1.MappingAttribute.CopyValueMapper.Target = types.StringValue(mappingAttributesItem.MappingAttribute.CopyValueMapper.Target)
					}
					if mappingAttributesItem.MappingAttribute.SetValueMapper != nil {
						mappingAttributes1.MappingAttribute.SetValueMapper = &tfTypes.SetValueMapper{}
						mappingAttributes1.MappingAttribute.SetValueMapper.Mode = types.StringValue(string(mappingAttributesItem.MappingAttribute.SetValueMapper.Mode))
						mappingAttributes1.MappingAttribute.SetValueMapper.Target = types.StringValue(mappingAttributesItem.MappingAttribute.SetValueMapper.Target)
						valueResult, _ := json.Marshal(mappingAttributesItem.MappingAttribute.SetValueMapper.Value)
						mappingAttributes1.MappingAttribute.SetValueMapper.Value = types.StringValue(string(valueResult))
					}
				}
				if mappingAttributesItem.MappingAttributeV2 != nil {
					mappingAttributes1.MappingAttributeV2 = &tfTypes.MappingAttributeV2{}
					if mappingAttributesItem.MappingAttributeV2.Operation.Any != nil {
						anyResult, _ := json.Marshal(mappingAttributesItem.MappingAttributeV2.Operation.Any)
						mappingAttributes1.MappingAttributeV2.Operation.Any = types.StringValue(string(anyResult))
					}
					if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode != nil {
						mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode = &tfTypes.OperationObjectNode{}
						mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Append = nil
						for _, appendItem := range mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Append {
							var append2 types.String
							append2Result, _ := json.Marshal(appendItem)
							append2 = types.StringValue(string(append2Result))
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Append = append(mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Append, append2)
						}
						mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Copy = types.StringPointerValue(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Copy)
						if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random == nil {
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random = nil
						} else {
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random = &tfTypes.RandomOperation{}
							if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.One != nil {
								mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random.One = &tfTypes.One{}
								mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random.One.Type = types.StringValue(string(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.One.Type))
							}
							if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two != nil {
								mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random.Two = &tfTypes.Two{}
								if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Max != nil {
									mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Max = types.NumberValue(big.NewFloat(float64(*mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Max)))
								} else {
									mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Max = types.NumberNull()
								}
								if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Min != nil {
									mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Min = types.NumberValue(big.NewFloat(float64(*mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Min)))
								} else {
									mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Min = types.NumberNull()
								}
								mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Type = types.StringValue(string(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Type))
							}
						}
						if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Set == nil {
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Set = types.StringNull()
						} else {
							setResult, _ := json.Marshal(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Set)
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Set = types.StringValue(string(setResult))
						}
						mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Template = types.StringPointerValue(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Template)
						if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq == nil {
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Uniq = nil
						} else {
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Uniq = &tfTypes.Uniq{}
							if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq.Boolean != nil {
								mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Uniq.Boolean = types.BoolPointerValue(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq.Boolean)
							}
							if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq.ArrayOfStr != nil {
								mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Uniq.ArrayOfStr = []types.String{}
								for _, v := range mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq.ArrayOfStr {
									mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Uniq.ArrayOfStr = append(mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.Uniq.ArrayOfStr, types.StringValue(v))
								}
							}
						}
						if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.AdditionalProperties == nil {
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.AdditionalProperties = types.StringNull()
						} else {
							additionalPropertiesResult, _ := json.Marshal(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.AdditionalProperties)
							mappingAttributes1.MappingAttributeV2.Operation.OperationObjectNode.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
						}
					}
					if mappingAttributesItem.MappingAttributeV2.Origin != nil {
						mappingAttributes1.MappingAttributeV2.Origin = types.StringValue(string(*mappingAttributesItem.MappingAttributeV2.Origin))
					} else {
						mappingAttributes1.MappingAttributeV2.Origin = types.StringNull()
					}
					mappingAttributes1.MappingAttributeV2.Target = types.StringValue(mappingAttributesItem.MappingAttributeV2.Target)
				}
				if mappingAttributesCount+1 > len(targets1.MappingAttributes) {
					targets1.MappingAttributes = append(targets1.MappingAttributes, mappingAttributes1)
				} else {
					targets1.MappingAttributes[mappingAttributesCount].MappingAttribute = mappingAttributes1.MappingAttribute
					targets1.MappingAttributes[mappingAttributesCount].MappingAttributeV2 = mappingAttributes1.MappingAttributeV2
				}
			}
			targets1.Name = types.StringPointerValue(targetsItem.Name)
			targets1.RelationAttributes = []tfTypes.RelationAttribute{}
			for relationAttributesCount, relationAttributesItem := range targetsItem.RelationAttributes {
				var relationAttributes1 tfTypes.RelationAttribute
				relationAttributes1.Mode = types.StringValue(string(relationAttributesItem.Mode))
				if relationAttributesItem.Origin != nil {
					relationAttributes1.Origin = types.StringValue(string(*relationAttributesItem.Origin))
				} else {
					relationAttributes1.Origin = types.StringNull()
				}
				if len(relationAttributesItem.RelatedTo) > 0 {
					relationAttributes1.RelatedTo = make(map[string]types.String)
					for key, value1 := range relationAttributesItem.RelatedTo {
						result, _ := json.Marshal(value1)
						relationAttributes1.RelatedTo[key] = types.StringValue(string(result))
					}
				}
				if relationAttributesItem.SourceFilter == nil {
					relationAttributes1.SourceFilter = nil
				} else {
					relationAttributes1.SourceFilter = &tfTypes.SourceFilter{}
					relationAttributes1.SourceFilter.Attribute = types.StringPointerValue(relationAttributesItem.SourceFilter.Attribute)
					relationAttributes1.SourceFilter.Limit = types.Int64PointerValue(relationAttributesItem.SourceFilter.Limit)
					relationAttributes1.SourceFilter.RelationTag = types.StringPointerValue(relationAttributesItem.SourceFilter.RelationTag)
					relationAttributes1.SourceFilter.Schema = types.StringPointerValue(relationAttributesItem.SourceFilter.Schema)
					relationAttributes1.SourceFilter.Self = types.BoolPointerValue(relationAttributesItem.SourceFilter.Self)
					relationAttributes1.SourceFilter.Tag = types.StringPointerValue(relationAttributesItem.SourceFilter.Tag)
				}
				relationAttributes1.Target = types.StringValue(relationAttributesItem.Target)
				relationAttributes1.TargetTags = []types.String{}
				for _, v := range relationAttributesItem.TargetTags {
					relationAttributes1.TargetTags = append(relationAttributes1.TargetTags, types.StringValue(v))
				}
				relationAttributes1.TargetTagsIncludeSource = types.BoolPointerValue(relationAttributesItem.TargetTagsIncludeSource)
				if relationAttributesCount+1 > len(targets1.RelationAttributes) {
					targets1.RelationAttributes = append(targets1.RelationAttributes, relationAttributes1)
				} else {
					targets1.RelationAttributes[relationAttributesCount].Mode = relationAttributes1.Mode
					targets1.RelationAttributes[relationAttributesCount].Origin = relationAttributes1.Origin
					targets1.RelationAttributes[relationAttributesCount].RelatedTo = relationAttributes1.RelatedTo
					targets1.RelationAttributes[relationAttributesCount].SourceFilter = relationAttributes1.SourceFilter
					targets1.RelationAttributes[relationAttributesCount].Target = relationAttributes1.Target
					targets1.RelationAttributes[relationAttributesCount].TargetTags = relationAttributes1.TargetTags
					targets1.RelationAttributes[relationAttributesCount].TargetTagsIncludeSource = relationAttributes1.TargetTagsIncludeSource
				}
			}
			targets1.TargetSchema = types.StringValue(targetsItem.TargetSchema)
			targets1.TargetUnique = []types.String{}
			for _, v := range targetsItem.TargetUnique {
				targets1.TargetUnique = append(targets1.TargetUnique, types.StringValue(v))
			}
			if targetsCount+1 > len(r.Targets) {
				r.Targets = append(r.Targets, targets1)
			} else {
				r.Targets[targetsCount].AllowFailure = targets1.AllowFailure
				r.Targets[targetsCount].ConditionMode = targets1.ConditionMode
				r.Targets[targetsCount].Conditions = targets1.Conditions
				r.Targets[targetsCount].ID = targets1.ID
				r.Targets[targetsCount].LinkbackRelationAttribute = targets1.LinkbackRelationAttribute
				r.Targets[targetsCount].LinkbackRelationTags = targets1.LinkbackRelationTags
				r.Targets[targetsCount].LoopConfig = targets1.LoopConfig
				r.Targets[targetsCount].MappingAttributes = targets1.MappingAttributes
				r.Targets[targetsCount].Name = targets1.Name
				r.Targets[targetsCount].RelationAttributes = targets1.RelationAttributes
				r.Targets[targetsCount].TargetSchema = targets1.TargetSchema
				r.Targets[targetsCount].TargetUnique = targets1.TargetUnique
			}
		}
	}
}
