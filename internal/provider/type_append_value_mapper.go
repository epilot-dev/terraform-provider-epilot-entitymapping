// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type AppendValueMapper struct {
	Mode         types.String   `tfsdk:"mode"`
	Source       types.String   `tfsdk:"source"`
	Target       types.String   `tfsdk:"target"`
	TargetUnique []types.String `tfsdk:"target_unique"`
	ValueJSON    types.String   `tfsdk:"value_json"`
}