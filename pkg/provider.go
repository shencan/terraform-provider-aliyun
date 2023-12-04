package pkg

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/shencan/terraform-provider-aliyun/pkg/models"

	"github.com/hashicorp/terraform-plugin-framework/provider/schema"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
)

var _ provider.Provider = &AliYunProvider{}

type AliYunProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &AliYunProvider{
			version: version,
		}
	}
}

func (a *AliYunProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "aliyun"
	resp.Version = a.version
}

func (a *AliYunProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"ak": schema.StringAttribute{
				Optional:    true,
				Description: "The aliyun sdk ak",
			},
			"sk": schema.StringAttribute{
				Optional:    true,
				Description: "The aliyun sdk sk",
			},
			"region": schema.StringAttribute{
				Optional:    true,
				Description: "The aliyun sdk region",
			},
		},
		Description:         "aliyun provider config",
		MarkdownDescription: "aliyun provider config",
	}
}

func (a *AliYunProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data models.ProviderModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if data.Ak.String() == "" || data.Sk.String() == "" || data.Region.String() == "" {
		resp.Diagnostics.AddError("Ak Sk Region", "is empty")
		return
	}

	client := NewProviderClient(ctx, data)
	resp.ResourceData = client
	resp.DataSourceData = client
}

func (a *AliYunProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	//return []func() datasource.DataSource{
	//	NewEssDataSource,
	//}
	return nil
}

func (a *AliYunProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewEssTagResource,
	}
}
