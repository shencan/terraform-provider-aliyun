package pkg

import (
	"context"
	"fmt"
	ess20220222 "github.com/alibabacloud-go/ess-20220222/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/shencan/terraform-provider-aliyun/pkg/models"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = &EssTagResource{}
var _ resource.ResourceWithConfigure = &EssTagResource{}

type EssTagResource struct {
	Client *ess20220222.Client
	Region string
}

func NewEssTagResource() resource.Resource {
	return &EssTagResource{}
}

func (esstag *EssTagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ess_tag"
}

func (esstag *EssTagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "`serviceaccount data source`",
		Attributes: map[string]schema.Attribute{
			"ess_id": schema.StringAttribute{
				Required: true,
			},
			"tags": schema.MapAttribute{
				ElementType: types.StringType,
				Required:    true,
			},
		},
	}
}

func (esstag *EssTagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan models.EssTagModels
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if plan.EssId.String() == "" {
		resp.Diagnostics.AddError("Ess_id is ", "empty")
	}

	if plan.Tags.IsNull() {
		resp.Diagnostics.AddError("tags is ", "empty")
	}

	essId := plan.EssId.ValueString()
	input := &ess20220222.TagResourcesRequest{
		RegionId: tea.String(esstag.Region),
		ResourceIds: []*string{
			&essId,
		},
		ResourceType: tea.String("scalinggroup"),
	}

	for k, v := range plan.Tags.Elements() {
		input.Tags = append(input.Tags, &ess20220222.TagResourcesRequestTags{
			Value: tea.String(v.String()),
			Key:   tea.String(k),
		})
	}

	tflog.Info(ctx, "Create Configured---->", map[string]any{
		"region": esstag.Region,
		"essId":  plan.EssId.ValueString(),
		"tags":   input.Tags,
	})

	_, err := esstag.Client.TagResources(input)
	if err != nil {
		resp.Diagnostics.AddError("could not create ess tag[create]", err.Error())
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (esstag *EssTagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state models.EssTagModels
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	essId := state.EssId.ValueString()
	response, err := esstag.Client.DescribeScalingGroups(&ess20220222.DescribeScalingGroupsRequest{
		RegionId: tea.String(esstag.Region),
		ScalingGroupIds: []*string{
			&essId,
		},
	})
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading DescribeScalingGroups",
			"Could not read DescribeScalingGroups "+state.EssId.ValueString()+"——>"+": "+err.Error(),
		)
	}

	if response.Body == nil {
		resp.Diagnostics.AddError("error reading ess DescribeScalingGroups", "response.Body is nil")
	}
	asgTag := make(map[string]string)
	for _, asg := range response.Body.ScalingGroups {
		for _, tag := range asg.Tags {
			asgTag[tea.StringValue(tag.TagKey)] = tea.StringValue(tag.TagValue)
		}
	}

	stateTag := make(map[string]attr.Value)
	for k, v := range state.Tags.Elements() {
		if val, ok := asgTag[k]; ok {
			stateTag[k] = types.StringValue(val)
		}
		stateTag[k] = v
	}

	value, _ := types.MapValue(types.StringType, stateTag)
	state.Tags = value

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		resp.Diagnostics.AddError(
			"Error Reading Ess Tag for state files",
			"Could not update Ess Tag "+state.EssId.ValueString(),
		)
		return
	}

}

func (esstag *EssTagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan models.EssTagModels
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	essId := plan.EssId.ValueString()
	input := &ess20220222.TagResourcesRequest{
		RegionId: tea.String(esstag.Region),
		ResourceIds: []*string{
			&essId,
		},
		ResourceType: tea.String("scalinggroup"),
	}

	for k, v := range plan.Tags.Elements() {
		input.Tags = append(input.Tags, &ess20220222.TagResourcesRequestTags{
			Value: tea.String(v.String()),
			Key:   tea.String(k),
		})
	}

	tflog.Info(ctx, "Update Configured---->", map[string]any{
		"region": esstag.Region,
		"essId":  plan.EssId.ValueString(),
		"tags":   input.Tags,
	})

	_, err := esstag.Client.TagResources(input)
	if err != nil {
		resp.Diagnostics.AddError("could not create ess tag[update]", err.Error())
		return
	}

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (esstag *EssTagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state models.EssTagModels
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	essId := state.EssId.ValueString()
	input := &ess20220222.UntagResourcesRequest{
		RegionId: tea.String(esstag.Region),
		ResourceIds: []*string{
			&essId,
		},
		ResourceType: tea.String("scalinggroup"),
	}
	for k := range state.Tags.Elements() {
		tmpKey := k
		input.TagKeys = append(input.TagKeys, &tmpKey)
	}

	tflog.Info(ctx, "Delete Configured---->", map[string]any{
		"region": esstag.Region,
		"essId":  state.EssId.ValueString(),
		"tagKey": input.TagKeys,
	})
	_, err := esstag.Client.UntagResources(input)
	if err != nil {
		resp.Diagnostics.AddError("could not delete ess tag", err.Error())
		return
	}
}

func (esstag *EssTagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	providerClient, ok := req.ProviderData.(*ProviderClient)
	if !ok {
		resp.Diagnostics.AddError(
			"断言 ProviderClient失败",
			fmt.Sprintf("Expected *ProviderData.ProviderClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	client, err := providerClient.NewEssClient()
	if err != nil {
		resp.Diagnostics.AddError("创建ess client 失败: ", err.Error())
	}
	esstag.Client = client
	esstag.Region = providerClient.GetRegion()
}
