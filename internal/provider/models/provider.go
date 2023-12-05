package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type ProviderModel struct {
	Ak     types.String `tfsdk:"ak"`
	Sk     types.String `tfsdk:"sk"`
	Region types.String `tfsdk:"region"`
}
