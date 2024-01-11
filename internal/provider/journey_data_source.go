// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/epilot-dev/terraform-provider-epilot-journey/internal/sdk"
	"github.com/epilot-dev/terraform-provider-epilot-journey/internal/sdk/pkg/models/operations"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &JourneyDataSource{}
var _ datasource.DataSourceWithConfigure = &JourneyDataSource{}

func NewJourneyDataSource() datasource.DataSource {
	return &JourneyDataSource{}
}

// JourneyDataSource is the data source implementation.
type JourneyDataSource struct {
	client *sdk.SDK
}

// JourneyDataSourceModel describes the data model.
type JourneyDataSourceModel struct {
	AdditionalProperties types.String `tfsdk:"additional_properties"`
	BrandID              types.String `tfsdk:"brand_id"`
	CreatedAt            types.String `tfsdk:"created_at"`
	CreatedBy            types.String `tfsdk:"created_by"`
	Design               *Design      `tfsdk:"design"`
	JourneyID            types.String `tfsdk:"journey_id"`
	LastModifiedAt       types.String `tfsdk:"last_modified_at"`
	Logics               []Logics     `tfsdk:"logics"`
	Name                 types.String `tfsdk:"name"`
	OrganizationID       types.String `tfsdk:"organization_id"`
	OrgID                types.String `tfsdk:"org_id"`
	Revisions            types.Number `tfsdk:"revisions"`
	Rules                []Rules      `tfsdk:"rules"`
	Settings             *Settings    `tfsdk:"settings"`
	Source               types.String `tfsdk:"source"`
	Steps                []Steps      `tfsdk:"steps"`
	Version              types.Number `tfsdk:"version"`
}

// Metadata returns the data source type name.
func (r *JourneyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_journey"
}

// Schema defines the schema for the data source.
func (r *JourneyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Journey DataSource",

		Attributes: map[string]schema.Attribute{
			"additional_properties": schema.StringAttribute{
				Computed:    true,
				Description: `Parsed as JSON.`,
			},
			"brand_id": schema.StringAttribute{
				Computed: true,
			},
			"created_at": schema.StringAttribute{
				Computed: true,
			},
			"created_by": schema.StringAttribute{
				Computed: true,
			},
			"design": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"logo_url": schema.StringAttribute{
						Computed: true,
					},
					"theme": schema.MapAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
				},
			},
			"journey_id": schema.StringAttribute{
				Required:    true,
				Description: `Journey ID`,
			},
			"last_modified_at": schema.StringAttribute{
				Computed: true,
			},
			"logics": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"actions": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
						"auto_generated_id": schema.StringAttribute{
							Computed: true,
						},
						"conditions": schema.ListAttribute{
							Computed:    true,
							ElementType: types.StringType,
						},
					},
				},
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"organization_id": schema.StringAttribute{
				Computed: true,
			},
			"org_id": schema.StringAttribute{
				Optional:    true,
				Description: `Organization ID`,
			},
			"revisions": schema.NumberAttribute{
				Computed: true,
			},
			"rules": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							Computed: true,
						},
						"source_type": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["journey", "step", "block"]`,
						},
						"target": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: `must be one of ["inject", "injectWithKey"]`,
						},
					},
				},
			},
			"settings": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"address_suggestions_file_url": schema.StringAttribute{
						Computed: true,
					},
					"description": schema.StringAttribute{
						Computed: true,
					},
					"design_id": schema.StringAttribute{
						Computed: true,
					},
					"embed_options": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"button": schema.SingleNestedAttribute{
								Computed: true,
								Attributes: map[string]schema.Attribute{
									"align": schema.StringAttribute{
										Computed:    true,
										Description: `must be one of ["left", "center", "right"]`,
									},
									"text": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							"lang": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["de", "en", "fr"]`,
							},
							"mode": schema.StringAttribute{
								Computed:    true,
								Description: `must be one of ["full-screen", "inline"]`,
							},
							"scroll_to_top": schema.BoolAttribute{
								Computed: true,
							},
							"top_bar": schema.BoolAttribute{
								Computed: true,
							},
							"width": schema.StringAttribute{
								Computed: true,
							},
						},
					},
					"entity_id": schema.StringAttribute{
						Computed: true,
					},
					"entity_tags": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"file_purposes": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"mappings_automation_id": schema.StringAttribute{
						Computed: true,
					},
					"organization_settings": schema.MapAttribute{
						Computed:    true,
						ElementType: types.BoolType,
					},
					"public_token": schema.StringAttribute{
						Computed: true,
					},
					"runtime_entities": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
					},
					"safe_mode_automation": schema.BoolAttribute{
						Computed: true,
					},
					"targeted_customer": schema.StringAttribute{
						Computed: true,
					},
					"template_id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"source": schema.StringAttribute{
				Optional:    true,
				Description: `What source ID. Journey or Entity ID`,
			},
			"steps": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"hide_next_button": schema.BoolAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"schema": schema.StringAttribute{
							Computed:    true,
							Description: `Parsed as JSON.`,
						},
						"show_step_name": schema.BoolAttribute{
							Computed: true,
						},
						"show_stepper": schema.BoolAttribute{
							Computed: true,
						},
						"show_stepper_labels": schema.BoolAttribute{
							Computed: true,
						},
						"show_step_subtitle": schema.BoolAttribute{
							Computed: true,
						},
						"step_id": schema.StringAttribute{
							Computed: true,
						},
						"sub_title": schema.StringAttribute{
							Computed: true,
						},
						"title": schema.StringAttribute{
							Computed: true,
						},
						"uischema": schema.StringAttribute{
							Computed:    true,
							Description: `Parsed as JSON.`,
						},
					},
				},
			},
			"version": schema.NumberAttribute{
				Computed: true,
			},
		},
	}
}

func (r *JourneyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (r *JourneyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *JourneyDataSourceModel
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

	id := data.JourneyID.ValueString()
	orgID := new(string)
	if !data.OrgID.IsUnknown() && !data.OrgID.IsNull() {
		*orgID = data.OrgID.ValueString()
	} else {
		orgID = nil
	}
	source := new(string)
	if !data.Source.IsUnknown() && !data.Source.IsNull() {
		*source = data.Source.ValueString()
	} else {
		source = nil
	}
	request := operations.GetJourneyRequest{
		ID:     id,
		OrgID:  orgID,
		Source: source,
	}
	res, err := r.client.Journeys.GetJourney(ctx, request)
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
	if res.Journey == nil {
		resp.Diagnostics.AddError("unexpected response from API. No response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedJourney(res.Journey)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}