package provider

import (
	"context"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	ess20220222 "github.com/alibabacloud-go/ess-20220222/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/shencan/terraform-provider-aliyun/internal/provider/models"
)

type ProviderClient struct {
	Ak     string
	Sk     string
	Region string
}

func NewProviderClient(_ context.Context, data models.ProviderModel) *ProviderClient {
	return &ProviderClient{
		Ak:     data.Ak.ValueString(),
		Sk:     data.Sk.ValueString(),
		Region: data.Region.ValueString(),
	}
}

func (c *ProviderClient) NewEssClient() (*ess20220222.Client, error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(c.Ak),
		AccessKeySecret: tea.String(c.Sk),
	}
	// config.Endpoint = tea.String("ess.aliyuncs.com")
	config.Endpoint = tea.String(fmt.Sprintf("ess.%s.aliyuncs.com", c.Region))
	return ess20220222.NewClient(config)
}

func (c *ProviderClient) GetRegion() string {
	return c.Region
}
