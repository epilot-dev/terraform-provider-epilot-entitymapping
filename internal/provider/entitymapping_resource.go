// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	speakeasy_boolplanmodifier "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/planmodifiers/boolplanmodifier"
	speakeasy_int64planmodifier "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/planmodifiers/int64planmodifier"
	speakeasy_listplanmodifier "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/planmodifiers/listplanmodifier"
	speakeasy_numberplanmodifier "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/planmodifiers/numberplanmodifier"
	speakeasy_objectplanmodifier "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/planmodifiers/objectplanmodifier"
	speakeasy_stringplanmodifier "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/planmodifiers/stringplanmodifier"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/models/operations"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/validators"
	speakeasy_listvalidators "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/validators/listvalidators"
	speakeasy_objectvalidators "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/validators/objectvalidators"
	speakeasy_stringvalidators "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/validators/stringvalidators"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/numberplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &EntityMappingResource{}
var _ resource.ResourceWithImportState = &EntityMappingResource{}

func NewEntityMappingResource() resource.Resource {
	return &EntityMappingResource{}
}

// EntityMappingResource defines the resource implementation.
type EntityMappingResource struct {
	client *sdk.SDK
}

// EntityMappingResourceModel describes the resource data model.
type EntityMappingResourceModel struct {
	ID      types.String           `tfsdk:"id"`
	Source  tfTypes.SourceConfig   `tfsdk:"source"`
	Targets []tfTypes.TargetConfig `tfsdk:"targets"`
	Version types.Int64            `tfsdk:"version"`
}

func (r *EntityMappingResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_entity_mapping"
}

