// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type SourceFilter struct {
	Attribute   types.String `tfsdk:"attribute"`
	Limit       types.Int64  `tfsdk:"limit"`
	RelationTag types.String `tfsdk:"relation_tag"`
	Schema      types.String `tfsdk:"schema"`
	Self        types.Bool   `tfsdk:"self"`
	Tag         types.String `tfsdk:"tag"`
}