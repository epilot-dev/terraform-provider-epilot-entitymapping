// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"math/big"
)

func (r *EntityMappingResourceModel) ToSharedMappingConfigV2() *shared.MappingConfigV2 {
	var id string
	id = r.ID.ValueString()

	var config *shared.Config
	if r.Source.Config != nil {
		var journeyRef *shared.JourneyRef
		if r.Source.Config.JourneyRef != nil {
			journeyID := new(string)
			if !r.Source.Config.JourneyRef.JourneyID.IsUnknown() && !r.Source.Config.JourneyRef.JourneyID.IsNull() {
				*journeyID = r.Source.Config.JourneyRef.JourneyID.ValueString()
			} else {
				journeyID = nil
			}
			journeyRef = &shared.JourneyRef{
				JourneyID: journeyID,
			}
		}
		if journeyRef != nil {
			config = &shared.Config{
				JourneyRef: journeyRef,
			}
		}
		var entityRef *shared.EntityRef
		if r.Source.Config.EntityRef != nil {
			var entityID string
			entityID = r.Source.Config.EntityRef.EntityID.ValueString()

			entitySchema := new(string)
			if !r.Source.Config.EntityRef.EntitySchema.IsUnknown() && !r.Source.Config.EntityRef.EntitySchema.IsNull() {
				*entitySchema = r.Source.Config.EntityRef.EntitySchema.ValueString()
			} else {
				entitySchema = nil
			}
			entityRef = &shared.EntityRef{
				EntityID:     entityID,
				EntitySchema: entitySchema,
			}
		}
		if entityRef != nil {
			config = &shared.Config{
				EntityRef: entityRef,
			}
		}
	}
	typeVar := new(shared.Type)
	if !r.Source.Type.IsUnknown() && !r.Source.Type.IsNull() {
		*typeVar = shared.Type(r.Source.Type.ValueString())
	} else {
		typeVar = nil
	}
	source := shared.SourceConfig{
		Config: config,
		Type:   typeVar,
	}
	var targets []shared.TargetConfig = []shared.TargetConfig{}
	for _, targetsItem := range r.Targets {
		allowFailure := new(bool)
		if !targetsItem.AllowFailure.IsUnknown() && !targetsItem.AllowFailure.IsNull() {
			*allowFailure = targetsItem.AllowFailure.ValueBool()
		} else {
			allowFailure = nil
		}
		var conditionMode interface{}
		if !targetsItem.ConditionMode.IsUnknown() && !targetsItem.ConditionMode.IsNull() {
			_ = json.Unmarshal([]byte(targetsItem.ConditionMode.ValueString()), &conditionMode)
		}
		var conditions interface{}
		if !targetsItem.Conditions.IsUnknown() && !targetsItem.Conditions.IsNull() {
			_ = json.Unmarshal([]byte(targetsItem.Conditions.ValueString()), &conditions)
		}
		id1 := new(string)
		if !targetsItem.ID.IsUnknown() && !targetsItem.ID.IsNull() {
			*id1 = targetsItem.ID.ValueString()
		} else {
			id1 = nil
		}
		linkbackRelationAttribute := new(string)
		if !targetsItem.LinkbackRelationAttribute.IsUnknown() && !targetsItem.LinkbackRelationAttribute.IsNull() {
			*linkbackRelationAttribute = targetsItem.LinkbackRelationAttribute.ValueString()
		} else {
			linkbackRelationAttribute = nil
		}
		var linkbackRelationTags []string = []string{}
		for _, linkbackRelationTagsItem := range targetsItem.LinkbackRelationTags {
			linkbackRelationTags = append(linkbackRelationTags, linkbackRelationTagsItem.ValueString())
		}
		var loopConfig *shared.LoopConfig
		if targetsItem.LoopConfig != nil {
			length := new(float64)
			if !targetsItem.LoopConfig.Length.IsUnknown() && !targetsItem.LoopConfig.Length.IsNull() {
				*length, _ = targetsItem.LoopConfig.Length.ValueBigFloat().Float64()
			} else {
				length = nil
			}
			sourcePath := new(string)
			if !targetsItem.LoopConfig.SourcePath.IsUnknown() && !targetsItem.LoopConfig.SourcePath.IsNull() {
				*sourcePath = targetsItem.LoopConfig.SourcePath.ValueString()
			} else {
				sourcePath = nil
			}
			loopConfig = &shared.LoopConfig{
				Length:     length,
				SourcePath: sourcePath,
			}
		}
		var mappingAttributes []shared.MappingAttributes = []shared.MappingAttributes{}
		for _, mappingAttributesItem := range targetsItem.MappingAttributes {
			if mappingAttributesItem.MappingAttributeV2 != nil {
				var operation shared.OperationNode
				var operationObjectNode *shared.OperationObjectNode
				if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode != nil {
					var additionalProperties interface{}
					if !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.AdditionalProperties.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.AdditionalProperties.IsNull() {
						_ = json.Unmarshal([]byte(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.AdditionalProperties.ValueString()), &additionalProperties)
					}
					var append1 []interface{} = []interface{}{}
					for _, appendItem := range mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Append {
						var appendTmp interface{}
						_ = json.Unmarshal([]byte(appendItem.ValueString()), &appendTmp)
						append1 = append(append1, appendTmp)
					}
					copy := new(string)
					if !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Copy.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Copy.IsNull() {
						*copy = mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Copy.ValueString()
					} else {
						copy = nil
					}
					var random *shared.RandomOperation
					if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random != nil {
						var one *shared.One
						if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.One != nil {
							typeVar1 := shared.RandomOperationSchemasType(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.One.Type.ValueString())
							one = &shared.One{
								Type: typeVar1,
							}
						}
						if one != nil {
							random = &shared.RandomOperation{
								One: one,
							}
						}
						var two *shared.Two
						if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two != nil {
							max := new(float64)
							if !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Max.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Max.IsNull() {
								*max, _ = mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Max.ValueBigFloat().Float64()
							} else {
								max = nil
							}
							min := new(float64)
							if !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Min.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Min.IsNull() {
								*min, _ = mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Min.ValueBigFloat().Float64()
							} else {
								min = nil
							}
							typeVar2 := shared.RandomOperationType(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Random.Two.Type.ValueString())
							two = &shared.Two{
								Max:  max,
								Min:  min,
								Type: typeVar2,
							}
						}
						if two != nil {
							random = &shared.RandomOperation{
								Two: two,
							}
						}
					}
					var set interface{}
					if !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Set.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Set.IsNull() {
						_ = json.Unmarshal([]byte(mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Set.ValueString()), &set)
					}
					template := new(string)
					if !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Template.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Template.IsNull() {
						*template = mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Template.ValueString()
					} else {
						template = nil
					}
					var uniq *shared.Uniq
					if mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq != nil {
						boolean := new(bool)
						if !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq.Boolean.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq.Boolean.IsNull() {
							*boolean = mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq.Boolean.ValueBool()
						} else {
							boolean = nil
						}
						if boolean != nil {
							uniq = &shared.Uniq{
								Boolean: boolean,
							}
						}
						var arrayOfStr []string = []string{}
						for _, arrayOfStrItem := range mappingAttributesItem.MappingAttributeV2.Operation.OperationObjectNode.Uniq.ArrayOfStr {
							arrayOfStr = append(arrayOfStr, arrayOfStrItem.ValueString())
						}
						if arrayOfStr != nil {
							uniq = &shared.Uniq{
								ArrayOfStr: arrayOfStr,
							}
						}
					}
					operationObjectNode = &shared.OperationObjectNode{
						AdditionalProperties: additionalProperties,
						Append:               append1,
						Copy:                 copy,
						Random:               random,
						Set:                  set,
						Template:             template,
						Uniq:                 uniq,
					}
				}
				if operationObjectNode != nil {
					operation = shared.OperationNode{
						OperationObjectNode: operationObjectNode,
					}
				}
				var anyVar interface{}
				if !mappingAttributesItem.MappingAttributeV2.Operation.Any.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Operation.Any.IsNull() {
					_ = json.Unmarshal([]byte(mappingAttributesItem.MappingAttributeV2.Operation.Any.ValueString()), &anyVar)
				}
				if anyVar != nil {
					operation = shared.OperationNode{
						Any: anyVar,
					}
				}
				origin := new(shared.AttributeOrigin)
				if !mappingAttributesItem.MappingAttributeV2.Origin.IsUnknown() && !mappingAttributesItem.MappingAttributeV2.Origin.IsNull() {
					*origin = shared.AttributeOrigin(mappingAttributesItem.MappingAttributeV2.Origin.ValueString())
				} else {
					origin = nil
				}
				var target string
				target = mappingAttributesItem.MappingAttributeV2.Target.ValueString()

				mappingAttributeV2 := shared.MappingAttributeV2{
					Operation: operation,
					Origin:    origin,
					Target:    target,
				}
				mappingAttributes = append(mappingAttributes, shared.MappingAttributes{
					MappingAttributeV2: &mappingAttributeV2,
				})
			}
			if mappingAttributesItem.MappingAttribute != nil {
				var mappingAttribute shared.MappingAttribute
				var setValueMapper *shared.SetValueMapper
				if mappingAttributesItem.MappingAttribute.SetValueMapper != nil {
					mode := shared.MappingAttributeMode(mappingAttributesItem.MappingAttribute.SetValueMapper.Mode.ValueString())
					var target1 string
					target1 = mappingAttributesItem.MappingAttribute.SetValueMapper.Target.ValueString()

					var value interface{}
					_ = json.Unmarshal([]byte(mappingAttributesItem.MappingAttribute.SetValueMapper.Value.ValueString()), &value)
					setValueMapper = &shared.SetValueMapper{
						Mode:   mode,
						Target: target1,
						Value:  value,
					}
				}
				if setValueMapper != nil {
					mappingAttribute = shared.MappingAttribute{
						SetValueMapper: setValueMapper,
					}
				}
				var copyValueMapper *shared.CopyValueMapper
				if mappingAttributesItem.MappingAttribute.CopyValueMapper != nil {
					mode1 := shared.MappingAttributeMode(mappingAttributesItem.MappingAttribute.CopyValueMapper.Mode.ValueString())
					var source1 string
					source1 = mappingAttributesItem.MappingAttribute.CopyValueMapper.Source.ValueString()

					var target2 string
					target2 = mappingAttributesItem.MappingAttribute.CopyValueMapper.Target.ValueString()

					copyValueMapper = &shared.CopyValueMapper{
						Mode:   mode1,
						Source: source1,
						Target: target2,
					}
				}
				if copyValueMapper != nil {
					mappingAttribute = shared.MappingAttribute{
						CopyValueMapper: copyValueMapper,
					}
				}
				var appendValueMapper *shared.AppendValueMapper
				if mappingAttributesItem.MappingAttribute.AppendValueMapper != nil {
					mode2 := shared.MappingAttributeMode(mappingAttributesItem.MappingAttribute.AppendValueMapper.Mode.ValueString())
					source2 := new(string)
					if !mappingAttributesItem.MappingAttribute.AppendValueMapper.Source.IsUnknown() && !mappingAttributesItem.MappingAttribute.AppendValueMapper.Source.IsNull() {
						*source2 = mappingAttributesItem.MappingAttribute.AppendValueMapper.Source.ValueString()
					} else {
						source2 = nil
					}
					var target3 string
					target3 = mappingAttributesItem.MappingAttribute.AppendValueMapper.Target.ValueString()

					var targetUnique []string = []string{}
					for _, targetUniqueItem := range mappingAttributesItem.MappingAttribute.AppendValueMapper.TargetUnique {
						targetUnique = append(targetUnique, targetUniqueItem.ValueString())
					}
					var valueJSON string
					valueJSON = mappingAttributesItem.MappingAttribute.AppendValueMapper.ValueJSON.ValueString()

					appendValueMapper = &shared.AppendValueMapper{
						Mode:         mode2,
						Source:       source2,
						Target:       target3,
						TargetUnique: targetUnique,
						ValueJSON:    valueJSON,
					}
				}
				if appendValueMapper != nil {
					mappingAttribute = shared.MappingAttribute{
						AppendValueMapper: appendValueMapper,
					}
				}
				mappingAttributes = append(mappingAttributes, shared.MappingAttributes{
					MappingAttribute: &mappingAttribute,
				})
			}
		}
		name := new(string)
		if !targetsItem.Name.IsUnknown() && !targetsItem.Name.IsNull() {
			*name = targetsItem.Name.ValueString()
		} else {
			name = nil
		}
		var relationAttributes []shared.RelationAttribute = []shared.RelationAttribute{}
		for _, relationAttributesItem := range targetsItem.RelationAttributes {
			mode3 := shared.Mode(relationAttributesItem.Mode.ValueString())
			origin1 := new(shared.AttributeOrigin)
			if !relationAttributesItem.Origin.IsUnknown() && !relationAttributesItem.Origin.IsNull() {
				*origin1 = shared.AttributeOrigin(relationAttributesItem.Origin.ValueString())
			} else {
				origin1 = nil
			}
			relatedTo := make(map[string]interface{})
			for relatedToKey, relatedToValue := range relationAttributesItem.RelatedTo {
				var relatedToInst interface{}
				_ = json.Unmarshal([]byte(relatedToValue.ValueString()), &relatedToInst)
				relatedTo[relatedToKey] = relatedToInst
			}
			var sourceFilter *shared.SourceFilter
			if relationAttributesItem.SourceFilter != nil {
				attribute := new(string)
				if !relationAttributesItem.SourceFilter.Attribute.IsUnknown() && !relationAttributesItem.SourceFilter.Attribute.IsNull() {
					*attribute = relationAttributesItem.SourceFilter.Attribute.ValueString()
				} else {
					attribute = nil
				}
				limit := new(int64)
				if !relationAttributesItem.SourceFilter.Limit.IsUnknown() && !relationAttributesItem.SourceFilter.Limit.IsNull() {
					*limit = relationAttributesItem.SourceFilter.Limit.ValueInt64()
				} else {
					limit = nil
				}
				relationTag := new(string)
				if !relationAttributesItem.SourceFilter.RelationTag.IsUnknown() && !relationAttributesItem.SourceFilter.RelationTag.IsNull() {
					*relationTag = relationAttributesItem.SourceFilter.RelationTag.ValueString()
				} else {
					relationTag = nil
				}
				schema := new(string)
				if !relationAttributesItem.SourceFilter.Schema.IsUnknown() && !relationAttributesItem.SourceFilter.Schema.IsNull() {
					*schema = relationAttributesItem.SourceFilter.Schema.ValueString()
				} else {
					schema = nil
				}
				self := new(bool)
				if !relationAttributesItem.SourceFilter.Self.IsUnknown() && !relationAttributesItem.SourceFilter.Self.IsNull() {
					*self = relationAttributesItem.SourceFilter.Self.ValueBool()
				} else {
					self = nil
				}
				tag := new(string)
				if !relationAttributesItem.SourceFilter.Tag.IsUnknown() && !relationAttributesItem.SourceFilter.Tag.IsNull() {
					*tag = relationAttributesItem.SourceFilter.Tag.ValueString()
				} else {
					tag = nil
				}
				sourceFilter = &shared.SourceFilter{
					Attribute:   attribute,
					Limit:       limit,
					RelationTag: relationTag,
					Schema:      schema,
					Self:        self,
					Tag:         tag,
				}
			}
			var target4 string
			target4 = relationAttributesItem.Target.ValueString()

			var targetTags []string = []string{}
			for _, targetTagsItem := range relationAttributesItem.TargetTags {
				targetTags = append(targetTags, targetTagsItem.ValueString())
			}
			targetTagsIncludeSource := new(bool)
			if !relationAttributesItem.TargetTagsIncludeSource.IsUnknown() && !relationAttributesItem.TargetTagsIncludeSource.IsNull() {
				*targetTagsIncludeSource = relationAttributesItem.TargetTagsIncludeSource.ValueBool()
			} else {
				targetTagsIncludeSource = nil
			}
			relationAttributes = append(relationAttributes, shared.RelationAttribute{
				Mode:                    mode3,
				Origin:                  origin1,
				RelatedTo:               relatedTo,
				SourceFilter:            sourceFilter,
				Target:                  target4,
				TargetTags:              targetTags,
				TargetTagsIncludeSource: targetTagsIncludeSource,
			})
		}
		var targetSchema string
		targetSchema = targetsItem.TargetSchema.ValueString()

		var targetUnique1 []string = []string{}
		for _, targetUniqueItem1 := range targetsItem.TargetUnique {
			targetUnique1 = append(targetUnique1, targetUniqueItem1.ValueString())
		}
		targets = append(targets, shared.TargetConfig{
			AllowFailure:              allowFailure,
			ConditionMode:             conditionMode,
			Conditions:                conditions,
			ID:                        id1,
			LinkbackRelationAttribute: linkbackRelationAttribute,
			LinkbackRelationTags:      linkbackRelationTags,
			LoopConfig:                loopConfig,
			MappingAttributes:         mappingAttributes,
			Name:                      name,
			RelationAttributes:        relationAttributes,
			TargetSchema:              targetSchema,
			TargetUnique:              targetUnique1,
		})
	}
	out := shared.MappingConfigV2{
		ID:      id,
		Source:  source,
		Targets: targets,
	}
	return &out
}

func (r *EntityMappingResourceModel) RefreshFromSharedMappingConfigV2(resp *shared.MappingConfigV2) {
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
