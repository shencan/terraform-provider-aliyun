package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type EssTagModels struct {
	EssId types.String `tfsdk:"ess_id"`
	Tags  types.Map    `tfsdk:"tags"`
}