func (r *EntityMappingResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "EntityMapping Resource",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
				},
				Description: `Mapping Config Id. Requires replacement if changed.`,
			},
			"source": schema.SingleNestedAttribute{
				Required: true,
				PlanModifiers: []planmodifier.Object{
					objectplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
				},
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.Object{
							objectplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
						},
						Attributes: map[string]schema.Attribute{
							"entity_ref": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Attributes: map[string]schema.Attribute{
									"entity_id": schema.StringAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Description: `id of the source entity to be mapped. Not Null; Requires replacement if changed.`,
										Validators: []validator.String{
											speakeasy_stringvalidators.NotNull(),
										},
									},
									"entity_schema": schema.StringAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Description: `schema of the source entity. Requires replacement if changed.`,
									},
								},
								Description: `Requires replacement if changed.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("journey_ref"),
									}...),
								},
							},
							"journey_ref": schema.SingleNestedAttribute{
								Computed: true,
								Optional: true,
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Attributes: map[string]schema.Attribute{
									"journey_id": schema.StringAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Description: `Requires replacement if changed.`,
									},
								},
								Description: `Requires replacement if changed.`,
								Validators: []validator.Object{
									objectvalidator.ConflictsWith(path.Expressions{
										path.MatchRelative().AtParent().AtName("entity_ref"),
									}...),
								},
							},
						},
						Description: `Requires replacement if changed.`,
					},
					"type": schema.StringAttribute{
						Computed: true,
						Optional: true,
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.RequiresReplaceIfConfigured(),
							speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
						},
						Description: `must be one of ["journey", "entity"]; Requires replacement if changed.`,
						Validators: []validator.String{
							stringvalidator.OneOf(
								"journey",
								"entity",
							),
						},
					},
				},
				Description: `Requires replacement if changed.`,
			},
			"targets": schema.ListNestedAttribute{
				Required: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.RequiresReplaceIfConfigured(),
					speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
				},
				NestedObject: schema.NestedAttributeObject{
					Validators: []validator.Object{
						speakeasy_objectvalidators.NotNull(),
					},
					PlanModifiers: []planmodifier.Object{
						objectplanmodifier.RequiresReplaceIfConfigured(),
						speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
					},
					Attributes: map[string]schema.Attribute{
						"allow_failure": schema.BoolAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Bool{
								boolplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
							},
							Description: `Pass it as true, when you don't want failures to interrupt the mapping process. Requires replacement if changed.`,
						},
						"condition_mode": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Description: `Requires replacement if changed.; Parsed as JSON.`,
							Validators: []validator.String{
								validators.IsValidJSON(),
							},
						},
						"conditions": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Description: `Requires replacement if changed.; Parsed as JSON.`,
							Validators: []validator.String{
								validators.IsValidJSON(),
							},
						},
						"id": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Description: `Identifier for target configuration. Useful for later usages when trying to identify which target config to map to. Requires replacement if changed.`,
						},
						"linkback_relation_attribute": schema.StringAttribute{
							Computed: true,
							Optional: true,
							Default:  stringdefault.StaticString("mapped_entities"),
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Description: `Relation attribute on the main entity where the target entity will be linked. Set to false to disable linkback. Default: "mapped_entities"; Requires replacement if changed.`,
						},
						"linkback_relation_tags": schema.ListAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
							},
							ElementType: types.StringType,
							Description: `Relation tags (labels) to include in main entity linkback relation attribute. Requires replacement if changed.`,
						},
						"loop_config": schema.SingleNestedAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.Object{
								objectplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
							},
							Attributes: map[string]schema.Attribute{
								"length": schema.NumberAttribute{
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.Number{
										numberplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_numberplanmodifier.SuppressDiff(speakeasy_numberplanmodifier.ExplicitSuppress),
									},
									Description: `a hard limit of how many times the loop is allowed to run. Requires replacement if changed.`,
								},
								"source_path": schema.StringAttribute{
									Computed: true,
									Optional: true,
									PlanModifiers: []planmodifier.String{
										stringplanmodifier.RequiresReplaceIfConfigured(),
										speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
									},
									Description: `path to the array from the entity payload. Requires replacement if changed.`,
								},
							},
							Description: `contains config in case of running in loop mode. Requires replacement if changed.`,
						},
						"mapping_attributes": schema.ListNestedAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
							},
							NestedObject: schema.NestedAttributeObject{
								Validators: []validator.Object{
									speakeasy_objectvalidators.NotNull(),
								},
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Attributes: map[string]schema.Attribute{
									"any": schema.StringAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Description: `Requires replacement if changed.; Parsed as JSON.`,
										Validators: []validator.String{
											stringvalidator.ConflictsWith(path.Expressions{
												path.MatchRelative().AtParent().AtName("mapping_attribute"),
											}...),
											validators.IsValidJSON(),
										},
									},
									"mapping_attribute": schema.SingleNestedAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
										},
										Attributes: map[string]schema.Attribute{
											"append_value_mapper": schema.SingleNestedAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.RequiresReplaceIfConfigured(),
													speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
												},
												Attributes: map[string]schema.Attribute{
													"mode": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `- copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property. Not Null; must be one of ["copy_if_exists", "append_if_exists", "set_value"]; Requires replacement if changed.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
															stringvalidator.OneOf(
																"copy_if_exists",
																"append_if_exists",
																"set_value",
															),
														},
													},
													"source": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `JSON source path for the value to be extracted from the main entity. Eg: steps[1].['Product Info'].price. Requires replacement if changed.`,
													},
													"target": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `JSON like target path for the attribute. Eg. last_name. Not Null; Requires replacement if changed.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
														},
													},
													"target_unique": schema.ListAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.List{
															listplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
														},
														ElementType: types.StringType,
														Description: `Array of keys which should be used when checking for uniqueness. Eg: [country, city, postal_code]. Requires replacement if changed.`,
													},
													"value_json": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `To be provided only when mapping json objects into a target attribute. Eg array of addresses. Not Null; Requires replacement if changed.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
														},
													},
												},
												Description: `Requires replacement if changed.`,
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("copy_value_mapper"),
														path.MatchRelative().AtParent().AtName("set_value_mapper"),
													}...),
												},
											},
											"copy_value_mapper": schema.SingleNestedAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.RequiresReplaceIfConfigured(),
													speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
												},
												Attributes: map[string]schema.Attribute{
													"mode": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `- copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property. Not Null; must be one of ["copy_if_exists", "append_if_exists", "set_value"]; Requires replacement if changed.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
															stringvalidator.OneOf(
																"copy_if_exists",
																"append_if_exists",
																"set_value",
															),
														},
													},
													"source": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `JSON source path for the value to be extracted from the main entity. Eg: steps[1].['Product Info'].price. Not Null; Requires replacement if changed.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
														},
													},
													"target": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `JSON like target path for the attribute. Eg. last_name. Not Null; Requires replacement if changed.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
														},
													},
												},
												Description: `Requires replacement if changed.`,
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("append_value_mapper"),
														path.MatchRelative().AtParent().AtName("set_value_mapper"),
													}...),
												},
											},
											"set_value_mapper": schema.SingleNestedAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.Object{
													objectplanmodifier.RequiresReplaceIfConfigured(),
													speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
												},
												Attributes: map[string]schema.Attribute{
													"mode": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `- copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property. Not Null; must be one of ["copy_if_exists", "append_if_exists", "set_value"]; Requires replacement if changed.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
															stringvalidator.OneOf(
																"copy_if_exists",
																"append_if_exists",
																"set_value",
															),
														},
													},
													"target": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `JSON like target path for the attribute. Eg. last_name. Not Null; Requires replacement if changed.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
														},
													},
													"value": schema.StringAttribute{
														Computed: true,
														Optional: true,
														PlanModifiers: []planmodifier.String{
															stringplanmodifier.RequiresReplaceIfConfigured(),
															speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
														},
														Description: `Any value to be set: string, number, string[], number[], JSON object, etc. It will override existing values, if any. Not Null; Requires replacement if changed.; Parsed as JSON.`,
														Validators: []validator.String{
															speakeasy_stringvalidators.NotNull(),
															validators.IsValidJSON(),
														},
													},
												},
												Description: `Requires replacement if changed.`,
												Validators: []validator.Object{
													objectvalidator.ConflictsWith(path.Expressions{
														path.MatchRelative().AtParent().AtName("append_value_mapper"),
														path.MatchRelative().AtParent().AtName("copy_value_mapper"),
													}...),
												},
											},
										},
										Description: `Requires replacement if changed.`,
										Validators: []validator.Object{
											objectvalidator.ConflictsWith(path.Expressions{
												path.MatchRelative().AtParent().AtName("any"),
											}...),
										},
									},
								},
							},
							Description: `Attribute mappings. Not Null; Requires replacement if changed.`,
							Validators: []validator.List{
								speakeasy_listvalidators.NotNull(),
							},
						},
						"name": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Description: `A name for this configuration. Requires replacement if changed.`,
						},
						"relation_attributes": schema.ListNestedAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
							},
							NestedObject: schema.NestedAttributeObject{
								Validators: []validator.Object{
									speakeasy_objectvalidators.NotNull(),
								},
								PlanModifiers: []planmodifier.Object{
									objectplanmodifier.RequiresReplaceIfConfigured(),
									speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
								},
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Description: `Not Null; must be one of ["append", "prepend", "set"]; Requires replacement if changed.`,
										Validators: []validator.String{
											speakeasy_stringvalidators.NotNull(),
											stringvalidator.OneOf(
												"append",
												"prepend",
												"set",
											),
										},
									},
									"origin": schema.StringAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Description: `Origin of an attribute. must be one of ["system_recommendation", "user_manually", "entity_updating_system_recommendation"]; Requires replacement if changed.`,
										Validators: []validator.String{
											stringvalidator.OneOf(
												"system_recommendation",
												"user_manually",
												"entity_updating_system_recommendation",
											),
										},
									},
									"source_filter": schema.SingleNestedAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.Object{
											objectplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_objectplanmodifier.SuppressDiff(speakeasy_objectplanmodifier.ExplicitSuppress),
										},
										Attributes: map[string]schema.Attribute{
											"attribute": schema.StringAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplaceIfConfigured(),
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
												},
												Description: `Filter by a specific relation attribute on the main entity. Requires replacement if changed.`,
											},
											"limit": schema.Int64Attribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.Int64{
													int64planmodifier.RequiresReplaceIfConfigured(),
													speakeasy_int64planmodifier.SuppressDiff(speakeasy_int64planmodifier.ExplicitSuppress),
												},
												Description: `Limit relations to maximum number (default, all matched relations). Requires replacement if changed.`,
											},
											"relation_tag": schema.StringAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplaceIfConfigured(),
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
												},
												Description: `Filter by relation tag (label) on the main entity. Requires replacement if changed.`,
											},
											"schema": schema.StringAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplaceIfConfigured(),
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
												},
												Description: `Filter by specific schema. Requires replacement if changed.`,
											},
											"self": schema.BoolAttribute{
												Computed: true,
												Optional: true,
												Default:  booldefault.StaticBool(false),
												PlanModifiers: []planmodifier.Bool{
													boolplanmodifier.RequiresReplaceIfConfigured(),
													speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
												},
												Description: `Picks main entity as relation (overrides other filters). Default: false; Requires replacement if changed.`,
											},
											"tag": schema.StringAttribute{
												Computed: true,
												Optional: true,
												PlanModifiers: []planmodifier.String{
													stringplanmodifier.RequiresReplaceIfConfigured(),
													speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
												},
												Description: `Filter by a specific tag on the related entity. Requires replacement if changed.`,
											},
										},
										Description: `A filter to identify which source entities to pick as relations from main entity. Requires replacement if changed.`,
									},
									"target": schema.StringAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.String{
											stringplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
										},
										Description: `Target attribute to store the relation in. Not Null; Requires replacement if changed.`,
										Validators: []validator.String{
											speakeasy_stringvalidators.NotNull(),
										},
									},
									"target_tags": schema.ListAttribute{
										Computed: true,
										Optional: true,
										PlanModifiers: []planmodifier.List{
											listplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
										},
										ElementType: types.StringType,
										Description: `Relation tags (labels) to set for the stored relations. Requires replacement if changed.`,
									},
									"target_tags_include_source": schema.BoolAttribute{
										Computed: true,
										Optional: true,
										Default:  booldefault.StaticBool(false),
										PlanModifiers: []planmodifier.Bool{
											boolplanmodifier.RequiresReplaceIfConfigured(),
											speakeasy_boolplanmodifier.SuppressDiff(speakeasy_boolplanmodifier.ExplicitSuppress),
										},
										Description: `Include all relation tags (labels) present on the main entity relation. Default: false; Requires replacement if changed.`,
									},
								},
							},
							Description: `Relation mappings. Requires replacement if changed.`,
						},
						"target_schema": schema.StringAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.String{
								stringplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_stringplanmodifier.SuppressDiff(speakeasy_stringplanmodifier.ExplicitSuppress),
							},
							Description: `Schema of target entity. Not Null; Requires replacement if changed.`,
							Validators: []validator.String{
								speakeasy_stringvalidators.NotNull(),
							},
						},
						"target_unique": schema.ListAttribute{
							Computed: true,
							Optional: true,
							PlanModifiers: []planmodifier.List{
								listplanmodifier.RequiresReplaceIfConfigured(),
								speakeasy_listplanmodifier.SuppressDiff(speakeasy_listplanmodifier.ExplicitSuppress),
							},
							ElementType: types.StringType,
							Description: `Unique key for target entity (see upsertEntity of Entity API). Requires replacement if changed.`,
						},
					},
				},
				Description: `Requires replacement if changed.`,
			},
			"version": schema.Int64Attribute{
				Computed: true,
			},
		},
	}
}

func (r *EntityMappingResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *EntityMappingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data *EntityMappingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(plan.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	mappingConfigV2 := data.ToSharedMappingConfigV2Input()
	var id string
	id = data.ID.ValueString()

	request := operations.PutMappingConfigRequest{
		MappingConfigV2: mappingConfigV2,
		ID:              id,
	}
	res, err := r.client.Mappings.PutMappingConfig(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.MappingConfigV2 != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMappingConfigV2(res.MappingConfigV2)
	refreshPlan(ctx, plan, &data, resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EntityMappingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *EntityMappingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var id string
	id = data.ID.ValueString()

	request := operations.GetMappingConfigRequest{
		ID: id,
	}
	res, err := r.client.Mappings.GetMappingConfig(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.MappingConfigV2 != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedMappingConfigV2(res.MappingConfigV2)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EntityMappingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data *EntityMappingResourceModel
	var plan types.Object

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	merge(ctx, req, resp, &data)
	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; all attributes marked as RequiresReplace

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *EntityMappingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data *EntityMappingResourceModel
	var item types.Object

	resp.Diagnostics.Append(req.State.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Not Implemented; entity does not have a configured DELETE operation
}

func (r *EntityMappingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), req.ID)...)
}
