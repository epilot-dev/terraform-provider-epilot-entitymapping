// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	tfTypes "github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/provider/types"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-entitymapping/internal/sdk/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &EntityMappingDataSource{}
var _ datasource.DataSourceWithConfigure = &EntityMappingDataSource{}

func NewEntityMappingDataSource() datasource.DataSource {
	return &EntityMappingDataSource{}
}

// EntityMappingDataSource is the data source implementation.
type EntityMappingDataSource struct {
	client *sdk.SDK
}

// EntityMappingDataSourceModel describes the data model.
type EntityMappingDataSourceModel struct {
	ID      types.String           `tfsdk:"id"`
	Source  tfTypes.SourceConfig   `tfsdk:"source"`
	Targets []tfTypes.TargetConfig `tfsdk:"targets"`
}

// Metadata returns the data source type name.
func (r *EntityMappingDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_entity_mapping"
}

// Schema defines the schema for the data source.
func (r *EntityMappingDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "EntityMapping DataSource",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Required:    true,
				Description: `Mapping Config Id`,
			},
			"source": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"config": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"entity_ref": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"entity_id": schema.StringAttribute{
										Computed:    true,
										Description: `id of the source entity to be mapped`,
									},
									"entity_schema": schema.StringAttribute{
										Computed:    true,
										Description: `schema of the source entity`,
									},
								},
							},
							"journey_ref": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"journey_id": schema.StringAttribute{
										Computed: true,
									},
								},
							},
						},
					},
					"type": schema.StringAttribute{
						Computed:    true,
						Description: `must be one of ["journey", "entity"]`,
					},
				},
			},
			"targets": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"allow_failure": schema.BoolAttribute{
							Computed:    true,
							Description: `Pass it as true, when you don't want failures to interrupt the mapping process.`,
						},
						"condition_mode": schema.StringAttribute{
							Computed:    true,
							Description: `Parsed as JSON.`,
						},
						"conditions": schema.StringAttribute{
							Computed:    true,
							Description: `Parsed as JSON.`,
						},
						"id": schema.StringAttribute{
							Computed:    true,
							Description: `Identifier for target configuration. Useful for later usages when trying to identify which target config to map to.`,
						},
						"linkback_relation_attribute": schema.StringAttribute{
							Computed: true,
							MarkdownDescription: `Relation attribute on the main entity where the target entity will be linked. Set to false to disable linkback` + "\n" +
								``,
						},
						"linkback_relation_tags": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `Relation tags (labels) to include in main entity linkback relation attribute`,
						},
						"loop_config": schema.SingleNestedAttribute{
							Computed: true,
							Attributes: map[string]schema.Attribute{
								"length": schema.NumberAttribute{
									Computed:    true,
									Description: `a hard limit of how many times the loop is allowed to run.`,
								},
								"source_path": schema.StringAttribute{
									Computed:    true,
									Description: `path to the array from the entity payload`,
								},
							},
							Description: `contains config in case of running in loop mode`,
						},
						"mapping_attributes": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"mapping_attribute": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"append_value_mapper": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"mode": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `- copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property.` + "\n" +
															`` + "\n" +
															`must be one of ["copy_if_exists", "append_if_exists", "set_value"]`,
													},
													"source": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `JSON source path for the value to be extracted from the main entity. Eg: steps[1].['Product Info'].price` + "\n" +
															``,
													},
													"target": schema.StringAttribute{
														Computed:    true,
														Description: `JSON like target path for the attribute. Eg. last_name`,
													},
													"target_unique": schema.ListAttribute{
														Computed:    true,
														ElementType: types.StringType,
														MarkdownDescription: `Array of keys which should be used when checking for uniqueness. Eg: [country, city, postal_code]` + "\n" +
															``,
													},
													"value_json": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `To be provided only when mapping json objects into a target attribute. Eg array of addresses.` + "\n" +
															``,
													},
												},
											},
											"copy_value_mapper": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"mode": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `- copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property.` + "\n" +
															`` + "\n" +
															`must be one of ["copy_if_exists", "append_if_exists", "set_value"]`,
													},
													"source": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `JSON source path for the value to be extracted from the main entity. Eg: steps[1].['Product Info'].price` + "\n" +
															``,
													},
													"target": schema.StringAttribute{
														Computed:    true,
														Description: `JSON like target path for the attribute. Eg. last_name`,
													},
												},
											},
											"set_value_mapper": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"mode": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `- copy_if_exists - it replaces the target attribute with the source value - append_if_exists - it currently replaces target attribute with array like values. Useful when you have multiple values to be added into one attribute. - set_value - it sets a value to a predefined value. Must be used together with value property.` + "\n" +
															`` + "\n" +
															`must be one of ["copy_if_exists", "append_if_exists", "set_value"]`,
													},
													"target": schema.StringAttribute{
														Computed:    true,
														Description: `JSON like target path for the attribute. Eg. last_name`,
													},
													"value": schema.StringAttribute{
														Computed: true,
														MarkdownDescription: `Any value to be set: string, number, string[], number[], JSON object, etc. It will override existing values, if any.` + "\n" +
															`` + "\n" +
															`Parsed as JSON.`,
													},
												},
											},
										},
									},
									"mapping_attribute_v2": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"operation": schema.SingleNestedAttribute{
												Computed: true,
												Attributes: map[string]schema.Attribute{
													"any": schema.StringAttribute{
														Computed:    true,
														Description: `Parsed as JSON.`,
													},
													"operation_object_node": schema.SingleNestedAttribute{
														Computed: true,
														Attributes: map[string]schema.Attribute{
															"append": schema.ListAttribute{
																Computed:    true,
																ElementType: types.StringType,
																Description: `Append to array`,
															},
															"copy": schema.StringAttribute{
																Computed:    true,
																Description: `Copy JSONPath value from source entity context`,
															},
															"random": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"one": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"type": schema.StringAttribute{
																				Computed:    true,
																				Description: `must be one of ["uuid", "nanoid"]`,
																			},
																		},
																	},
																	"two": schema.SingleNestedAttribute{
																		Computed: true,
																		Attributes: map[string]schema.Attribute{
																			"max": schema.NumberAttribute{
																				Computed: true,
																			},
																			"min": schema.NumberAttribute{
																				Computed: true,
																			},
																			"type": schema.StringAttribute{
																				Computed:    true,
																				Description: `must be one of ["number"]`,
																			},
																		},
																	},
																},
															},
															"set": schema.StringAttribute{
																Computed:    true,
																Description: `Parsed as JSON.`,
															},
															"template": schema.StringAttribute{
																Computed:    true,
																Description: `Define handlebars template to output a string`,
															},
															"uniq": schema.SingleNestedAttribute{
																Computed: true,
																Attributes: map[string]schema.Attribute{
																	"boolean": schema.BoolAttribute{
																		Computed: true,
																	},
																	"array_of_str": schema.ListAttribute{
																		Computed:    true,
																		ElementType: types.StringType,
																	},
																},
																Description: `Unique array`,
															},
															"additional_properties": schema.StringAttribute{
																Computed:    true,
																Description: `Parsed as JSON.`,
															},
														},
													},
												},
												Description: `Mapping operation nodes are either primitive values or operation node objects`,
											},
											"origin": schema.StringAttribute{
												Computed:    true,
												Description: `Origin of an attribute. must be one of ["system_recommendation", "user_manually", "entity_updating_system_recommendation"]`,
											},
											"target": schema.StringAttribute{
												Computed:    true,
												Description: `Target JSON path for the attribute to set`,
											},
										},
									},
								},
							},
							Description: `Attribute mappings`,
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: `A name for this configuration`,
						},
						"relation_attributes": schema.ListNestedAttribute{
							Computed: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"mode": schema.StringAttribute{
										Computed:    true,
										Description: `must be one of ["append", "prepend", "set"]`,
									},
									"origin": schema.StringAttribute{
										Computed:    true,
										Description: `Origin of an attribute. must be one of ["system_recommendation", "user_manually", "entity_updating_system_recommendation"]`,
									},
									"related_to": schema.MapAttribute{
										Computed:    true,
										ElementType: types.StringType,
									},
									"source_filter": schema.SingleNestedAttribute{
										Computed: true,
										Attributes: map[string]schema.Attribute{
											"attribute": schema.StringAttribute{
												Computed:    true,
												Description: `Filter by a specific relation attribute on the main entity`,
											},
											"limit": schema.Int64Attribute{
												Computed:    true,
												Description: `Limit relations to maximum number (default, all matched relations)`,
											},
											"relation_tag": schema.StringAttribute{
												Computed:    true,
												Description: `Filter by relation tag (label) on the main entity`,
											},
											"schema": schema.StringAttribute{
												Computed:    true,
												Description: `Filter by specific schema`,
											},
											"self": schema.BoolAttribute{
												Computed:    true,
												Description: `Picks main entity as relation (overrides other filters)`,
											},
											"tag": schema.StringAttribute{
												Computed:    true,
												Description: `Filter by a specific tag on the related entity`,
											},
										},
										Description: `A filter to identify which source entities to pick as relations from main entity`,
									},
									"target": schema.StringAttribute{
										Computed:    true,
										Description: `Target attribute to store the relation in`,
									},
									"target_tags": schema.ListAttribute{
										Computed:    true,
										ElementType: types.StringType,
										Description: `Relation tags (labels) to set for the stored relations`,
									},
									"target_tags_include_source": schema.BoolAttribute{
										Computed:    true,
										Description: `Include all relation tags (labels) present on the main entity relation`,
									},
								},
							},
							Description: `Relation mappings`,
						},
						"target_schema": schema.StringAttribute{
							Computed:    true,
							Description: `Schema of target entity`,
						},
						"target_unique": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
							Description: `Unique key for target entity (see upsertEntity of Entity API)`,
						},
					},
				},
			},
		},
	}
}

func (r *EntityMappingDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.SDK)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.SDK, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *EntityMappingDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *EntityMappingDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
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

	id := data.ID.ValueString()
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