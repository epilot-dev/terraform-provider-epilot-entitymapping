// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type Uniq struct {
	Boolean    types.Bool     `tfsdk:"boolean" tfPlanOnly:"true"`
	ArrayOfStr []types.String `tfsdk:"array_of_str" tfPlanOnly:"true"`
}
